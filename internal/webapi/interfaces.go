package webapi

import (
	"context"

	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type ()

type (
	ITPBankGetBankAccount interface {
		Get(context.Context, *model.WhereInputPartner) (*model.BankAccountPartner, error)
	}
	ITPBankInfo interface {
		GetName() string
	}
	ITPBankValidateTransaction interface {
		Validate(context.Context, *model.TransactionCreateInputPartner) error
	}
	ITPBankPreValidateTransaction interface {
		PreValidate(context.Context, *model.TransactionCreateInputPartner) (*model.TransactionCreateInputPartner, error)
	}
	ITPBankCreateTransaction interface {
		Create(context.Context, *model.TransactionCreateInputPartner) error
	}
	ITPBankAPI interface {
		ITPBankGetBankAccount
		ITPBankInfo
		ITPBankValidateTransaction
		ITPBankPreValidateTransaction
		ITPBankCreateTransaction
	}
)
