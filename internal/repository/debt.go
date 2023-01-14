package repository

import (
	"context"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/debt"
	entTxc "github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/transaction"
)

type (
	DebtFullfillRepository struct {
		client *ent.Client
	}
	debtFullfillRepository struct {
		dUR  UpdateModelRepository[*model.Debt, *model.DebtUpdateInput]
		bAUR UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
		bALR ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
		tCR  CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
	}
)

func GetDebtIsNextRepository(
	client *ent.Client,
) IIsNextModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput] {
	return getIsNextModelRepostiory(GetDebtListRepository(client))
}

func GetDebtListRepository(
	client *ent.Client,
) ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput] {
	return ent.NewDebtReadRepository(client)
}

func GetDebtCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Debt, *model.DebtCreateInput] {
	return ent.NewDebtCreateRepository(client, false)
}

func GetDebtUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Debt, *model.DebtUpdateInput] {
	return ent.NewDebtUpdateRepository(client, false)
}

func GetDebtDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.Debt] {
	return ent.NewDebtDeleteRepository(client, false)
}

func GetDebtFulfillRepository(
	client *ent.Client,
) IDebtFullfillRepository {
	return &DebtFullfillRepository{client: client}
}

func newDebtFullfillRepository(
	client *ent.Client,
) *debtFullfillRepository {
	return &debtFullfillRepository{
		dUR:  GetDebtUpdateRepository(client),
		bALR: GetBankAccountListRepository(client),
		bAUR: GetBankAccountUpdateRepository(client),
		tCR:  GetTransactionCreateRepository(client),
	}
}

func (s *debtFullfillRepository) subtractReceiverBankAccount(ctx context.Context, e *model.Debt) (*model.BankAccount, error) {
	bAs, _ := s.bALR.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.BankAccountWhereInput{
		ID: e.ReceiverID,
	})
	bA := bAs[0]
	aM, _ := e.Amount.Float64()
	bA, err := s.bAUR.Update(ctx, bA, &model.BankAccountUpdateInput{
		CashOut: generic.GetPointer(aM),
	})
	if err != nil {
		return nil, err
	}
	return bA, nil
}
func (s *debtFullfillRepository) addOwnerBankAccount(ctx context.Context, e *model.Debt) (*model.BankAccount, error) {
	bAs, _ := s.bALR.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.BankAccountWhereInput{
		ID: e.OwnerID,
	})
	bA := bAs[0]
	aM, _ := e.Amount.Float64()
	bA, err := s.bAUR.Update(ctx, bA, &model.BankAccountUpdateInput{
		CashIn: generic.GetPointer(aM),
	})
	if err != nil {
		return nil, err
	}
	return bA, nil
}
func (s *debtFullfillRepository) createTransaction(ctx context.Context, e *model.Debt) (*model.Transaction, error) {
	i := &model.TransactionCreateInput{
		SenderID:                  e.ReceiverID,
		SenderBankAccountNumber:   e.ReceiverBankAccountNumber,
		SenderName:                e.ReceiverName,
		SenderBankName:            e.ReceiverBankName,
		ReceiverID:                e.OwnerID,
		ReceiverBankAccountNumber: e.OwnerBankAccountNumber,
		ReceiverBankName:          e.OwnerBankName,
		ReceiverName:              e.OwnerName,
		Status:                    generic.GetPointer(entTxc.StatusSuccess),
		Amount:                    e.Amount,
		TransactionType:           entTxc.TransactionTypeInternal,
		Description:               &e.Description,
	}
	res, err := s.tCR.Create(ctx, i)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *debtFullfillRepository) updateStatus(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	if i == nil {
		i = new(model.DebtUpdateInput)
	}
	i.Status = generic.GetPointer(debt.StatusFulfilled)
	return s.dUR.Update(ctx, e, i)
}

func (s *debtFullfillRepository) fulfill(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	_, err := s.subtractReceiverBankAccount(ctx, e)
	if err != nil {
		return nil, err
	}
	_, err = s.addOwnerBankAccount(ctx, e)
	if err != nil {
		return nil, err
	}
	txc, err := s.createTransaction(ctx, e)
	if err != nil {
		return nil, err
	}
	if i == nil {
		i = new(model.DebtUpdateInput)
	}
	i.TransactionID = generic.GetPointer(txc.ID)
	e, err = s.updateStatus(ctx, e, i)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (s *DebtFullfillRepository) Fulfill(ctx context.Context, e *model.Debt, i *model.DebtUpdateInput) (*model.Debt, error) {
	if err := transaction.WithTx(ctx, s.client, func(tx *ent.Tx) error {
		var errInner error
		driver := newDebtFullfillRepository(tx.Client())
		e, errInner = driver.fulfill(ctx, e, i)
		return errInner
	}); err != nil {
		return nil, err
	}
	return e, nil
}
