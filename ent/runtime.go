// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/TcMits/wnc-final/ent/admin"
	"github.com/TcMits/wnc-final/ent/bankaccount"
	"github.com/TcMits/wnc-final/ent/contact"
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/employee"
	"github.com/TcMits/wnc-final/ent/schema"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	adminMixin := schema.Admin{}.Mixin()
	adminMixinFields0 := adminMixin[0].Fields()
	_ = adminMixinFields0
	adminFields := schema.Admin{}.Fields()
	_ = adminFields
	// adminDescCreateTime is the schema descriptor for create_time field.
	adminDescCreateTime := adminMixinFields0[0].Descriptor()
	// admin.DefaultCreateTime holds the default value on creation for the create_time field.
	admin.DefaultCreateTime = adminDescCreateTime.Default.(func() time.Time)
	// adminDescUpdateTime is the schema descriptor for update_time field.
	adminDescUpdateTime := adminMixinFields0[1].Descriptor()
	// admin.DefaultUpdateTime holds the default value on creation for the update_time field.
	admin.DefaultUpdateTime = adminDescUpdateTime.Default.(func() time.Time)
	// admin.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	admin.UpdateDefaultUpdateTime = adminDescUpdateTime.UpdateDefault.(func() time.Time)
	// adminDescJwtTokenKey is the schema descriptor for jwt_token_key field.
	adminDescJwtTokenKey := adminFields[1].Descriptor()
	// admin.DefaultJwtTokenKey holds the default value on creation for the jwt_token_key field.
	admin.DefaultJwtTokenKey = adminDescJwtTokenKey.Default.(func() string)
	// adminDescPassword is the schema descriptor for password field.
	adminDescPassword := adminFields[2].Descriptor()
	// admin.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	admin.PasswordValidator = adminDescPassword.Validators[0].(func(string) error)
	// adminDescUsername is the schema descriptor for username field.
	adminDescUsername := adminFields[3].Descriptor()
	// admin.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	admin.UsernameValidator = func() func(string) error {
		validators := adminDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// adminDescFirstName is the schema descriptor for first_name field.
	adminDescFirstName := adminFields[4].Descriptor()
	// admin.DefaultFirstName holds the default value on creation for the first_name field.
	admin.DefaultFirstName = adminDescFirstName.Default.(string)
	// admin.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	admin.FirstNameValidator = adminDescFirstName.Validators[0].(func(string) error)
	// adminDescLastName is the schema descriptor for last_name field.
	adminDescLastName := adminFields[5].Descriptor()
	// admin.DefaultLastName holds the default value on creation for the last_name field.
	admin.DefaultLastName = adminDescLastName.Default.(string)
	// admin.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	admin.LastNameValidator = adminDescLastName.Validators[0].(func(string) error)
	// adminDescIsActive is the schema descriptor for is_active field.
	adminDescIsActive := adminFields[6].Descriptor()
	// admin.DefaultIsActive holds the default value on creation for the is_active field.
	admin.DefaultIsActive = adminDescIsActive.Default.(bool)
	// adminDescID is the schema descriptor for id field.
	adminDescID := adminFields[0].Descriptor()
	// admin.DefaultID holds the default value on creation for the id field.
	admin.DefaultID = adminDescID.Default.(func() uuid.UUID)
	bankaccountMixin := schema.BankAccount{}.Mixin()
	bankaccountMixinFields0 := bankaccountMixin[0].Fields()
	_ = bankaccountMixinFields0
	bankaccountFields := schema.BankAccount{}.Fields()
	_ = bankaccountFields
	// bankaccountDescCreateTime is the schema descriptor for create_time field.
	bankaccountDescCreateTime := bankaccountMixinFields0[0].Descriptor()
	// bankaccount.DefaultCreateTime holds the default value on creation for the create_time field.
	bankaccount.DefaultCreateTime = bankaccountDescCreateTime.Default.(func() time.Time)
	// bankaccountDescUpdateTime is the schema descriptor for update_time field.
	bankaccountDescUpdateTime := bankaccountMixinFields0[1].Descriptor()
	// bankaccount.DefaultUpdateTime holds the default value on creation for the update_time field.
	bankaccount.DefaultUpdateTime = bankaccountDescUpdateTime.Default.(func() time.Time)
	// bankaccount.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	bankaccount.UpdateDefaultUpdateTime = bankaccountDescUpdateTime.UpdateDefault.(func() time.Time)
	// bankaccountDescCashIn is the schema descriptor for cash_in field.
	bankaccountDescCashIn := bankaccountFields[2].Descriptor()
	// bankaccount.CashInValidator is a validator for the "cash_in" field. It is called by the builders before save.
	bankaccount.CashInValidator = bankaccountDescCashIn.Validators[0].(func(float64) error)
	// bankaccountDescCashOut is the schema descriptor for cash_out field.
	bankaccountDescCashOut := bankaccountFields[3].Descriptor()
	// bankaccount.CashOutValidator is a validator for the "cash_out" field. It is called by the builders before save.
	bankaccount.CashOutValidator = bankaccountDescCashOut.Validators[0].(func(float64) error)
	// bankaccountDescAccountNumber is the schema descriptor for account_number field.
	bankaccountDescAccountNumber := bankaccountFields[4].Descriptor()
	// bankaccount.DefaultAccountNumber holds the default value on creation for the account_number field.
	bankaccount.DefaultAccountNumber = bankaccountDescAccountNumber.Default.(func() string)
	// bankaccount.AccountNumberValidator is a validator for the "account_number" field. It is called by the builders before save.
	bankaccount.AccountNumberValidator = bankaccountDescAccountNumber.Validators[0].(func(string) error)
	// bankaccountDescIsForPayment is the schema descriptor for is_for_payment field.
	bankaccountDescIsForPayment := bankaccountFields[5].Descriptor()
	// bankaccount.DefaultIsForPayment holds the default value on creation for the is_for_payment field.
	bankaccount.DefaultIsForPayment = bankaccountDescIsForPayment.Default.(bool)
	// bankaccountDescID is the schema descriptor for id field.
	bankaccountDescID := bankaccountFields[0].Descriptor()
	// bankaccount.DefaultID holds the default value on creation for the id field.
	bankaccount.DefaultID = bankaccountDescID.Default.(func() uuid.UUID)
	contactMixin := schema.Contact{}.Mixin()
	contactMixinFields0 := contactMixin[0].Fields()
	_ = contactMixinFields0
	contactFields := schema.Contact{}.Fields()
	_ = contactFields
	// contactDescCreateTime is the schema descriptor for create_time field.
	contactDescCreateTime := contactMixinFields0[0].Descriptor()
	// contact.DefaultCreateTime holds the default value on creation for the create_time field.
	contact.DefaultCreateTime = contactDescCreateTime.Default.(func() time.Time)
	// contactDescUpdateTime is the schema descriptor for update_time field.
	contactDescUpdateTime := contactMixinFields0[1].Descriptor()
	// contact.DefaultUpdateTime holds the default value on creation for the update_time field.
	contact.DefaultUpdateTime = contactDescUpdateTime.Default.(func() time.Time)
	// contact.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	contact.UpdateDefaultUpdateTime = contactDescUpdateTime.UpdateDefault.(func() time.Time)
	// contactDescAccountNumber is the schema descriptor for account_number field.
	contactDescAccountNumber := contactFields[2].Descriptor()
	// contact.AccountNumberValidator is a validator for the "account_number" field. It is called by the builders before save.
	contact.AccountNumberValidator = contactDescAccountNumber.Validators[0].(func(string) error)
	// contactDescSuggestName is the schema descriptor for suggest_name field.
	contactDescSuggestName := contactFields[3].Descriptor()
	// contact.SuggestNameValidator is a validator for the "suggest_name" field. It is called by the builders before save.
	contact.SuggestNameValidator = contactDescSuggestName.Validators[0].(func(string) error)
	// contactDescID is the schema descriptor for id field.
	contactDescID := contactFields[0].Descriptor()
	// contact.DefaultID holds the default value on creation for the id field.
	contact.DefaultID = contactDescID.Default.(func() uuid.UUID)
	customerMixin := schema.Customer{}.Mixin()
	customerMixinFields0 := customerMixin[0].Fields()
	_ = customerMixinFields0
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescCreateTime is the schema descriptor for create_time field.
	customerDescCreateTime := customerMixinFields0[0].Descriptor()
	// customer.DefaultCreateTime holds the default value on creation for the create_time field.
	customer.DefaultCreateTime = customerDescCreateTime.Default.(func() time.Time)
	// customerDescUpdateTime is the schema descriptor for update_time field.
	customerDescUpdateTime := customerMixinFields0[1].Descriptor()
	// customer.DefaultUpdateTime holds the default value on creation for the update_time field.
	customer.DefaultUpdateTime = customerDescUpdateTime.Default.(func() time.Time)
	// customer.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	customer.UpdateDefaultUpdateTime = customerDescUpdateTime.UpdateDefault.(func() time.Time)
	// customerDescJwtTokenKey is the schema descriptor for jwt_token_key field.
	customerDescJwtTokenKey := customerFields[1].Descriptor()
	// customer.DefaultJwtTokenKey holds the default value on creation for the jwt_token_key field.
	customer.DefaultJwtTokenKey = customerDescJwtTokenKey.Default.(func() string)
	// customerDescPassword is the schema descriptor for password field.
	customerDescPassword := customerFields[2].Descriptor()
	// customer.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	customer.PasswordValidator = customerDescPassword.Validators[0].(func(string) error)
	// customerDescUsername is the schema descriptor for username field.
	customerDescUsername := customerFields[3].Descriptor()
	// customer.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	customer.UsernameValidator = func() func(string) error {
		validators := customerDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// customerDescFirstName is the schema descriptor for first_name field.
	customerDescFirstName := customerFields[4].Descriptor()
	// customer.DefaultFirstName holds the default value on creation for the first_name field.
	customer.DefaultFirstName = customerDescFirstName.Default.(string)
	// customer.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	customer.FirstNameValidator = customerDescFirstName.Validators[0].(func(string) error)
	// customerDescLastName is the schema descriptor for last_name field.
	customerDescLastName := customerFields[5].Descriptor()
	// customer.DefaultLastName holds the default value on creation for the last_name field.
	customer.DefaultLastName = customerDescLastName.Default.(string)
	// customer.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	customer.LastNameValidator = customerDescLastName.Validators[0].(func(string) error)
	// customerDescPhoneNumber is the schema descriptor for phone_number field.
	customerDescPhoneNumber := customerFields[6].Descriptor()
	// customer.PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	customer.PhoneNumberValidator = customerDescPhoneNumber.Validators[0].(func(string) error)
	// customerDescEmail is the schema descriptor for email field.
	customerDescEmail := customerFields[7].Descriptor()
	// customer.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	customer.EmailValidator = func() func(string) error {
		validators := customerDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// customerDescIsActive is the schema descriptor for is_active field.
	customerDescIsActive := customerFields[8].Descriptor()
	// customer.DefaultIsActive holds the default value on creation for the is_active field.
	customer.DefaultIsActive = customerDescIsActive.Default.(bool)
	// customerDescID is the schema descriptor for id field.
	customerDescID := customerFields[0].Descriptor()
	// customer.DefaultID holds the default value on creation for the id field.
	customer.DefaultID = customerDescID.Default.(func() uuid.UUID)
	debtMixin := schema.Debt{}.Mixin()
	debtMixinFields0 := debtMixin[0].Fields()
	_ = debtMixinFields0
	debtFields := schema.Debt{}.Fields()
	_ = debtFields
	// debtDescCreateTime is the schema descriptor for create_time field.
	debtDescCreateTime := debtMixinFields0[0].Descriptor()
	// debt.DefaultCreateTime holds the default value on creation for the create_time field.
	debt.DefaultCreateTime = debtDescCreateTime.Default.(func() time.Time)
	// debtDescUpdateTime is the schema descriptor for update_time field.
	debtDescUpdateTime := debtMixinFields0[1].Descriptor()
	// debt.DefaultUpdateTime holds the default value on creation for the update_time field.
	debt.DefaultUpdateTime = debtDescUpdateTime.Default.(func() time.Time)
	// debt.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	debt.UpdateDefaultUpdateTime = debtDescUpdateTime.UpdateDefault.(func() time.Time)
	// debtDescOwnerBankAccountNumber is the schema descriptor for owner_bank_account_number field.
	debtDescOwnerBankAccountNumber := debtFields[1].Descriptor()
	// debt.OwnerBankAccountNumberValidator is a validator for the "owner_bank_account_number" field. It is called by the builders before save.
	debt.OwnerBankAccountNumberValidator = debtDescOwnerBankAccountNumber.Validators[0].(func(string) error)
	// debtDescOwnerBankName is the schema descriptor for owner_bank_name field.
	debtDescOwnerBankName := debtFields[2].Descriptor()
	// debt.OwnerBankNameValidator is a validator for the "owner_bank_name" field. It is called by the builders before save.
	debt.OwnerBankNameValidator = debtDescOwnerBankName.Validators[0].(func(string) error)
	// debtDescOwnerName is the schema descriptor for owner_name field.
	debtDescOwnerName := debtFields[3].Descriptor()
	// debt.OwnerNameValidator is a validator for the "owner_name" field. It is called by the builders before save.
	debt.OwnerNameValidator = debtDescOwnerName.Validators[0].(func(string) error)
	// debtDescReceiverBankAccountNumber is the schema descriptor for receiver_bank_account_number field.
	debtDescReceiverBankAccountNumber := debtFields[5].Descriptor()
	// debt.ReceiverBankAccountNumberValidator is a validator for the "receiver_bank_account_number" field. It is called by the builders before save.
	debt.ReceiverBankAccountNumberValidator = debtDescReceiverBankAccountNumber.Validators[0].(func(string) error)
	// debtDescReceiverBankName is the schema descriptor for receiver_bank_name field.
	debtDescReceiverBankName := debtFields[6].Descriptor()
	// debt.ReceiverBankNameValidator is a validator for the "receiver_bank_name" field. It is called by the builders before save.
	debt.ReceiverBankNameValidator = debtDescReceiverBankName.Validators[0].(func(string) error)
	// debtDescReceiverName is the schema descriptor for receiver_name field.
	debtDescReceiverName := debtFields[7].Descriptor()
	// debt.ReceiverNameValidator is a validator for the "receiver_name" field. It is called by the builders before save.
	debt.ReceiverNameValidator = debtDescReceiverName.Validators[0].(func(string) error)
	// debtDescDescription is the schema descriptor for description field.
	debtDescDescription := debtFields[11].Descriptor()
	// debt.DefaultDescription holds the default value on creation for the description field.
	debt.DefaultDescription = debtDescDescription.Default.(string)
	// debtDescID is the schema descriptor for id field.
	debtDescID := debtFields[0].Descriptor()
	// debt.DefaultID holds the default value on creation for the id field.
	debt.DefaultID = debtDescID.Default.(func() uuid.UUID)
	employeeMixin := schema.Employee{}.Mixin()
	employeeMixinFields0 := employeeMixin[0].Fields()
	_ = employeeMixinFields0
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescCreateTime is the schema descriptor for create_time field.
	employeeDescCreateTime := employeeMixinFields0[0].Descriptor()
	// employee.DefaultCreateTime holds the default value on creation for the create_time field.
	employee.DefaultCreateTime = employeeDescCreateTime.Default.(func() time.Time)
	// employeeDescUpdateTime is the schema descriptor for update_time field.
	employeeDescUpdateTime := employeeMixinFields0[1].Descriptor()
	// employee.DefaultUpdateTime holds the default value on creation for the update_time field.
	employee.DefaultUpdateTime = employeeDescUpdateTime.Default.(func() time.Time)
	// employee.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	employee.UpdateDefaultUpdateTime = employeeDescUpdateTime.UpdateDefault.(func() time.Time)
	// employeeDescJwtTokenKey is the schema descriptor for jwt_token_key field.
	employeeDescJwtTokenKey := employeeFields[1].Descriptor()
	// employee.DefaultJwtTokenKey holds the default value on creation for the jwt_token_key field.
	employee.DefaultJwtTokenKey = employeeDescJwtTokenKey.Default.(func() string)
	// employeeDescPassword is the schema descriptor for password field.
	employeeDescPassword := employeeFields[2].Descriptor()
	// employee.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	employee.PasswordValidator = employeeDescPassword.Validators[0].(func(string) error)
	// employeeDescUsername is the schema descriptor for username field.
	employeeDescUsername := employeeFields[3].Descriptor()
	// employee.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	employee.UsernameValidator = func() func(string) error {
		validators := employeeDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// employeeDescFirstName is the schema descriptor for first_name field.
	employeeDescFirstName := employeeFields[4].Descriptor()
	// employee.DefaultFirstName holds the default value on creation for the first_name field.
	employee.DefaultFirstName = employeeDescFirstName.Default.(string)
	// employee.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	employee.FirstNameValidator = employeeDescFirstName.Validators[0].(func(string) error)
	// employeeDescLastName is the schema descriptor for last_name field.
	employeeDescLastName := employeeFields[5].Descriptor()
	// employee.DefaultLastName holds the default value on creation for the last_name field.
	employee.DefaultLastName = employeeDescLastName.Default.(string)
	// employee.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	employee.LastNameValidator = employeeDescLastName.Validators[0].(func(string) error)
	// employeeDescIsActive is the schema descriptor for is_active field.
	employeeDescIsActive := employeeFields[6].Descriptor()
	// employee.DefaultIsActive holds the default value on creation for the is_active field.
	employee.DefaultIsActive = employeeDescIsActive.Default.(bool)
	// employeeDescID is the schema descriptor for id field.
	employeeDescID := employeeFields[0].Descriptor()
	// employee.DefaultID holds the default value on creation for the id field.
	employee.DefaultID = employeeDescID.Default.(func() uuid.UUID)
	transactionMixin := schema.Transaction{}.Mixin()
	transactionMixinFields0 := transactionMixin[0].Fields()
	_ = transactionMixinFields0
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescCreateTime is the schema descriptor for create_time field.
	transactionDescCreateTime := transactionMixinFields0[0].Descriptor()
	// transaction.DefaultCreateTime holds the default value on creation for the create_time field.
	transaction.DefaultCreateTime = transactionDescCreateTime.Default.(func() time.Time)
	// transactionDescUpdateTime is the schema descriptor for update_time field.
	transactionDescUpdateTime := transactionMixinFields0[1].Descriptor()
	// transaction.DefaultUpdateTime holds the default value on creation for the update_time field.
	transaction.DefaultUpdateTime = transactionDescUpdateTime.Default.(func() time.Time)
	// transaction.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	transaction.UpdateDefaultUpdateTime = transactionDescUpdateTime.UpdateDefault.(func() time.Time)
	// transactionDescReceiverBankAccountNumber is the schema descriptor for receiver_bank_account_number field.
	transactionDescReceiverBankAccountNumber := transactionFields[3].Descriptor()
	// transaction.ReceiverBankAccountNumberValidator is a validator for the "receiver_bank_account_number" field. It is called by the builders before save.
	transaction.ReceiverBankAccountNumberValidator = transactionDescReceiverBankAccountNumber.Validators[0].(func(string) error)
	// transactionDescReceiverBankName is the schema descriptor for receiver_bank_name field.
	transactionDescReceiverBankName := transactionFields[4].Descriptor()
	// transaction.ReceiverBankNameValidator is a validator for the "receiver_bank_name" field. It is called by the builders before save.
	transaction.ReceiverBankNameValidator = transactionDescReceiverBankName.Validators[0].(func(string) error)
	// transactionDescReceiverName is the schema descriptor for receiver_name field.
	transactionDescReceiverName := transactionFields[5].Descriptor()
	// transaction.ReceiverNameValidator is a validator for the "receiver_name" field. It is called by the builders before save.
	transaction.ReceiverNameValidator = transactionDescReceiverName.Validators[0].(func(string) error)
	// transactionDescSenderBankAccountNumber is the schema descriptor for sender_bank_account_number field.
	transactionDescSenderBankAccountNumber := transactionFields[7].Descriptor()
	// transaction.SenderBankAccountNumberValidator is a validator for the "sender_bank_account_number" field. It is called by the builders before save.
	transaction.SenderBankAccountNumberValidator = transactionDescSenderBankAccountNumber.Validators[0].(func(string) error)
	// transactionDescSenderBankName is the schema descriptor for sender_bank_name field.
	transactionDescSenderBankName := transactionFields[8].Descriptor()
	// transaction.SenderBankNameValidator is a validator for the "sender_bank_name" field. It is called by the builders before save.
	transaction.SenderBankNameValidator = transactionDescSenderBankName.Validators[0].(func(string) error)
	// transactionDescSenderName is the schema descriptor for sender_name field.
	transactionDescSenderName := transactionFields[9].Descriptor()
	// transaction.SenderNameValidator is a validator for the "sender_name" field. It is called by the builders before save.
	transaction.SenderNameValidator = transactionDescSenderName.Validators[0].(func(string) error)
	// transactionDescDescription is the schema descriptor for description field.
	transactionDescDescription := transactionFields[13].Descriptor()
	// transaction.DefaultDescription holds the default value on creation for the description field.
	transaction.DefaultDescription = transactionDescDescription.Default.(string)
	// transactionDescID is the schema descriptor for id field.
	transactionDescID := transactionFields[0].Descriptor()
	// transaction.DefaultID holds the default value on creation for the id field.
	transaction.DefaultID = transactionDescID.Default.(func() uuid.UUID)
}
