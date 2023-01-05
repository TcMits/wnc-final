package model

import "github.com/TcMits/wnc-final/ent"

type (
	Debt                      = ent.Debt
	DebtOrderInput            = ent.DebtOrderInput
	DebtWhereInput            = ent.DebtWhereInput
	DebtCreateInput           = ent.DebtCreateInput
	DebtUpdateInput           = ent.DebtUpdateInput
	DebtFulfillWithTokenInput struct {
		Token string
		Otp   string
	}
	DebtFulfillResp struct {
		Token string
	}
)
