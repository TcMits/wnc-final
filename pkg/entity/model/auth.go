package model

type (
	CustomerLoginInput struct {
		Username *string
		Password *string
	}
	CustomerChangePasswordInput struct {
		OldPassword     string
		Password        string
		ConfirmPassword string
		HashPwd         *string
	}
	CustomerChangePasswordWithTokenInput struct {
		Token           string
		Otp             string
		Password        string
		ConfirmPassword string
		HashPwd         *string
		User            *Customer
	}
	CustomerForgetPasswordInput struct {
		Email string
		User  *Customer
	}
	CustomerForgetPasswordResp struct {
		Token string
	}
)

type (
	EmployeeLoginInput struct {
		Username *string
		Password *string
	}
	AdminLoginInput struct {
		Username *string
		Password *string
	}
	PartnerLoginInput struct {
		ApiKey *string
	}
)
