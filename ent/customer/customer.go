// Code generated by ent, DO NOT EDIT.

package customer

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldJwtTokenKey holds the string denoting the jwt_token_key field in the database.
	FieldJwtTokenKey = "jwt_token_key"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldIsActive holds the string denoting the is_active field in the database.
	FieldIsActive = "is_active"
	// EdgeBankAccounts holds the string denoting the bank_accounts edge name in mutations.
	EdgeBankAccounts = "bank_accounts"
	// EdgeContacts holds the string denoting the contacts edge name in mutations.
	EdgeContacts = "contacts"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// BankAccountsTable is the table that holds the bank_accounts relation/edge.
	BankAccountsTable = "bank_accounts"
	// BankAccountsInverseTable is the table name for the BankAccount entity.
	// It exists in this package in order to avoid circular dependency with the "bankaccount" package.
	BankAccountsInverseTable = "bank_accounts"
	// BankAccountsColumn is the table column denoting the bank_accounts relation/edge.
	BankAccountsColumn = "customer_id"
	// ContactsTable is the table that holds the contacts relation/edge.
	ContactsTable = "contacts"
	// ContactsInverseTable is the table name for the Contact entity.
	// It exists in this package in order to avoid circular dependency with the "contact" package.
	ContactsInverseTable = "contacts"
	// ContactsColumn is the table column denoting the contacts relation/edge.
	ContactsColumn = "owner_id"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldJwtTokenKey,
	FieldPassword,
	FieldUsername,
	FieldFirstName,
	FieldLastName,
	FieldPhoneNumber,
	FieldEmail,
	FieldIsActive,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultJwtTokenKey holds the default value on creation for the "jwt_token_key" field.
	DefaultJwtTokenKey func() string
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DefaultFirstName holds the default value on creation for the "first_name" field.
	DefaultFirstName string
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// DefaultLastName holds the default value on creation for the "last_name" field.
	DefaultLastName string
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultIsActive holds the default value on creation for the "is_active" field.
	DefaultIsActive bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
