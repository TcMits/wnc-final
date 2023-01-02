package model

import "github.com/TcMits/wnc-final/ent"

type (
	Transaction            = ent.Transaction
	TransactionOrderInput  = ent.TransactionOrderInput
	TransactionWhereInput  = ent.TransactionWhereInput
	TransactionCreateInput = ent.TransactionCreateInput
	TransactionUpdateInput = ent.TransactionUpdateInput
	TransactionCreateResp  struct {
		*Transaction
		Token string
	}
)
