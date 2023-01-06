package model

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/debt"
)

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

var DebtStatus = []string{debt.StatusCancelled.String(), debt.StatusFulfilled.String(), debt.StatusPending.String()}
