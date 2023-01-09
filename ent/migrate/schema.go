// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdminsColumns holds the columns for the "admins" table.
	AdminsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "jwt_token_key", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, SchemaType: map[string]string{"mysql": "char(32)"}},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "is_active", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// AdminsTable holds the schema information for the "admins" table.
	AdminsTable = &schema.Table{
		Name:       "admins",
		Columns:    AdminsColumns,
		PrimaryKey: []*schema.Column{AdminsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "admin_create_time",
				Unique:  false,
				Columns: []*schema.Column{AdminsColumns[1]},
			},
		},
	}
	// BankAccountsColumns holds the columns for the "bank_accounts" table.
	BankAccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "cash_in", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(20,2)", "postgres": "numeric"}},
		{Name: "cash_out", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(20,2)", "postgres": "numeric"}},
		{Name: "account_number", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "is_for_payment", Type: field.TypeBool, Default: false},
		{Name: "customer_id", Type: field.TypeUUID},
	}
	// BankAccountsTable holds the schema information for the "bank_accounts" table.
	BankAccountsTable = &schema.Table{
		Name:       "bank_accounts",
		Columns:    BankAccountsColumns,
		PrimaryKey: []*schema.Column{BankAccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bank_accounts_customers_customer",
				Columns:    []*schema.Column{BankAccountsColumns[7]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.Restrict,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "bankaccount_account_number",
				Unique:  false,
				Columns: []*schema.Column{BankAccountsColumns[5]},
			},
			{
				Name:    "bankaccount_customer_id",
				Unique:  false,
				Columns: []*schema.Column{BankAccountsColumns[7]},
			},
			{
				Name:    "bankaccount_create_time",
				Unique:  false,
				Columns: []*schema.Column{BankAccountsColumns[1]},
			},
			{
				Name:    "bankaccount_update_time",
				Unique:  false,
				Columns: []*schema.Column{BankAccountsColumns[2]},
			},
		},
	}
	// ContactsColumns holds the columns for the "contacts" table.
	ContactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "account_number", Type: field.TypeString, Size: 255},
		{Name: "suggest_name", Type: field.TypeString, Size: 255},
		{Name: "bank_name", Type: field.TypeString, Size: 2147483647},
		{Name: "owner_id", Type: field.TypeUUID},
	}
	// ContactsTable holds the schema information for the "contacts" table.
	ContactsTable = &schema.Table{
		Name:       "contacts",
		Columns:    ContactsColumns,
		PrimaryKey: []*schema.Column{ContactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "contacts_customers_owner",
				Columns:    []*schema.Column{ContactsColumns[6]},
				RefColumns: []*schema.Column{CustomersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "contact_suggest_name",
				Unique:  false,
				Columns: []*schema.Column{ContactsColumns[4]},
			},
			{
				Name:    "contact_create_time",
				Unique:  false,
				Columns: []*schema.Column{ContactsColumns[1]},
			},
			{
				Name:    "contact_owner_id",
				Unique:  false,
				Columns: []*schema.Column{ContactsColumns[6]},
			},
			{
				Name:    "contact_account_number_bank_name_owner_id",
				Unique:  true,
				Columns: []*schema.Column{ContactsColumns[3], ContactsColumns[5], ContactsColumns[6]},
			},
		},
	}
	// CustomersColumns holds the columns for the "customers" table.
	CustomersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "jwt_token_key", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, SchemaType: map[string]string{"mysql": "char(32)"}},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "phone_number", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 255},
		{Name: "is_active", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// CustomersTable holds the schema information for the "customers" table.
	CustomersTable = &schema.Table{
		Name:       "customers",
		Columns:    CustomersColumns,
		PrimaryKey: []*schema.Column{CustomersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "customer_create_time",
				Unique:  false,
				Columns: []*schema.Column{CustomersColumns[1]},
			},
		},
	}
	// DebtsColumns holds the columns for the "debts" table.
	DebtsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "owner_bank_account_number", Type: field.TypeString, Size: 255},
		{Name: "owner_bank_name", Type: field.TypeString, Size: 255},
		{Name: "owner_name", Type: field.TypeString, Size: 256},
		{Name: "receiver_bank_account_number", Type: field.TypeString, Size: 255},
		{Name: "receiver_bank_name", Type: field.TypeString, Size: 255},
		{Name: "receiver_name", Type: field.TypeString, Size: 256},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "cancelled", "fulfilled"}, Default: "pending"},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(20,2)", "postgres": "numeric"}},
		{Name: "owner_id", Type: field.TypeUUID},
		{Name: "receiver_id", Type: field.TypeUUID},
		{Name: "transaction_id", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// DebtsTable holds the schema information for the "debts" table.
	DebtsTable = &schema.Table{
		Name:       "debts",
		Columns:    DebtsColumns,
		PrimaryKey: []*schema.Column{DebtsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "debts_bank_accounts_owner",
				Columns:    []*schema.Column{DebtsColumns[12]},
				RefColumns: []*schema.Column{BankAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "debts_bank_accounts_receiver",
				Columns:    []*schema.Column{DebtsColumns[13]},
				RefColumns: []*schema.Column{BankAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "debts_transactions_debt",
				Columns:    []*schema.Column{DebtsColumns[14]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "debt_owner_id",
				Unique:  false,
				Columns: []*schema.Column{DebtsColumns[12]},
			},
			{
				Name:    "debt_receiver_id",
				Unique:  false,
				Columns: []*schema.Column{DebtsColumns[13]},
			},
			{
				Name:    "debt_transaction_id",
				Unique:  false,
				Columns: []*schema.Column{DebtsColumns[14]},
			},
			{
				Name:    "debt_create_time",
				Unique:  false,
				Columns: []*schema.Column{DebtsColumns[1]},
			},
			{
				Name:    "debt_update_time",
				Unique:  false,
				Columns: []*schema.Column{DebtsColumns[2]},
			},
		},
	}
	// EmployeesColumns holds the columns for the "employees" table.
	EmployeesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "jwt_token_key", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 255, SchemaType: map[string]string{"mysql": "char(32)"}},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 128},
		{Name: "first_name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "last_name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "is_active", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// EmployeesTable holds the schema information for the "employees" table.
	EmployeesTable = &schema.Table{
		Name:       "employees",
		Columns:    EmployeesColumns,
		PrimaryKey: []*schema.Column{EmployeesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "employee_create_time",
				Unique:  false,
				Columns: []*schema.Column{EmployeesColumns[1]},
			},
		},
	}
	// TransactionsColumns holds the columns for the "transactions" table.
	TransactionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"draft", "verified", "success"}, Default: "draft"},
		{Name: "receiver_bank_account_number", Type: field.TypeString, Size: 255},
		{Name: "receiver_bank_name", Type: field.TypeString, Size: 255},
		{Name: "receiver_name", Type: field.TypeString, Size: 256},
		{Name: "sender_bank_account_number", Type: field.TypeString, Size: 255},
		{Name: "sender_bank_name", Type: field.TypeString, Size: 255},
		{Name: "sender_name", Type: field.TypeString, Size: 256},
		{Name: "amount", Type: field.TypeFloat64, SchemaType: map[string]string{"mysql": "decimal(20,2)", "postgres": "numeric"}},
		{Name: "transaction_type", Type: field.TypeEnum, Enums: []string{"internal", "external"}},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
		{Name: "source_transaction_id", Type: field.TypeUUID, Unique: true, Nullable: true},
		{Name: "receiver_id", Type: field.TypeUUID, Nullable: true},
		{Name: "sender_id", Type: field.TypeUUID},
	}
	// TransactionsTable holds the schema information for the "transactions" table.
	TransactionsTable = &schema.Table{
		Name:       "transactions",
		Columns:    TransactionsColumns,
		PrimaryKey: []*schema.Column{TransactionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "transactions_transactions_fee_transaction",
				Columns:    []*schema.Column{TransactionsColumns[13]},
				RefColumns: []*schema.Column{TransactionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "transactions_bank_accounts_receiver",
				Columns:    []*schema.Column{TransactionsColumns[14]},
				RefColumns: []*schema.Column{BankAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "transactions_bank_accounts_sender",
				Columns:    []*schema.Column{TransactionsColumns[15]},
				RefColumns: []*schema.Column{BankAccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "transaction_source_transaction_id",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[13]},
			},
			{
				Name:    "transaction_receiver_id",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[14]},
			},
			{
				Name:    "transaction_sender_id",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[15]},
			},
			{
				Name:    "transaction_create_time",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[1]},
			},
			{
				Name:    "transaction_update_time",
				Unique:  false,
				Columns: []*schema.Column{TransactionsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdminsTable,
		BankAccountsTable,
		ContactsTable,
		CustomersTable,
		DebtsTable,
		EmployeesTable,
		TransactionsTable,
	}
)

func init() {
	BankAccountsTable.ForeignKeys[0].RefTable = CustomersTable
	ContactsTable.ForeignKeys[0].RefTable = CustomersTable
	DebtsTable.ForeignKeys[0].RefTable = BankAccountsTable
	DebtsTable.ForeignKeys[1].RefTable = BankAccountsTable
	DebtsTable.ForeignKeys[2].RefTable = TransactionsTable
	TransactionsTable.ForeignKeys[0].RefTable = TransactionsTable
	TransactionsTable.ForeignKeys[1].RefTable = BankAccountsTable
	TransactionsTable.ForeignKeys[2].RefTable = BankAccountsTable
}
