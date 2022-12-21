package ent

import "fmt"

func (s *BankAccount) GetBalance() float64 {
	return s.CashIn - s.CashOut
}

func (s *BankAccount) IsBalanceSufficient(amount float64) error {
	bl := s.GetBalance()
	if amount > bl {
		return fmt.Errorf("insufficient balance")
	}
	return nil
}

func (s *Customer) GetName() string {
	return s.FirstName + " " + s.LastName
}
