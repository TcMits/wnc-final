package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/TcMits/wnc-final/pkg/tool/password"
	"github.com/TcMits/wnc-final/pkg/tool/template"
)

type (
	CustomerLoginUseCase struct {
		UC1        usecase.ICustomerGetUserUseCase
		SecretKey  *string
		RefreshTTL time.Duration
		AccessTTL  time.Duration
	}
	CustomerValidateLoginInputUseCase struct {
		UC1 usecase.ICustomerGetUserUseCase
	}
	CustomerRenewAccessTokenUseCase struct {
		UC1       usecase.ICustomerGetUserUseCase
		UC2       usecase.ICustomerUpdateUseCase
		SecretKey *string
		AccessTTL time.Duration
	}
	CustomerLogoutUseCase struct {
		UC1 usecase.ICustomerUpdateUseCase
	}
	CustomerValidateChangePasswordWithTokenUseCase struct {
		UC1 usecase.ICustomerConfigUseCase
		UC2 usecase.ICustomerGetFirstUseCase
	}
	CustomerChangePasswordWithTokenUseCase struct {
		UC1 usecase.ICustomerUpdateUseCase
	}
	CustomerForgetPasswordUseCase struct {
		UC1                  usecase.ICustomerConfigUseCase
		TaskExecutor         task.IExecuteTask[*mail.EmailPayload]
		ForgetPwdSubjectMail *string
		ForgetPwdMailTemp    *string
		OtpTimeout           time.Duration
	}
	CustomerValidateForgetPassword struct {
		UC1 usecase.ICustomerGetFirstUseCase
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

func invalidateToken[ModelType, UpdateInput any](
	ctx context.Context,
	handler func(context.Context, ModelType, UpdateInput) (ModelType, error),
	user ModelType,
	input UpdateInput,
) (ModelType, error) {
	user, err := handler(ctx, user, input)
	if err != nil {
		return generic.Zero[ModelType](), usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.auth.invalidateToken: %s", err))
	}
	return user, nil
}

func invalidateCustomerToken(
	ctx context.Context,
	handler usecase.ICustomerUpdateUseCase,
	user *model.Customer,
) (*model.Customer, error) {
	user, err := invalidateToken(ctx, handler.Update, user, &model.CustomerUpdateInput{
		ClearJwtTokenKey: true,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *CustomerLoginUseCase) Login(ctx context.Context, input *model.CustomerLoginInput) (any, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	entity := entityAny.(*model.Customer)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	payload := map[string]any{
		"username": entity.Username,
		"password": entity.Password,
		"jwt_key":  entity.JwtTokenKey,
	}
	tokenPair, err := jwt.NewTokenPair(
		*s.SecretKey,
		payload,
		payload,
		s.AccessTTL,
		s.RefreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.Login: %w", err))
	}
	return tokenPair, nil
}
func (s *CustomerValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.CustomerLoginInput,
) (*model.CustomerLoginInput, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Customer)
	if entity == nil {
		return nil, usecase.ValidationError((fmt.Errorf("invalid username")))
	}
	if !entity.IsActive {
		return nil, usecase.ValidationError(fmt.Errorf("user is not active"))
	}
	err = password.ValidatePassword(entity.Password, *input.Password)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("password is invalid"))
	}
	return input, nil
}

func (s *CustomerRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *s.SecretKey)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid token"))
	}
	userAny, err := s.UC1.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, err
	}
	user := userAny.(*model.Customer)
	_, err = invalidateCustomerToken(ctx, s.UC2, user)
	if err != nil {
		return nil, err
	}
	token, err := jwt.NewAccessToken(payload, *s.SecretKey, s.AccessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.CustomerRenewAccessTokenUseCase.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (s *CustomerLogoutUseCase) Logout(
	ctx context.Context,
) error {
	user := usecase.GetUserAsCustomer(ctx)
	_, err := invalidateCustomerToken(ctx, s.UC1, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerForgetPasswordUseCase) ForgetPassword(ctx context.Context, i *model.CustomerForgetPasswordInput) (*model.CustomerForgetPasswordResp, error) {
	otp := usecase.GenerateOTP(6)
	msg, err := template.RenderFileToStr(*s.ForgetPwdMailTemp, map[string]string{
		"otp":     otp,
		"name":    i.User.GetName(),
		"expires": fmt.Sprintf("%.0f", s.OtpTimeout.Minutes()),
	}, ctx)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.CustomerForgetPasswordUseCase.ForgetPassword: %s", err))
	}
	err = s.TaskExecutor.ExecuteTask(ctx, &mail.EmailPayload{
		Subject: *s.ForgetPwdSubjectMail,
		Message: *msg,
		To:      []string{i.User.Email},
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.CustomerForgetPasswordUseCase.ForgetPassword: %s", err))
	}
	otpHashValue, err := usecase.GenerateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, i.User), otp))
	if err != nil {
		return nil, err
	}
	tk, err := usecase.GenerateForgetPwdToken(
		ctx,
		otpHashValue,
		i.User.Email,
		*s.UC1.GetSecret(),
		s.OtpTimeout,
	)
	if err != nil {
		return nil, err
	}
	return &model.CustomerForgetPasswordResp{Token: tk}, nil
}

