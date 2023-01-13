package tpbank

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/TcMits/wnc-final/internal/webapi"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/TcMits/wnc-final/pkg/tool/template"
)

type (
	TPBankSDK struct {
		SecretKey              string
		RSAPrivateK            string
		ApiKey                 string
		AccessToken            string
		Layout                 string
		BaseURL                string
		AuthAPI                string
		BankAccountsAPI        string
		ValidateTransactionAPI string
		CreateTransactionAPI   string
	}
	TPBankInfo struct {
		Name string
	}
	TPBankValidateTransaction struct {
		ApiKey      string
		BaseURL     string
		AccessToken string
	}
	TPBankAPI struct {
		webapi.ITPBankGetBankAccount
		webapi.ITPBankInfo
		webapi.ITPBankPreValidateTransaction
		webapi.ITPBankValidateTransaction
		webapi.ITPBankCreateTransaction
	}
)

func (s *TPBankSDK) PreValidate(ctx context.Context, i *model.TransactionCreateInputPartner) (*model.TransactionCreateInputPartner, error) {
	data, err := template.RenderToStr(s.Layout, map[string]string{
		"receiver_bank_account_number": i.ReceiverBankAccountNumber,
		"sender_bank_account_number":   i.SenderBankAccountNumber,
		"sender_name":                  i.SenderName,
		"amount":                       i.Amount.String(),
		"description":                  i.Description,
	}, ctx)
	if err != nil {
		return nil, err
	}
	hashData := password.GenerateHashData(ctx, s.SecretKey, *data)
	i.Token = hashData
	sig, err := password.GenerateSignature(ctx, hashData, s.RSAPrivateK)
	if err != nil {
		return nil, err
	}
	i.Signature = sig
	return i, nil
}

