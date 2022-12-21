package transaction

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

const (
	FeeAmount        float64 = 1
	ProductOwnerName string  = ""
	FeeDesc          string  = "Fee transaction"
)

type (
	CustomerTransactionCreateUseCase struct {
		cBaGetter  usecase.ICustomerBankAccountGetFirstUseCase
		repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
		cBaUpdater usecase.ICustomerBankAccountUpdateUseCase
	}
	CustomerTransactionUpdateUseCase struct {
		cBaUpdater usecase.ICustomerBankAccountUpdateUseCase
		cBaGetter  usecase.ICustomerBankAccountGetFirstUseCase
		repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput]
	}
	CustomerTransactionListUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionValidateCreateInputUseCase struct {
		cBaGetter usecase.ICustomerBankAccountGetFirstUseCase
		cGetter   usecase.ICustomerGetFirstUseCase
		repoList  repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionUseCase struct {
		usecase.ICustomerTransactionValidateCreateInputUseCase
		usecase.ICustomerTransactionCreateUseCase
		usecase.ICustomerTransactionListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
	}
)

func NewCustomerTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListUseCase {
	return &CustomerTransactionListUseCase{
		repoList: repoList,
	}
}

func NewCustomerTransactionCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	rLBa repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rUBa repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.ICustomerTransactionCreateUseCase {
	return &CustomerTransactionCreateUseCase{
		repoCreate: repoCreate,
		cBaGetter:  bankaccount.NewCustomerBankAccountGetFirstUseCase(rLBa),
		cBaUpdater: bankaccount.NewCustomerBankAccountUpdateUseCase(rUBa),
	}
}

func NewCustomerTransactionValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerTransactionValidateCreateInputUseCase {
	return &CustomerTransactionValidateCreateInputUseCase{
		repoList:  repoList,
		cBaGetter: bankaccount.NewCustomerBankAccountGetFirstUseCase(rlba),
		cGetter:   customer.NewCustomerGetFirstUseCase(rlc),
	}
}

func NewCustomerTransactionUseCase(
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rUBa repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	sk *string,
) usecase.ICustomerTransactionUseCase {
	return &CustomerTransactionUseCase{
		ICustomerTransactionCreateUseCase:              NewCustomerTransactionCreateUseCase(repoCreate, rlba, rUBa),
		ICustomerTransactionValidateCreateInputUseCase: NewCustomerTransactionValidateCreateInputUseCase(repoList, rlba, rlc),
		ICustomerTransactionListUseCase:                NewCustomerTransactionListUseCase(repoList),
		ICustomerConfigUseCase:                         config.NewCustomerConfigUseCase(sk),
		ICustomerGetUserUseCase:                        me.NewCustomerGetUserUseCase(rlc),
	}
}

func (uc *CustomerTransactionCreateUseCase) createTxcFee(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	return uc.Create(ctx, i, false)
}

//	func (uc *CustomerTransactionUpdateUseCase) Update(ctx context.Context, e *model.Transaction, i *model.TransactionUpdateInput, isFeePaidByMe bool) (*model.Transaction, error) {
//		if e.TransactionType == transaction.TransactionTypeInternal {
//			sender, _ := uc.cBaGetter.GetFirst(ctx, nil, &model.BankAccountWhereInput{
//				ID: e.SenderID,
//			})
//			am, _ := i.Amount.Float64()
//			cshOut := sender.CashOut - am
//			_, err := uc.cBaUpdater.Update(ctx, sender, &model.BankAccountUpdateInput{
//				CashOut: &cshOut,
//			})
//			if err != nil {
//				return nil, err
//			}
//			if i.ReceiverID != nil {
//				receiver, _ := uc.cBaGetter.GetFirst(ctx, nil, &model.BankAccountWhereInput{
//					ID: i.ReceiverID,
//				})
//				cshIn := am + receiver.CashIn
//				_, err := uc.cBaUpdater.Update(ctx, receiver, &model.BankAccountUpdateInput{
//					CashIn: &cshIn,
//				})
//				if err != nil {
//					return nil, err
//				}
//			}
//			txc, err := uc.repoUpdate.Update(ctx, e, i)
//			if err != nil {
//				return nil, err
//			}
//			if i.SourceTransactionID == nil {
//				stsSuc := transaction.StatusSuccess
//				desc := FeeDesc
//				ni := &model.TransactionCreateInput{
//					SourceTransactionID:     &txc.ID,
//					SenderID:                i.SenderID,
//					SenderBankAccountNumber: i.SenderBankAccountNumber,
//					SenderBankName:          i.SenderBankName,
//					Amount:                  decimal.NewFromFloat(FeeAmount),
//					Status:                  &stsSuc,
//					TransactionType:         transaction.TransactionTypeInternal,
//					Description:             &desc,
//				}
//				if isFeePaidByMe {
//					ni.SenderID = i.SenderID
//					ni.SenderBankAccountNumber = i.SenderBankAccountNumber
//					ni.SenderBankName = i.SenderBankName
//				} else {
//					ni.SenderID = *i.ReceiverID
//					ni.SenderBankAccountNumber = i.ReceiverBankAccountNumber
//					ni.SenderBankName = i.ReceiverBankName
//				}
//				_, err = uc.createTxcFee(ctx, ni)
//				if err != nil {
//					return nil, err
//				}
//			}
//		}
//	}
func (uc *CustomerTransactionCreateUseCase) Create(ctx context.Context, i *model.TransactionCreateInput, isFeePaidByMe bool) (*model.Transaction, error) {
	return uc.repoCreate.Create(ctx, i)
}

func (uc *CustomerTransactionValidateCreateInputUseCase) Validate(ctx context.Context, i *model.TransactionCreateInput, isFeePaidByMe bool) (*model.TransactionCreateInput, error) {
	ba, err := uc.cBaGetter.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		ID: &i.SenderID,
	})
	if err != nil {
		return nil, err
	}
	if ba == nil {
		return nil, usecase.WrapError(fmt.Errorf("bank account sender is invalid"))
	}
	if !ba.IsForPayment {
		return nil, usecase.WrapError(fmt.Errorf("bank account sender is not for payment"))
	}
	am, _ := i.Amount.Float64()
	if isFeePaidByMe {
		if err = ba.IsBalanceSufficient(am + FeeAmount); err != nil {
			return nil, usecase.WrapError(err)
		}
	} else if err = ba.IsBalanceSufficient(am); err != nil {
		return nil, usecase.WrapError(err)
	}
	if i.TransactionType == transaction.TransactionTypeInternal {
		baOther, err := uc.cBaGetter.GetFirst(ctx, nil, &model.BankAccountWhereInput{
			ID: i.ReceiverID,
		})
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
			if err = baOther.IsBalanceSufficient(FeeAmount); err != nil {
				return nil, usecase.WrapError(err)
			}
		}
		other, err := uc.cGetter.GetFirst(ctx, nil, &model.CustomerWhereInput{
			ID: &baOther.CustomerID,
		})
		if err != nil {
			return nil, err
		}
		i.ReceiverBankAccountNumber = baOther.AccountNumber
		i.ReceiverBankName = ProductOwnerName
		i.ReceiverName = other.GetName()
	}
	stsDrf := transaction.StatusDraft
	i.Status = &stsDrf
	i.SenderBankAccountNumber = ba.AccountNumber
	i.SenderBankName = ProductOwnerName
	uAny := ctx.Value("user")
	user, _ := uAny.(*model.Customer)
	i.SenderName = user.GetName()
	return i, nil
}

func (uc *CustomerTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	return uc.repoList.List(ctx, limit, offset, o, w)
}
