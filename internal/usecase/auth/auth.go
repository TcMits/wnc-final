package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/admin"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/employee"
	"github.com/TcMits/wnc-final/internal/usecase/partner"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
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
	CustomerGetUserUseCase struct {
		gFUC usecase.ICustomerGetFirstUseCase
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

func NewCustomerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetUserUseCase {
	uc := &CustomerGetUserUseCase{
		gFUC: customer.NewCustomerGetFirstUseCase(repoList),
	}
	return uc
}
func NewCustomerAuthUseCase(
	taskExctor task.IExecuteTask[*mail.EmailPayload],
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
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
	gUUC := NewCustomerGetUserUseCase(repoList)
	uc := &CustomerAuthUseCase{
		ICustomerGetUserUseCase:                         gUUC,
		ICustomerConfigUseCase:                          config.NewCustomerConfigUseCase(secretKey, prodOwnerName, fee, feeDesc),
		ICustomerForgetPasswordUseCase:                  NewCustomerForgetPasswordUseCase(taskExctor, secretKey, prodOwnerName, feeDesc, forgetPwdEmailSubject, forgetPwdEmailTemplate, fee, otpTimeout),
		ICustomerValidateForgetPasswordUsecase:          NewCustomerValidateForgetPasswordUseCase(repoList),
		ICustomerChangePasswordWithTokenUseCase:         NewCustomerChangePasswordWithTokenUseCase(repoUpdate),
		ICustomerValidateChangePasswordWithTokenUseCase: NewCustomerValidateChangePasswordWithTokenUseCase(repoList, secretKey, prodOwnerName, feeDesc, fee),
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

func invalidateToken[ModelType, UpdateInput any](
	ctx context.Context,
	handler func(context.Context, ModelType, UpdateInput) (ModelType, error),
	user ModelType,
	input UpdateInput,
) (ModelType, error) {
	user, err := handler(ctx, user, input)
	if err != nil {
		return generic.Zero[ModelType](), usecase.WrapError(fmt.Errorf("internal.usecase.auth.auth.invalidateToken: %s", err))
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

func (uc *CustomerLoginUseCase) Login(ctx context.Context, input *model.CustomerLoginInput) (any, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
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

func (uc *CustomerRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *uc.secretKey)
	if err != nil {
		return nil, usecase.ValidationError(fmt.Errorf("invalid token"))
	}
	userAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, err
	}
	user := userAny.(*model.Customer)
	_, err = invalidateCustomerToken(ctx, uc.cUUC, user)
	if err != nil {
		return nil, err
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
	_, err := invalidateCustomerToken(ctx, uc.cUUC, user)
	if err != nil {
		return err
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
		return nil, usecase.ValidationError(fmt.Errorf("user does not exist"))
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
	user, err := s.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{Email: &email})
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
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.CustomerValidateChangePasswordWithTokenUseCase.ValidateChangePasswordWithToken: %w", err))
	}
	i.HashPwd = &hashPwd
	i.User = user
	return i, nil
}

func (useCase *CustomerGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.ValidationError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := useCase.gFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{
		Or: []*model.CustomerWhereInput{
			{Username: &username}, {PhoneNumber: &username}, {Email: &username},
		},
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

// employee
type (
	EmployeeGetUserUseCase struct {
		gFUC usecase.IEmployeeGetFirstUseCase
	}
	EmployeeLoginUseCase struct {
		gUUC       usecase.IEmployeeGetUserUseCase
		secretKey  *string
		refreshTTL time.Duration
		accessTTL  time.Duration
	}
	EmployeeValidateLoginInputUseCase struct {
		gUUC usecase.IEmployeeGetUserUseCase
	}
	EmployeeRenewAccessTokenUseCase struct {
		gUUC      usecase.IEmployeeGetUserUseCase
		secretKey *string
		accessTTL time.Duration
		eUUC      usecase.IEmployeeUpdateUseCase
	}
	EmployeeLogoutUseCase struct {
		eUUC usecase.IEmployeeUpdateUseCase
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

func NewEmployeeGetUserUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeGetUserUseCase {
	uc := &EmployeeGetUserUseCase{
		gFUC: employee.NewEmployeeGetFirstUseCase(repoList),
	}
	return uc
}

func (s *EmployeeGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.gFUC.GetFirst(ctx, nil, &model.EmployeeWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *EmployeeLoginUseCase) Login(ctx context.Context, input *model.EmployeeLoginInput) (any, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
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
		*uc.secretKey,
		payload,
		payload,
		uc.accessTTL,
		uc.refreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.EmployeeLoginUseCase.Login: %w", err))
	}
	return tokenPair, nil
}
func (uc *EmployeeValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.EmployeeLoginInput,
) (*model.EmployeeLoginInput, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
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

func (uc *EmployeeRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *uc.secretKey)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid token"))
	}
	userAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, err
	}
	user := userAny.(*model.Employee)
	_, err = invalidateEmployeeToken(ctx, uc.eUUC, user)
	if err != nil {
		return nil, err
	}
	token, err := jwt.NewAccessToken(payload, *uc.secretKey, uc.accessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.EmployeeRenewAccessTokenUseCase.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (uc *EmployeeLogoutUseCase) Logout(
	ctx context.Context,
) error {
	user := usecase.GetUserAsEmployee(ctx)
	_, err := invalidateEmployeeToken(ctx, uc.eUUC, user)
	if err != nil {
		return err
	}
	return nil
}

func NewEmployeeAuthUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput],
	secretKey,
	prodOwnerName *string,
	refreshTTL,
	accessTTL time.Duration,
) usecase.IEmployeeAuthUseCase {
	gUUC := NewEmployeeGetUserUseCase(repoList)
	uc := &EmployeeAuthUseCase{
		EmployeeLoginUseCase: &EmployeeLoginUseCase{
			gUUC:       gUUC,
			secretKey:  secretKey,
			refreshTTL: refreshTTL,
			accessTTL:  accessTTL,
		},
		EmployeeValidateLoginInputUseCase: &EmployeeValidateLoginInputUseCase{
			gUUC: gUUC,
		},
		EmployeeRenewAccessTokenUseCase: &EmployeeRenewAccessTokenUseCase{
			gUUC:      gUUC,
			secretKey: secretKey,
			accessTTL: accessTTL,
			eUUC:      employee.NewEmployeeUpdateUseCase(repoUpdate),
		},
		EmployeeLogoutUseCase: &EmployeeLogoutUseCase{
			eUUC: employee.NewEmployeeUpdateUseCase(repoUpdate),
		},
		IEmployeeGetUserUseCase: NewEmployeeGetUserUseCase(repoList),
		IEmployeeConfigUseCase:  config.NewEmployeeConfigUseCase(secretKey, prodOwnerName),
	}
	return uc
}

// Admin
type (
	AdminGetUserUseCase struct {
		gFUC usecase.IAdminGetFirstUseCase
	}
	AdminLoginUseCase struct {
		gUUC       usecase.IAdminGetUserUseCase
		secretKey  *string
		refreshTTL time.Duration
		accessTTL  time.Duration
	}
	AdminValidateLoginInputUseCase struct {
		gUUC usecase.IAdminGetUserUseCase
	}
	AdminRenewAccessTokenUseCase struct {
		gUUC      usecase.IAdminGetUserUseCase
		secretKey *string
		accessTTL time.Duration
		eUUC      usecase.IAdminUpdateUseCase
	}
	AdminLogoutUseCase struct {
		eUUC usecase.IAdminUpdateUseCase
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

func NewAdminGetUserUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
) usecase.IAdminGetUserUseCase {
	uc := &AdminGetUserUseCase{
		gFUC: admin.NewAdminGetFirstUseCase(repoList),
	}
	return uc
}

func (s *AdminGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.gFUC.GetFirst(ctx, nil, &model.AdminWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *AdminLoginUseCase) Login(ctx context.Context, input *model.AdminLoginInput) (any, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
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
		*uc.secretKey,
		payload,
		payload,
		uc.accessTTL,
		uc.refreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.AdminLoginUseCase.Login: %w", err))
	}
	return tokenPair, nil
}
func (uc *AdminValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.AdminLoginInput,
) (*model.AdminLoginInput, error) {
	entityAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": *input.Username})
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

func (uc *AdminRenewAccessTokenUseCase) RenewToken(
	ctx context.Context,
	refreshToken *string,
) (any, error) {
	payload, err := jwt.ParseJWT(*refreshToken, *uc.secretKey)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid token"))
	}
	userAny, err := uc.gUUC.GetUser(ctx, map[string]any{"username": payload["username"]})
	if err != nil {
		return nil, err
	}
	user := userAny.(*model.Admin)
	_, err = invalidateAdminToken(ctx, uc.eUUC, user)
	if err != nil {
		return nil, err
	}
	token, err := jwt.NewAccessToken(payload, *uc.secretKey, uc.accessTTL)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.AdminRenewAccessTokenUseCase.RenewToken: %w", err))
	}
	return &jwt.TokenPair{
		RefreshToken: refreshToken,
		AccessToken:  &token,
	}, nil
}

func (uc *AdminLogoutUseCase) Logout(
	ctx context.Context,
) error {
	user := usecase.GetUserAsAdmin(ctx)
	_, err := invalidateAdminToken(ctx, uc.eUUC, user)
	if err != nil {
		return err
	}
	return nil
}

func NewAdminAuthUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Admin, *model.AdminUpdateInput],
	secretKey,
	prodOwnerName *string,
	refreshTTL,
	accessTTL time.Duration,
) usecase.IAdminAuthUseCase {
	gUUC := NewAdminGetUserUseCase(repoList)
	uc := &AdminAuthUseCase{
		AdminLoginUseCase: &AdminLoginUseCase{
			gUUC:       gUUC,
			secretKey:  secretKey,
			refreshTTL: refreshTTL,
			accessTTL:  accessTTL,
		},
		AdminValidateLoginInputUseCase: &AdminValidateLoginInputUseCase{
			gUUC: gUUC,
		},
		AdminRenewAccessTokenUseCase: &AdminRenewAccessTokenUseCase{
			gUUC:      gUUC,
			secretKey: secretKey,
			accessTTL: accessTTL,
			eUUC:      admin.NewAdminUpdateUseCase(repoUpdate),
		},
		AdminLogoutUseCase: &AdminLogoutUseCase{
			eUUC: admin.NewAdminUpdateUseCase(repoUpdate),
		},
		IAdminGetUserUseCase: NewAdminGetUserUseCase(repoList),
		IAdminConfigUseCase:  config.NewAdminConfigUseCase(secretKey, prodOwnerName),
	}
	return uc
}

// partner
type (
	PartnerGetUserUseCase struct {
		gFUC usecase.IPartnerGetFirstUseCase
	}
	PartnerLoginUseCase struct {
		gUUC       usecase.IPartnerGetUserUseCase
		secretKey  *string
		refreshTTL time.Duration
		accessTTL  time.Duration
	}
	PartnerValidateLoginInputUseCase struct {
		gUUC usecase.IPartnerGetUserUseCase
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
	u, err := s.gFUC.GetFirst(ctx, nil, &model.PartnerWhereInput{
		APIKey: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (s *PartnerLoginUseCase) Login(ctx context.Context, input *model.PartnerLoginInput) (any, error) {
	entityAny, err := s.gUUC.GetUser(ctx, map[string]any{"username": *input.ApiKey})
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
		*s.secretKey,
		payload,
		payload,
		s.accessTTL,
		s.refreshTTL,
	)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.auth.auth.PartnerLoginUseCase.Login: %w", err))
	}
	return tokenPair, nil
}

func (s *PartnerValidateLoginInputUseCase) ValidateLoginInput(
	ctx context.Context,
	input *model.PartnerLoginInput,
) (*model.PartnerLoginInput, error) {
	entityAny, err := s.gUUC.GetUser(ctx, map[string]any{"username": *input.ApiKey})
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

func NewPartnerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
) usecase.IPartnerGetUserUseCase {
	uc := &PartnerGetUserUseCase{
		gFUC: partner.NewPartnerGetFirstUseCase(repoList),
	}
	return uc
}
func NewPartnerAuthUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
	secretKey *string,
	refreshTTL,
	accessTTL time.Duration,
) usecase.IPartnerAuthUseCase {
	gUUC := NewPartnerGetUserUseCase(repoList)
	uc := &PartnerAuthUseCase{
		PartnerLoginUseCase: &PartnerLoginUseCase{
			gUUC:       gUUC,
			secretKey:  secretKey,
			refreshTTL: refreshTTL,
			accessTTL:  accessTTL,
		},
		PartnerValidateLoginInputUseCase: &PartnerValidateLoginInputUseCase{
			gUUC: gUUC,
		},
	}
	return uc
}
