package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

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
