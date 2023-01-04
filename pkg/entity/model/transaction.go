package model

import "github.com/TcMits/wnc-final/ent"

type (
	Transaction                   = ent.Transaction
	TransactionOrderInput         = ent.TransactionOrderInput
	TransactionWhereInput         = ent.TransactionWhereInput
	TransactionUpdateInput        = ent.TransactionUpdateInput
	TransactionCreateInput        = ent.TransactionCreateInput
	TransactionCreateUseCaseInput struct {
		*TransactionCreateInput
		IsFeePaidByMe bool
	}
	TransactionConfirmUseCaseInput struct {
		Otp   string
		Token string
	}
	TransactionCreateResp struct {
		*Transaction
		Token string
	}
)
