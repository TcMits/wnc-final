package transaction

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/shopspring/decimal"
)

func (uc *CustomerTransactionConfirmSuccessUseCase) createTxcFee(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	return uc.tCUC.Create(ctx, i)
}

func (uc *CustomerTransactionUpdateUseCase) Update(ctx context.Context, e *model.Transaction, i *model.TransactionUpdateInput) (*model.Transaction, error) {
	if e.Status == transaction.StatusSuccess {
		return e, usecase.WrapError(fmt.Errorf("the transaction can not be updated"))
	}
	e, err := uc.repoUpdate.Update(ctx, e, i)
	if err != nil {
		return e, usecase.WrapError(err)
	}
	return e, nil
}
func (uc *CustomerTransactionConfirmSuccessUseCase) subtractSenderBankAccount(ctx context.Context, txc *model.Transaction) (*model.BankAccount, error) {
	bk, _ := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		ID: txc.SenderID,
	})
	am, _ := txc.Amount.Float64()
	bk, err := uc.bAUUC.Update(ctx, bk, &model.BankAccountUpdateInput{
		CashOut: generic.GetPointer(bk.CashOut - am),
	})
	if err != nil {
		return nil, err
	}
	return bk, nil
}
func (uc *CustomerTransactionConfirmSuccessUseCase) addReceiverBankAccount(ctx context.Context, txc *model.Transaction) (*model.BankAccount, error) {
	bk, _ := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		ID: txc.ReceiverID,
	})
	am, _ := txc.Amount.Float64()
	bk, err := uc.bAUUC.Update(ctx, bk, &model.BankAccountUpdateInput{
		CashIn: generic.GetPointer(bk.CashIn + am),
	})
	if err != nil {
		return nil, err
	}
	return bk, nil
}
func (uc *CustomerTransactionValidateConfirmInputUseCase) ValidateConfirmInput(ctx context.Context, e *model.Transaction, token *string) error {
	if e.Status == transaction.StatusDraft {
		pl, err := usecase.ParseConfirmTxcToken(ctx, *token, *uc.cfUC.GetSecret())
		if err != nil {
			return usecase.WrapError(err)
		}
		iFPBMAny, ok := pl["is_fee_paid_by_me"]
		if !ok {
			return usecase.WrapError(fmt.Errorf("invalid token"))
		}
		_, ok = iFPBMAny.(bool)
		if !ok {
			return usecase.WrapError(fmt.Errorf("invalid token"))
		}
		return nil
	}
	return usecase.WrapError(fmt.Errorf("cannot confirm %s transaction", e.Status))
}
func (uc *CustomerTransactionConfirmSuccessUseCase) ConfirmAsSuccess(ctx context.Context, e *model.Transaction, token *string) (*model.Transaction, error) {
	if e.TransactionType == transaction.TransactionTypeInternal {
		_, err := uc.subtractSenderBankAccount(ctx, e)
		if err != nil {
			return nil, err
		}
		_, err = uc.addReceiverBankAccount(ctx, e)
		if err != nil {
			return nil, err
		}
		txc, err := uc.tUUC.Update(ctx, e, &model.TransactionUpdateInput{
			Status: generic.GetPointer(transaction.StatusSuccess),
		})
		if err != nil {
			return nil, err
		}
		ni := &model.TransactionCreateInput{
			SourceTransactionID: &txc.ID,
			Amount:              decimal.NewFromFloat(*uc.cfUC.GetFeeAmount()),
			Status:              generic.GetPointer(transaction.StatusSuccess),
			TransactionType:     transaction.TransactionTypeInternal,
			Description:         uc.cfUC.GetFeeDesc(),
		}
		pl, _ := usecase.ParseConfirmTxcToken(ctx, *token, *uc.cfUC.GetSecret())
		iFPBMAny := pl["is_fee_paid_by_me"]
		isFeePaidByMe, _ := iFPBMAny.(bool)
		if isFeePaidByMe {
			ni.SenderID = *e.SenderID
			ni.SenderBankAccountNumber = e.SenderBankAccountNumber
			ni.SenderBankName = e.SenderBankName
		} else {
			ni.SenderID = *e.ReceiverID
			ni.SenderBankAccountNumber = e.ReceiverBankAccountNumber
			ni.SenderBankName = e.ReceiverBankName
		}
		_, err = uc.createTxcFee(ctx, ni)
		if err != nil {
			return nil, err
		}
		return txc, nil
	}
	return nil, usecase.WrapError(fmt.Errorf("unhandled external transaction case"))
}
func (uc *CustomerTransactionCreateUseCase) Create(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	return uc.repoCreate.Create(ctx, i)
}

