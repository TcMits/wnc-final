package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/error/wrapper"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/TcMits/wnc-final/pkg/tool/template"
)

type (
	CustomerLoginUseCase struct {
		gUUC       usecase.ICustomerGetUserUseCase
		secretKey  *string
		refreshTTL time.Duration
		accessTTL  time.Duration
	}
	CustomerValidateLoginInputUseCase struct {
		gUUC usecase.ICustomerGetUserUseCase
	}
	CustomerRenewAccessTokenUseCase struct {
		gUUC      usecase.ICustomerGetUserUseCase
		secretKey *string
		accessTTL time.Duration
		cUUC      usecase.ICustomerUpdateUseCase
	}
	CustomerLogoutUseCase struct {
		cUUC usecase.ICustomerUpdateUseCase
	}
	CustomerValidateChangePasswordWithTokenUseCase struct {
		cfUC  usecase.ICustomerConfigUseCase
		cGFUC usecase.ICustomerGetFirstUseCase
	}
	CustomerChangePasswordWithTokenUseCase struct {
		cUUC usecase.ICustomerUpdateUseCase
	}
	CustomerForgetPasswordUseCase struct {
		cfUC                 usecase.ICustomerConfigUseCase
		taskExecutor         task.IExecuteTask[*mail.EmailPayload]
		forgetPwdSubjectMail *string
		forgetPwdMailTemp    *string
		otpTimeout           time.Duration
	}
	CustomerValidateForgetPassword struct {
		cGFUC usecase.ICustomerGetFirstUseCase
	}
	CustomerAuthUseCase struct {
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerForgetPasswordUseCase
		usecase.ICustomerValidateForgetPasswordUsecase
		usecase.ICustomerChangePasswordWithTokenUseCase
		usecase.ICustomerValidateChangePasswordWithTokenUseCase
		*CustomerLoginUseCase
		*CustomerValidateLoginInputUseCase
		*CustomerRenewAccessTokenUseCase
		*CustomerLogoutUseCase
	}
)

func NewCustomerForgetPasswordUseCase(
	taskExctor task.IExecuteTask[*mail.EmailPayload],
	sk,
	prodOwnerName,
	feeDesc,
	forgetPwdEmailSubject,
	forgetPwdEmailTemplate *string,
	fee *float64,
	otpTimeout time.Duration,
) usecase.ICustomerForgetPasswordUseCase {
	return &CustomerForgetPasswordUseCase{
		cfUC:                 config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		taskExecutor:         taskExctor,
		forgetPwdSubjectMail: forgetPwdEmailSubject,
		forgetPwdMailTemp:    forgetPwdEmailTemplate,
		otpTimeout:           otpTimeout,
	}
}
func NewCustomerValidateForgetPasswordUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerValidateForgetPasswordUsecase {
	return &CustomerValidateForgetPassword{
		cGFUC: customer.NewCustomerGetFirstUseCase(rlc),
	}
}
func NewCustomerChangePasswordWithTokenUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
) usecase.ICustomerChangePasswordWithTokenUseCase {
	return &CustomerChangePasswordWithTokenUseCase{cUUC: customer.NewCustomerUpdateUseCase(repoUpdate)}
}
func NewCustomerValidateChangePasswordWithTokenUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	secretKey,
	prodOwnerName,
	feeDesc *string,
	fee *float64,
) usecase.ICustomerValidateChangePasswordWithTokenUseCase {
	return &CustomerValidateChangePasswordWithTokenUseCase{
		cfUC:  config.NewCustomerConfigUseCase(secretKey, prodOwnerName, fee, feeDesc),
		cGFUC: customer.NewCustomerGetFirstUseCase(rlc),
	}
}

func NewCustomerAuthUseCase(
	taskExctor task.IExecuteTask[*mail.EmailPayload],
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	secretKey,
	prodOwnerName,
	feeDesc,
	forgetPwdEmailSubject,
	forgetPwdEmailTemplate *string,
	fee *float64,
	otpTimeout,
	refreshTTL,
	accessTTL time.Duration,
) usecase.ICustomerAuthUseCase {
	gUUC := me.NewCustomerGetUserUseCase(repoList)
	uc := &CustomerAuthUseCase{
		ICustomerGetUserUseCase:                         gUUC,
		ICustomerConfigUseCase:                          config.NewCustomerConfigUseCase(secretKey, prodOwnerName, fee, feeDesc),
		ICustomerForgetPasswordUseCase:                  NewCustomerForgetPasswordUseCase(taskExctor, secretKey, prodOwnerName, feeDesc, forgetPwdEmailSubject, forgetPwdEmailTemplate, fee, otpTimeout),
		ICustomerValidateForgetPasswordUsecase:          NewCustomerValidateForgetPasswordUseCase(rlc),
		ICustomerChangePasswordWithTokenUseCase:         NewCustomerChangePasswordWithTokenUseCase(repoUpdate),
		ICustomerValidateChangePasswordWithTokenUseCase: NewCustomerValidateChangePasswordWithTokenUseCase(rlc, secretKey, prodOwnerName, feeDesc, fee),
		CustomerLoginUseCase: &CustomerLoginUseCase{
			gUUC:       gUUC,
			secretKey:  secretKey,
			refreshTTL: refreshTTL,
			accessTTL:  accessTTL,
		},
		CustomerValidateLoginInputUseCase: &CustomerValidateLoginInputUseCase{
			gUUC: gUUC,
		},
		CustomerRenewAccessTokenUseCase: &CustomerRenewAccessTokenUseCase{
			gUUC:      gUUC,
			secretKey: secretKey,
			accessTTL: accessTTL,
			cUUC:      customer.NewCustomerUpdateUseCase(repoUpdate),
		},
		CustomerLogoutUseCase: &CustomerLogoutUseCase{
			cUUC: customer.NewCustomerUpdateUseCase(repoUpdate),
		},
	}
	return uc
}

