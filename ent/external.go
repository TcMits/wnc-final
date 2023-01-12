package ent

import (
	"context"
)

func (c *Client) Flush(ctx context.Context) error {
	var err error
	_, err = c.Debt.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.Transaction.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.BankAccount.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.Contact.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.Customer.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.Employee.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.Admin.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	_, err = c.Partner.Delete().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