func (uc *CustomerTransactionValidateCreateInputUseCase) doesHaveDraftTxc(ctx context.Context, i *model.TransactionCreateInput) error {
	user := usecase.GetUserAsCustomer(ctx)
	l, o := 1, 0
	stsDrf := transaction.StatusDraft
	entities, err := uc.tLUC.List(ctx, &l, &o, nil, &model.TransactionWhereInput{
		HasSenderWith: []*model.BankAccountWhereInput{
			{
				CustomerID: &user.ID,
			},
		},
		Status: &stsDrf,
	})
	if err != nil {
		return usecase.WrapError(err)
	}
	if len(entities) > 0 {
		return usecase.WrapError(fmt.Errorf("there is a draft transaction to be processed. Cannot create a new transaction"))
	}
	return nil
}

func (uc *CustomerTransactionValidateCreateInputUseCase) Validate(ctx context.Context, i *model.TransactionCreateInput, isFeePaidByMe bool) (*model.TransactionCreateInput, error) {
	ba, err := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: &i.SenderID})
	if err != nil {
		return nil, err
	}
	if ba == nil {
		return nil, usecase.WrapError(fmt.Errorf("bank account sender is invalid"))
	}
	if !ba.IsForPayment {
		return nil, usecase.WrapError(fmt.Errorf("bank account sender is not for payment"))
	}
	err = uc.doesHaveDraftTxc(ctx, i)
	if err != nil {
		return nil, err
	}
	am, _ := i.Amount.Float64()
	if isFeePaidByMe {
		if err = ba.IsBalanceSufficient(am + *uc.cfUC.GetFeeAmount()); err != nil {
			return nil, usecase.WrapError(err)
		}
	} else if err = ba.IsBalanceSufficient(am); err != nil {
		return nil, usecase.WrapError(err)
	}
	if i.TransactionType == transaction.TransactionTypeInternal {
		baOther, err := uc.bAGFUC.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: i.ReceiverID})
		if err != nil {
			return nil, err
		}
		if baOther == nil {
			return nil, usecase.WrapError(fmt.Errorf("bank account receiver is invalid"))
		}
		if !baOther.IsForPayment {
			return nil, usecase.WrapError(fmt.Errorf("bank account receiver is not for payment"))
		}
		if !isFeePaidByMe {
			if err = baOther.IsBalanceSufficient(*uc.cfUC.GetFeeAmount()); err != nil {
				return nil, usecase.WrapError(err)
			}
		}
		other, err := uc.cGFUC.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: &baOther.CustomerID})
		if err != nil {
			return nil, err
		}
		i.ReceiverBankAccountNumber = baOther.AccountNumber
		i.ReceiverBankName = *uc.cfUC.GetProductOwnerName()
		i.ReceiverName = other.GetName()
	}
	i.Status = generic.GetPointer(transaction.StatusDraft)
	i.SenderBankAccountNumber = ba.AccountNumber
	i.SenderBankName = *uc.cfUC.GetProductOwnerName()
	user := usecase.GetUserAsCustomer(ctx)
	i.SenderName = user.GetName()
	return i, nil
}

func (uc *CustomerTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	return uc.repoList.List(ctx, limit, offset, o, w)
}

func (uc *CustomerTransactionListMyTxcUseCase) ListMyTxc(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.TransactionWhereInput)
	}
	w.Or = []*model.TransactionWhereInput{
		{
			HasReceiverWith: []*model.BankAccountWhereInput{{CustomerID: &user.ID}},
			HasSenderWith:   []*model.BankAccountWhereInput{{CustomerID: &user.ID}},
		},
	}
	return uc.tLUC.List(ctx, limit, offset, o, w)
}

func (uc *CustomerTransactionGetFirstMyTxcUseCase) GetFirstMyTxc(ctx context.Context, o *model.TransactionOrderInput, w *model.TransactionWhereInput) (*model.Transaction, error) {
	l, of := 1, 0
	entities, err := uc.tLMTUC.ListMyTxc(ctx, &l, &of, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.transaction.GetFirstMyTxc: %s", err))
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}
