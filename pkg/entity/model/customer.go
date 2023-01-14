package model

import "github.com/TcMits/wnc-final/ent"

type (
	Customer                   = ent.Customer
	CustomerOrderInput         = ent.CustomerOrderInput
	CustomerWhereInput         = ent.CustomerWhereInput
	CustomerCreateInput        = ent.CustomerCreateInput
	CustomerUpdateInput        = ent.CustomerUpdateInput
	EmployeeCreateCustomerResp struct {
		*Customer
		*BankAccount
	}
)
