package ent

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/bankaccount"
)

func (s *BankAccount) GetBalance() float64 {
	return s.CashIn - s.CashOut
}

func (s *BankAccount) IsBalanceSufficient(amount float64) (bool, error) {
	bl := s.GetBalance()
	if amount > bl {
		return false, nil
	}
	return true, nil
}

func (s *Customer) GetName() string {
	return NormalizeName(s.FirstName, s.LastName)
}

func NormalizeName(n ...string) string {
	var name string
	for idx, e := range n {
		if idx > 0 {
			name += " "
		}
		name += e
	}
	return name
}

func refreshFromDB(ctx context.Context, c *Client, eAny any) (any, error) {
	var result any
	var err error
	switch eAny.(type) {
	case *BankAccount:
		e, _ := eAny.(*BankAccount)
		result, err = c.BankAccount.Query().Where(bankaccount.ID(e.ID)).First(ctx)
		if err != nil {
			return nil, err
		}
	default:
		result = eAny
	}
	return result, nil
}

// RefreshBankAccountFromDB refresh state of bank account entity
func RefreshBankAccountFromDB(ctx context.Context, c *Client, e *BankAccount) (*BankAccount, error) {
	eAny, err := refreshFromDB(ctx, c, e)
	if err != nil {
		return nil, err
	}
	e, ok := eAny.(*BankAccount)
	if !ok {
		return nil, fmt.Errorf("invalid type")
	}
	return e, nil
}