func (s *CustomerValidateForgetPassword) ValidateForgetPassword(ctx context.Context, i *model.CustomerForgetPasswordInput) (*model.CustomerForgetPasswordInput, error) {
	user, err := s.UC1.GetFirst(ctx, nil, &model.CustomerWhereInput{Email: &i.Email})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, usecase.ValidationError(fmt.Errorf("user does not exist"))
	}
	i.User = user
	return i, nil
}

func (s *CustomerChangePasswordWithTokenUseCase) ChangePasswordWithToken(ctx context.Context, i *model.CustomerChangePasswordWithTokenInput) error {
	_, err := s.UC1.Update(ctx, i.User, &model.CustomerUpdateInput{
		ClearPassword: true,
		Password:      i.HashPwd,
	})
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.CustomerChangePasswordWithTokenUseCase.ChangePasswordWithToken: %s", err))
	}
	return nil
}

func (s *CustomerValidateChangePasswordWithTokenUseCase) ValidateChangePasswordWithToken(ctx context.Context, i *model.CustomerChangePasswordWithTokenInput) (*model.CustomerChangePasswordWithTokenInput, error) {
	pl, err := usecase.ParseToken(ctx, i.Token, *s.UC1.GetSecret())
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("token expired"))
	}
	eAny, ok := pl["email"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("invaid token due to email missing"))
	}
	tkAny, ok := pl["token"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("invaid token due to token missing"))
	}
	tk, ok := tkAny.(string)
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("invaid token due to token wrong type"))
	}
	email, ok := eAny.(string)
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("invaid token due to email wrong type"))
	}
	user, err := s.UC2.GetFirst(ctx, nil, &model.CustomerWhereInput{Email: &email})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, usecase.ValidationError(fmt.Errorf("user does not exist"))
	}
	err = usecase.ValidateHashInfo(usecase.MakeOTPValue(usecase.EmbedUser(ctx, user), i.Otp), tk)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("otp invalid"))
	}
	if i.Password != i.ConfirmPassword {
		return nil, usecase.ValidationError(fmt.Errorf("password not match"))
	}
	if err = password.ValidatePassword(user.Password, i.Password); err == nil {
		return nil, usecase.ValidationError(fmt.Errorf("new password match old password is not allowed"))
	}
	hashPwd, err := password.GetHashPassword(i.Password)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.CustomerValidateChangePasswordWithTokenUseCase.ValidateChangePasswordWithToken: %w", err))
	}
	i.HashPwd = &hashPwd
	i.User = user
	return i, nil
}

// employee
type (
	EmployeeLoginUseCase struct {
		UC1        usecase.IEmployeeGetUserUseCase
		SecretKey  *string
		RefreshTTL time.Duration
		AccessTTL  time.Duration
	}
	EmployeeValidateLoginInputUseCase struct {
		UC1 usecase.IEmployeeGetUserUseCase
	}
	EmployeeRenewAccessTokenUseCase struct {
		UC1       usecase.IEmployeeGetUserUseCase
		UC2       usecase.IEmployeeUpdateUseCase
		SecretKey *string
		AccessTTL time.Duration
	}
	EmployeeLogoutUseCase struct {
		UC1 usecase.IEmployeeUpdateUseCase
	}
	EmployeeAuthUseCase struct {
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeConfigUseCase
		*EmployeeLoginUseCase
		*EmployeeValidateLoginInputUseCase
		*EmployeeRenewAccessTokenUseCase
		*EmployeeLogoutUseCase
	}
)

func invalidateEmployeeToken(
	ctx context.Context,
	handler usecase.IEmployeeUpdateUseCase,
	user *model.Employee,
) (*model.Employee, error) {
	user, err := invalidateToken(ctx, handler.Update, user, &model.EmployeeUpdateInput{
		ClearJwtTokenKey: true,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *EmployeeLoginUseCase) Login(ctx context.Context, input *model.EmployeeLoginInput) (any, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Employee)
	payload := map[string]any{
		"username": entity.Username,
		"password": entity.Password,
		"jwt_key":  entity.JwtTokenKey,
	}
	tokenPair, err := jwt.NewTokenPair(
		*s.SecretKey,
		payload,
		payload,
		s.AccessTTL,
		s.RefreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.EmployeeLoginUseCase.Login: %w", err))
	}
	return tokenPair, nil
}
func (s *EmployeeValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.EmployeeLoginInput,
) (*model.EmployeeLoginInput, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Employee)
	if entity == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid username"))
	}
	if !entity.IsActive {
		return nil, usecase.ValidationError(fmt.Errorf("user is not active"))
	}
	err = password.ValidatePassword(entity.Password, *input.Password)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("password is invalid"))
	}
	return input, nil
}

