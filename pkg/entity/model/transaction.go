package model

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/transaction"
)

type (
	ActorType                     string
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
	PartnerTransactionCreateInput struct {
		*TransactionCreateInput
		Token     string
		Signature string
	}
)

var TransactionStatus = []string{transaction.StatusDraft.String(), transaction.StatusVerified.String(), transaction.StatusSuccess.String()}