func (s *TPBankSDK) getValidateTransactionURL(ctx context.Context) string {
	u, _ := url.JoinPath(s.BaseURL, s.ValidateTransactionAPI)
	return u
}
func (s *TPBankSDK) getAuthURL(ctx context.Context) string {
	u, _ := url.JoinPath(s.BaseURL, s.AuthAPI)
	return u
}
func (s *TPBankSDK) getBankAccountURL(ctx context.Context) string {
	u, _ := url.JoinPath(s.BaseURL, s.BankAccountsAPI)
	return u
}
func (s *TPBankSDK) getCreateTransactionURL(ctx context.Context) string {
	return path.Join(s.BaseURL, s.CreateTransactionAPI)
}
func (s *TPBankSDK) makeRequest(ctx context.Context, method, url string, body, headers map[string]string) (*http.Response, error) {
	var handler func(context.Context, string, map[string]string, map[string]string) (*http.Response, error)
	switch method {
	case http.MethodGet:
		handler = MakeGetRequest
	case http.MethodPost:
		handler = MakePostRequest
	}
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Authorization"] = fmt.Sprintf("Bearer %s", s.AccessToken)
	resp, err := handler(ctx, url, body, headers)
	if err != nil {
		return nil, err
	}
	tryTimes := 0
	for (resp.StatusCode == 499 || resp.StatusCode == 498) && tryTimes < 3 {
		err = s.auth(ctx)
		if err != nil {
			return nil, err
		}
		headers["Authorization"] = fmt.Sprintf("Bearer %s", s.AccessToken)
		tryTimes += 1
		resp, err = handler(ctx, url, body, headers)
		if err != nil {
			return nil, err
		}
	}
	return resp, err
}
func (s *TPBankSDK) makePostRequest(ctx context.Context, url string, body, headers map[string]string) (*http.Response, error) {
	return s.makeRequest(ctx, http.MethodPost, url, body, headers)
}
func (s *TPBankSDK) makeGetRequest(ctx context.Context, url string, headers map[string]string) (*http.Response, error) {
	return s.makeRequest(ctx, http.MethodGet, url, nil, headers)
}
func (s *TPBankSDK) Validate(ctx context.Context, i *model.TransactionCreateInputPartner) error {
	resp, err := s.makePostRequest(ctx, s.getValidateTransactionURL(ctx), map[string]string{
		"amount":                       i.Amount.String(),
		"description":                  i.Description,
		"token":                        i.Token,
		"signature":                    i.Signature,
		"sender_name":                  i.SenderName,
		"sender_bank_account_number":   i.SenderBankAccountNumber,
		"receiver_bank_account_number": i.ReceiverBankAccountNumber,
	}, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	type r struct {
		Message string `json:"message"`
	}
	t := new(r)
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	return fmt.Errorf(t.Message)
}
func (s *TPBankSDK) Get(ctx context.Context, w *model.WhereInputPartner) (*model.BankAccountPartner, error) {
	if w == nil {
		w = &model.WhereInputPartner{AccountNumber: "111111111111"}
	}
	url, err := AddQueryParams(s.getBankAccountURL(ctx), map[string]string{"account_number": w.AccountNumber})
	if err != nil {
		return nil, err
	}
	resp, err := s.makeGetRequest(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}
	fmt.Sprintln(resp.StatusCode)
	t := new(model.BankAccountPartner)
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
func (s *TPBankSDK) auth(ctx context.Context) error {
	res, err := s.makePostRequest(ctx, s.getAuthURL(ctx), map[string]string{
		"api_key": s.ApiKey,
	}, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid payload")
	}
	type tokenPair struct {
		AccessToken string `json:"access_token"`
	}
	t := new(tokenPair)
	err = json.NewDecoder(res.Body).Decode(t)
	if err != nil {
		return err
	}
	s.AccessToken = t.AccessToken
	return nil
}
func (s *TPBankSDK) Create(ctx context.Context, i *model.TransactionCreateInputPartner) error {
	resp, err := s.makePostRequest(ctx, s.getCreateTransactionURL(ctx), map[string]string{
		"amount":                       i.Amount.String(),
		"description":                  i.Description,
		"token":                        i.Token,
		"signature":                    i.Signature,
		"sender_name":                  i.SenderName,
		"sender_bank_account_number":   i.SenderBankAccountNumber,
		"receiver_bank_account_number": i.ReceiverBankAccountNumber,
	}, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusCreated {
		return nil
	}
	type r struct {
		Message string `json:"message"`
	}
	t := new(r)
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return err
	}
	return fmt.Errorf(t.Message)
}

func (s *TPBankInfo) GetName() string {
	return s.Name
}

func MakeGetRequest(ctx context.Context, url string, body, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if headers == nil {
		headers = make(map[string]string)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if err != nil {
		return nil, err
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func MakePostRequest(ctx context.Context, url string, body, headers map[string]string) (*http.Response, error) {
	postBody, _ := json.Marshal(body)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(postBody))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, err
}

func AddQueryParams(base string, opts map[string]string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	q := u.Query()
	if opts == nil {
		opts = make(map[string]string)
	}
	for k, v := range opts {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func NewTPBankInfo(name string) webapi.ITPBankInfo {
	return &TPBankInfo{Name: name}
}
func NewTPBankSDK(
	apiKey,
	rsaPrivateK,
	secretKey,
	layout,
	baseUrl,
	authAPI,
	bankAccountAPI,
	createTransactionAPI,
	validateAPI string,
) *TPBankSDK {
	return &TPBankSDK{
		ApiKey:                 apiKey,
		BaseURL:                baseUrl,
		AuthAPI:                authAPI,
		BankAccountsAPI:        bankAccountAPI,
		ValidateTransactionAPI: validateAPI,
		Layout:                 layout,
		SecretKey:              secretKey,
		RSAPrivateK:            rsaPrivateK,
		CreateTransactionAPI:   createTransactionAPI,
	}
}

func NewTPBankAPI(
	name,
	apiKey,
	rsaPrivateK,
	secretKey,
	layout,
	baseUrl,
	authAPI,
	bankAccountAPI,
	createTransactionAPI,
	validateAPI string,
) webapi.ITPBankAPI {
	sdk := NewTPBankSDK(
		apiKey,
		rsaPrivateK,
		secretKey,
		layout,
		baseUrl,
		authAPI,
		bankAccountAPI,
		createTransactionAPI,
		validateAPI,
	)
	return &TPBankAPI{
		ITPBankGetBankAccount:         sdk,
		ITPBankPreValidateTransaction: sdk,
		ITPBankInfo:                   NewTPBankInfo(name),
		ITPBankCreateTransaction:      sdk,
		ITPBankValidateTransaction:    sdk,
	}
}