func (s *EmployeeRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *s.SecretKey)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid token"))
	}
	userAny, err := s.UC1.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, err
	}
	user := userAny.(*model.Employee)
	_, err = invalidateEmployeeToken(ctx, s.UC2, user)
	if err != nil {
		return nil, err
	}
	token, err := jwt.NewAccessToken(payload, *s.SecretKey, s.AccessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.EmployeeRenewAccessTokenUseCase.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (s *EmployeeLogoutUseCase) Logout(
	ctx context.Context,
) error {
	user := usecase.GetUserAsEmployee(ctx)
	_, err := invalidateEmployeeToken(ctx, s.UC1, user)
	if err != nil {
		return err
	}
	return nil
}

// Admin
type (
	AdminLoginUseCase struct {
		UC1        usecase.IAdminGetUserUseCase
		SecretKey  *string
		RefreshTTL time.Duration
		AccessTTL  time.Duration
	}
	AdminValidateLoginInputUseCase struct {
		UC1 usecase.IAdminGetUserUseCase
	}
	AdminRenewAccessTokenUseCase struct {
		UC1       usecase.IAdminGetUserUseCase
		UC2       usecase.IAdminUpdateUseCase
		SecretKey *string
		AccessTTL time.Duration
	}
	AdminLogoutUseCase struct {
		UC1 usecase.IAdminUpdateUseCase
	}
	AdminAuthUseCase struct {
		usecase.IAdminGetUserUseCase
		usecase.IAdminConfigUseCase
		*AdminLoginUseCase
		*AdminValidateLoginInputUseCase
		*AdminRenewAccessTokenUseCase
		*AdminLogoutUseCase
	}
)

func invalidateAdminToken(
	ctx context.Context,
	handler usecase.IAdminUpdateUseCase,
	user *model.Admin,
) (*model.Admin, error) {
	user, err := invalidateToken(ctx, handler.Update, user, &model.AdminUpdateInput{
		ClearJwtTokenKey: true,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AdminLoginUseCase) Login(ctx context.Context, input *model.AdminLoginInput) (any, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Admin)
	payload := map[string]any{
		"username": entity.Username,
		"password": entity.Password,
		"jwt_key":  entity.JwtTokenKey,
	}
	tokenPair, err := jwt.NewTokenPair(
		*s.SecretKey,
		payload,
		payload,
		s.AccessTTL,
		s.RefreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.AdminLoginUseCase.Login: %w", err))
	}
	return tokenPair, nil
}
func (s *AdminValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.AdminLoginInput,
) (*model.AdminLoginInput, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.Username})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Admin)
	if entity == nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid username"))
	}
	if !entity.IsActive {
		return nil, usecase.ValidationError(fmt.Errorf("user is not active"))
	}
	err = password.ValidatePassword(entity.Password, *input.Password)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("password is invalid"))
	}
	return input, nil
}

func (s *AdminRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *s.SecretKey)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid token"))
	}
	userAny, err := s.UC1.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, err
	}
	user := userAny.(*model.Admin)
	_, err = invalidateAdminToken(ctx, s.UC2, user)
	if err != nil {
		return nil, err
	}
	token, err := jwt.NewAccessToken(payload, *s.SecretKey, s.AccessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.AdminRenewAccessTokenUseCase.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (s *AdminLogoutUseCase) Logout(
	ctx context.Context,
) error {
	user := usecase.GetUserAsAdmin(ctx)
	_, err := invalidateAdminToken(ctx, s.UC1, user)
	if err != nil {
		return err
	}
	return nil
}

// partner
type (
	PartnerGetUserUseCase struct {
		UC1 usecase.IPartnerGetFirstUseCase
	}
	PartnerLoginUseCase struct {
		UC1        usecase.IPartnerGetUserUseCase
		SecretKey  *string
		RefreshTTL time.Duration
		AccessTTL  time.Duration
	}
	PartnerValidateLoginInputUseCase struct {
		UC1 usecase.IPartnerGetUserUseCase
	}
	PartnerAuthUseCase struct {
		*PartnerLoginUseCase
		*PartnerValidateLoginInputUseCase
	}
)

func (s *PartnerGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.UC1.GetFirst(ctx, nil, &model.PartnerWhereInput{
		APIKey: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (s *PartnerLoginUseCase) Login(ctx context.Context, input *model.PartnerLoginInput) (any, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.ApiKey})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Partner)
	if err != nil {
		return nil, usecase.WrapError(err)
	}
	payload := map[string]any{
		"username":  entity.APIKey,
		"is_active": entity.IsActive,
	}
	tokenPair, err := jwt.NewTokenPair(
		*s.SecretKey,
		payload,
		payload,
		s.AccessTTL,
		s.RefreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.auth.auth.PartnerLoginUseCase.Login: %w", err))
	}
	return tokenPair, nil
}

func (s *PartnerValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.PartnerLoginInput,
) (*model.PartnerLoginInput, error) {
	entityAny, err := s.UC1.GetUser(ctx, map[string]any{"username": *input.ApiKey})
	if err != nil {
		return nil, err
	}
	entity := entityAny.(*model.Partner)
	if entity == nil {
		return nil, usecase.ValidationError((fmt.Errorf("invalid api key")))
	}
	if !entity.IsActive {
		return nil, usecase.ValidationError(fmt.Errorf("user is not active"))
	}
	return input, nil
}
