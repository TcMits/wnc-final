package repository

import (
	"context"

	"github.com/TcMits/wnc-final/ent"
	entTxc "github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/transaction"
)

type TransactionConfirmSuccessRepository struct {
	client *ent.Client
}
type PartnerTransactionCreateRepository struct {
	client *ent.Client
}
type partnerTransactionCreateRepository struct {
	r1 CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
	r2 UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	r3 ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
}

type transactionConfirmSuccessRepository struct {
	bAUR UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	bALR ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	tUR  UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput]
	tCR  CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
}

func GetTransactionListRepository(
	client *ent.Client,
) ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return ent.NewTransactionReadRepository(client)
}

func GetTransactionCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Transaction, *model.TransactionCreateInput] {
	return ent.NewTransactionCreateRepository(client, false)
}

func GetTransactionUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput] {
	return ent.NewTransactionUpdateRepository(client, false)
}

func GetTransactionConfirmSuccessRepository(
	client *ent.Client,
) ITransactionConfirmSuccessRepository {
	return &TransactionConfirmSuccessRepository{
		client: client,
	}
}
func GetTransactionIsNextRepository(
	client *ent.Client,
) IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return getIsNextModelRepostiory(GetTransactionListRepository(client))
}
func GetPartnerTransactionCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Transaction, *model.TransactionCreateInput] {
	return &PartnerTransactionCreateRepository{
		client: client,
	}
}

func newTransactionConfirmSuccessRepository(
	client *ent.Client,
) *transactionConfirmSuccessRepository {
	return &transactionConfirmSuccessRepository{
		bAUR: GetBankAccountUpdateRepository(client),
		bALR: GetBankAccountListRepository(client),
		tUR:  GetTransactionUpdateRepository(client),
		tCR:  GetTransactionCreateRepository(client),
	}
}
func newPartnerTransactionConfirmSuccessRepository(
	client *ent.Client,
) *partnerTransactionCreateRepository {
	return &partnerTransactionCreateRepository{
		r1: GetPartnerTransactionCreateRepository(client),
		r2: GetBankAccountUpdateRepository(client),
		r3: GetBankAccountListRepository(client),
	}
}

func (s *transactionConfirmSuccessRepository) subtractSenderBankAccount(ctx context.Context, txc *model.Transaction) (*model.BankAccount, error) {
	bAs, _ := s.bALR.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.BankAccountWhereInput{
		ID: txc.SenderID,
	})
	bA := bAs[0]
	aM, _ := txc.Amount.Float64()
	bA, err := s.bAUR.Update(ctx, bA, &model.BankAccountUpdateInput{
		CashOut: generic.GetPointer(aM),
	})
	if err != nil {
		return nil, err
	}
	return bA, nil
}
func (s *transactionConfirmSuccessRepository) addReceiverBankAccount(ctx context.Context, txc *model.Transaction) (*model.BankAccount, error) {
	bAs, _ := s.bALR.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.BankAccountWhereInput{
		ID: txc.ReceiverID,
	})
	bA := bAs[0]
	aM, _ := txc.Amount.Float64()
	bA, err := s.bAUR.Update(ctx, bA, &model.BankAccountUpdateInput{
		CashIn: generic.GetPointer(aM),
	})
	if err != nil {
		return nil, err
	}
	return bA, nil
}

func (s *transactionConfirmSuccessRepository) confirm(ctx context.Context, e *model.Transaction, feeTxcInput *model.TransactionCreateInput) (*model.Transaction, error) {
	_, err := s.subtractSenderBankAccount(ctx, e)
	if err != nil {
		return nil, err
	}
	_, err = s.addReceiverBankAccount(ctx, e)
	if err != nil {
		return nil, err
	}
	txc, err := s.tUR.Update(ctx, e, &model.TransactionUpdateInput{
		Status: generic.GetPointer(entTxc.StatusSuccess),
	})
	if err != nil {
		return nil, err
	}
	_, err = s.tCR.Create(ctx, feeTxcInput)
	if err != nil {
		return nil, err
	}
	return txc, nil
}
func (s *partnerTransactionCreateRepository) addReceiverBankAccount(ctx context.Context, txc *model.Transaction) (*model.BankAccount, error) {
	bAs, _ := s.r3.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.BankAccountWhereInput{
		ID: txc.ReceiverID,
	})
	bA := bAs[0]
	aM, _ := txc.Amount.Float64()
	bA, err := s.r2.Update(ctx, bA, &model.BankAccountUpdateInput{
		CashIn: generic.GetPointer(aM),
	})
	if err != nil {
		return nil, err
	}
	return bA, nil
}
func (s *partnerTransactionCreateRepository) create(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	txc, err := s.r1.Create(ctx, i)
	if err != nil {
		return nil, err
	}
	_, err = s.addReceiverBankAccount(ctx, txc)
	if err != nil {
		return nil, err
	}
	return txc, nil
}

func (s *TransactionConfirmSuccessRepository) ConfirmSuccess(ctx context.Context, e *model.Transaction, f *model.TransactionCreateInput) (*model.Transaction, error) {
	if err := transaction.WithTx(ctx, s.client, func(tx *ent.Tx) error {
		var errInner error
		confirmer := newTransactionConfirmSuccessRepository(tx.Client())
		e, errInner = confirmer.confirm(ctx, e, f)
		return errInner
	}); err != nil {
		return nil, err
	}
	return e, nil
}

func (s *PartnerTransactionCreateRepository) Create(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	var e *model.Transaction
	if err := transaction.WithTx(ctx, s.client, func(tx *ent.Tx) error {
		var errInner error
		driver := newPartnerTransactionConfirmSuccessRepository(tx.Client())
		e, errInner = driver.create(ctx, i)
		return errInner
	}); err != nil {
		return nil, err
	}
	return e, nil
}