func invalidateToken(
	ctx context.Context,
	handler usecase.ICustomerUpdateUseCase,
	user *model.Customer,
) (*model.Customer, error) {
	user, err := handler.Update(ctx, user, &model.CustomerUpdateInput{
		ClearJwtTokenKey: true,
	})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	return user, nil
}

func (uc *CustomerLoginUseCase) Login(ctx context.Context, input *model.CustomerLoginInput) (any, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	entity := entityAny.(*model.Customer)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	if !entity.IsActive {
		return nil, usecase.WrapError(fmt.Errorf("user is not active"))
	}
	payload := map[string]any{
		"username": entity.Username,
		"password": entity.Password,
		"jwt_key":  entity.JwtTokenKey,
	}
	tokenPair, err := jwt.NewTokenPair(
		*uc.secretKey,
		payload,
		payload,
		uc.accessTTL,
		uc.refreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.Login: %w", err))
	}
	return tokenPair, nil
}
func (uc *CustomerValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.CustomerLoginInput,
) (*model.CustomerLoginInput, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, err
	}
	if entityAny == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid username"))
	}
	entity := entityAny.(*model.Customer)
	err = password.ValidatePassword(entity.Password, *input.Password)
	if err != nil {
		return nil, usecase.WrapError(wrapper.NewValidationError(fmt.Errorf("password is invalid")))
	}
	return input, nil
}

func (uc *CustomerRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *uc.secretKey)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	userAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	user := userAny.(*model.Customer)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	_, err = invalidateToken(ctx, uc.cUUC, user)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	token, err := jwt.NewAccessToken(payload, *uc.secretKey, uc.accessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerRenewAccessTokenUseCase.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (uc *CustomerLogoutUseCase) Logout(
	ctx context.Context,
) error {
	user := usecase.GetUserAsCustomer(ctx)
	_, err := invalidateToken(ctx, uc.cUUC, user)
	if err != nil {
		return usecase.WrapError(err)
	}
	return nil
}

func (s *CustomerForgetPasswordUseCase) ForgetPassword(ctx context.Context, i *model.CustomerForgetPasswordInput) (*model.CustomerForgetPasswordResp, error) {
	otp := usecase.GenerateOTP(6)
	msg, err := template.RenderToStr(*s.forgetPwdMailTemp, map[string]string{
		"otp":     otp,
		"name":    i.User.GetName(),
		"expires": fmt.Sprintf("%.0f", s.otpTimeout.Minutes()),
	}, ctx)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerForgetPasswordUseCase.ForgetPassword: %s", err))
	}
	err = s.taskExecutor.ExecuteTask(ctx, &mail.EmailPayload{
		Subject: *s.forgetPwdSubjectMail,
		Message: *msg,
		To:      []string{i.User.Email},
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerForgetPasswordUseCase.ForgetPassword: %s", err))
	}
	otpHashValue, err := usecase.GenerateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, i.User), otp))
	if err != nil {
		return nil, err
	}
	tk, err := usecase.GenerateForgetPwdToken(
		ctx,
		otpHashValue,
		i.User.Email,
		*s.cfUC.GetSecret(),
		s.otpTimeout,
	)
	if err != nil {
		return nil, err
	}
	return &model.CustomerForgetPasswordResp{Token: tk}, nil
}

func (s *CustomerValidateForgetPassword) ValidateForgetPassword(ctx context.Context, i *model.CustomerForgetPasswordInput) (*model.CustomerForgetPasswordInput, error) {
	user, err := s.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{Email: &i.Email})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, usecase.WrapError(fmt.Errorf("user does not exist"))
	}
	i.User = user
	return i, nil
}

func (s *CustomerChangePasswordWithTokenUseCase) ChangePasswordWithToken(ctx context.Context, i *model.CustomerChangePasswordWithTokenInput) error {
	_, err := s.cUUC.Update(ctx, i.User, &model.CustomerUpdateInput{
		ClearPassword: true,
		Password:      i.HashPwd,
	})
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerChangePasswordWithTokenUseCase.ChangePasswordWithToken: %s", err))
	}
	return nil
}

func (s *CustomerValidateChangePasswordWithTokenUseCase) ValidateChangePasswordWithToken(ctx context.Context, i *model.CustomerChangePasswordWithTokenInput) (*model.CustomerChangePasswordWithTokenInput, error) {
	pl, err := usecase.ParseToken(ctx, i.Token, *s.cfUC.GetSecret())
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("token expired"))
	}
	eAny, ok := pl["email"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("invaid token due to email missing"))
	}
	tkAny, ok := pl["token"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("invaid token due to token missing"))
	}
	tk, ok := tkAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("invaid token due to token wrong type"))
	}
	email, ok := eAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("invaid token due to email wrong type"))
	}
	user, err := s.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{Email: &email})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, usecase.WrapError(fmt.Errorf("user does not exist"))
	}
	err = usecase.ValidateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, user), i.Otp), tk)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("otp invalid"))
	}
	if i.Password != i.ConfirmPassword {
		return nil, usecase.WrapError(fmt.Errorf("password not match"))
	}
	if err = password.ValidatePassword(user.Password, i.Password); err == nil {
		return nil, usecase.WrapError(fmt.Errorf("new password match old password is not allowed"))
	}
	hashPwd, err := password.GetHashPassword(i.Password)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerValidateChangePasswordWithTokenUseCase.ValidateChangePasswordWithToken: %w", err))
	}
	i.HashPwd = &hashPwd
	i.User = user
	return i, nil
}
