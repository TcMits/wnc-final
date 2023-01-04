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
)
