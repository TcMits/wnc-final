package constructors

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/auth"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
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
	return &auth.CustomerForgetPasswordUseCase{
		UC1:                  NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		TaskExecutor:         taskExctor,
		ForgetPwdSubjectMail: forgetPwdEmailSubject,
		ForgetPwdMailTemp:    forgetPwdEmailTemplate,
		OtpTimeout:           otpTimeout,
	}
}
func NewCustomerValidateForgetPasswordUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerValidateForgetPasswordUsecase {
	return &auth.CustomerValidateForgetPassword{
		UC1: NewCustomerGetFirstUseCase(rlc),
	}
}
func NewCustomerChangePasswordWithTokenUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
) usecase.ICustomerChangePasswordWithTokenUseCase {
	return &auth.CustomerChangePasswordWithTokenUseCase{UC1: NewCustomerUpdateUseCase(repoUpdate)}
}
func NewCustomerValidateChangePasswordWithTokenUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	secretKey,
	prodOwnerName,
	feeDesc *string,
	fee *float64,
) usecase.ICustomerValidateChangePasswordWithTokenUseCase {
	return &auth.CustomerValidateChangePasswordWithTokenUseCase{
		UC1: NewCustomerConfigUseCase(secretKey, prodOwnerName, fee, feeDesc),
		UC2: NewCustomerGetFirstUseCase(rlc),
	}
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
	uc := &auth.CustomerAuthUseCase{
		ICustomerGetUserUseCase:                         gUUC,
		ICustomerConfigUseCase:                          NewCustomerConfigUseCase(secretKey, prodOwnerName, fee, feeDesc),
		ICustomerForgetPasswordUseCase:                  NewCustomerForgetPasswordUseCase(taskExctor, secretKey, prodOwnerName, feeDesc, forgetPwdEmailSubject, forgetPwdEmailTemplate, fee, otpTimeout),
		ICustomerValidateForgetPasswordUsecase:          NewCustomerValidateForgetPasswordUseCase(repoList),
		ICustomerChangePasswordWithTokenUseCase:         NewCustomerChangePasswordWithTokenUseCase(repoUpdate),
		ICustomerValidateChangePasswordWithTokenUseCase: NewCustomerValidateChangePasswordWithTokenUseCase(repoList, secretKey, prodOwnerName, feeDesc, fee),
		CustomerLoginUseCase: &auth.CustomerLoginUseCase{
			UC1:        gUUC,
			SecretKey:  secretKey,
			RefreshTTL: refreshTTL,
			AccessTTL:  accessTTL,
		},
		CustomerValidateLoginInputUseCase: &auth.CustomerValidateLoginInputUseCase{
			UC1: gUUC,
		},
		CustomerRenewAccessTokenUseCase: &auth.CustomerRenewAccessTokenUseCase{
			UC1:       gUUC,
			UC2:       NewCustomerUpdateUseCase(repoUpdate),
			SecretKey: secretKey,
			AccessTTL: accessTTL,
		},
		CustomerLogoutUseCase: &auth.CustomerLogoutUseCase{
			UC1: NewCustomerUpdateUseCase(repoUpdate),
		},
	}
	return uc
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
	uc := &auth.EmployeeAuthUseCase{
		EmployeeLoginUseCase: &auth.EmployeeLoginUseCase{
			UC1:        gUUC,
			SecretKey:  secretKey,
			RefreshTTL: refreshTTL,
			AccessTTL:  accessTTL,
		},
		EmployeeValidateLoginInputUseCase: &auth.EmployeeValidateLoginInputUseCase{
			UC1: gUUC,
		},
		EmployeeRenewAccessTokenUseCase: &auth.EmployeeRenewAccessTokenUseCase{
			UC1:       gUUC,
			SecretKey: secretKey,
			AccessTTL: accessTTL,
			UC2:       NewEmployeeUpdateUseCase(repoUpdate),
		},
		EmployeeLogoutUseCase: &auth.EmployeeLogoutUseCase{
			UC1: NewEmployeeUpdateUseCase(repoUpdate),
		},
		IEmployeeGetUserUseCase: gUUC,
		IEmployeeConfigUseCase:  NewEmployeeConfigUseCase(secretKey, prodOwnerName),
	}
	return uc
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
	uc := &auth.AdminAuthUseCase{
		AdminLoginUseCase: &auth.AdminLoginUseCase{
			UC1:        gUUC,
			SecretKey:  secretKey,
			RefreshTTL: refreshTTL,
			AccessTTL:  accessTTL,
		},
		AdminValidateLoginInputUseCase: &auth.AdminValidateLoginInputUseCase{
			UC1: gUUC,
		},
		AdminRenewAccessTokenUseCase: &auth.AdminRenewAccessTokenUseCase{
			UC1:       gUUC,
			UC2:       NewAdminUpdateUseCase(repoUpdate),
			SecretKey: secretKey,
			AccessTTL: accessTTL,
		},
		AdminLogoutUseCase: &auth.AdminLogoutUseCase{
			UC1: NewAdminUpdateUseCase(repoUpdate),
		},
		IAdminGetUserUseCase: gUUC,
		IAdminConfigUseCase:  NewAdminConfigUseCase(secretKey, prodOwnerName),
	}
	return uc
}

func NewPartnerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
) usecase.IPartnerGetUserUseCase {
	uc := &auth.PartnerGetUserUseCase{
		UC1: NewPartnerGetFirstUseCase(repoList),
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
	uc := &auth.PartnerAuthUseCase{
		PartnerLoginUseCase: &auth.PartnerLoginUseCase{
			UC1:        gUUC,
			SecretKey:  secretKey,
			RefreshTTL: refreshTTL,
			AccessTTL:  accessTTL,
		},
		PartnerValidateLoginInputUseCase: &auth.PartnerValidateLoginInputUseCase{
			UC1: gUUC,
		},
	}
	return uc
}
