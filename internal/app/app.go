// Package app configures and runs application.
package app

import (
	"fmt"

	"github.com/TcMits/wnc-final/cmd/createuser/createuser"
	"github.com/TcMits/wnc-final/cmd/gendata/gendata"
	"github.com/TcMits/wnc-final/cmd/migrate/migrate"
	"github.com/TcMits/wnc-final/config"
	v1 "github.com/TcMits/wnc-final/internal/controller/http/v1"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase/constructors"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// setup
	if !cfg.App.Debug {
		migrate.Migrate()
		createuser.CreateUser()
		gendata.GenData()
	}

	// Repository
	client, err := datastore.NewClient(cfg.DB.URL, cfg.DB.PoolMax, cfg.App.Debug)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - NewClient: %w", err))
	}
	defer client.Close()

	// HTTP Server
	handler := v1.NewHandler()

	// Broker Server sent event
	b := sse.NewBroker(l)

	// Customer Usecase
	cUC1 := constructors.NewCustomerUseCase(
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	CMeUc := constructors.NewCustomerMeUseCase(
		repository.GetCustomerListRepository(client),
		repository.GetCustomerUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	CAuthUc := constructors.NewCustomerAuthUseCase(
		task.GetEmailTaskExecutor(cfg.Mail.Host, cfg.Mail.User, cfg.Mail.Password, cfg.Mail.SenderName, cfg.Mail.Port, l),
		repository.GetCustomerListRepository(client),
		repository.GetCustomerUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.Mail.ConfirmEmailSubject,
		&cfg.Mail.ConfirmEmailTemplate,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.TransactionUseCase.FeeAmount,
		cfg.Mail.OTPTimeout,
		cfg.AuthUseCase.AccessTTL,
		cfg.AuthUseCase.RefreshTTL,
	)
	cBankAccountUc := constructors.NewCustomerBankAccountUseCase(
		repository.GetBankAccountUpdateRepository(client),
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountIsNextRepository(client),
		repository.GetBankAccountDeleteRepository(client),
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		cfg.TransactionUseCase.Layout,
		cfg.BaseURL,
		cfg.AuthAPI,
		cfg.BankAccountAPI,
		cfg.ValidateAPI,
		cfg.CreateTransactionAPI,
		cfg.TPBank.Name,
		cfg.TPBank.ApiKey,
		cfg.TPBank.SecretKey,
		cfg.TPBank.PrivateKey,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	cTxcUc := constructors.NewCustomerTransactionUseCase(
		task.GetEmailTaskExecutor(cfg.Mail.Host, cfg.Mail.User, cfg.Mail.Password, cfg.Mail.SenderName, cfg.Mail.Port, l),
		repository.GetTransactionConfirmSuccessRepository(client, cfg.TransactionUseCase.Layout, cfg.BaseURL, cfg.AuthAPI, cfg.BankAccountAPI, cfg.ValidateAPI, cfg.CreateTransactionAPI, cfg.TPBank.Name, cfg.TPBank.ApiKey, cfg.TPBank.SecretKey, cfg.TPBank.PrivateKey),
		repository.GetTransactionCreateRepository(client),
		repository.GetTransactionListRepository(client),
		repository.GetTransactionUpdateRepository(client),
		repository.GetTransactionIsNextRepository(client),
		repository.GetCustomerListRepository(client),
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.Mail.ConfirmEmailSubject,
		&cfg.Mail.ConfirmEmailTemplate,
		cfg.TransactionUseCase.Layout,
		cfg.BaseURL,
		cfg.AuthAPI,
		cfg.BankAccountAPI,
		cfg.ValidateAPI,
		cfg.CreateTransactionAPI,
		cfg.TPBank.Name,
		cfg.TPBank.ApiKey,
		cfg.TPBank.SecretKey,
		cfg.TPBank.PrivateKey,
		&cfg.TransactionUseCase.FeeAmount,
		cfg.Mail.OTPTimeout,
	)
	cDUc := constructors.NewCustomerDebtUseCase(
		repository.GetDebtListRepository(client),
		repository.GetDebtCreateRepository(client),
		repository.GetDebtUpdateRepository(client),
		repository.GetDebtFulfillRepository(client),
		repository.GetDebtIsNextRepository(client),
		repository.GetCustomerListRepository(client),
		repository.GetBankAccountListRepository(client),
		task.GetDebtTaskExecutor(b, l),
		task.GetEmailTaskExecutor(cfg.Mail.Host, cfg.Mail.User, cfg.Mail.Password, cfg.Mail.SenderName, cfg.Mail.Port, l),
		&cfg.App.SecretKey,
		&cfg.Mail.ConfirmEmailSubject,
		&cfg.Mail.ConfirmEmailTemplate,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.TransactionUseCase.FeeAmount,
		cfg.Mail.OTPTimeout,
	)
	cStreamUc := constructors.NewCustomerStreamUseCase(
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	cCUc := constructors.NewCustomerContactUseCase(
		repository.GetContactListRepository(client),
		repository.GetContactUpdateRepository(client),
		repository.GetContactCreateRepository(client),
		repository.GetContactDeleteRepository(client),
		repository.GetContactIsNextRepository(client),
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		cfg.TransactionUseCase.Layout,
		cfg.BaseURL,
		cfg.AuthAPI,
		cfg.BankAccountAPI,
		cfg.ValidateAPI,
		cfg.CreateTransactionAPI,
		cfg.TPBank.Name,
		cfg.TPBank.ApiKey,
		cfg.TPBank.SecretKey,
		cfg.TPBank.PrivateKey,
		&cfg.TransactionUseCase.FeeAmount,
	)
	cCOUc := constructors.NewOptionUseCase(
		&cfg.App.SecretKey,
		&cfg.App.Name,
		cfg.TPBank.Name,
	)

	// Employee UseCase
	eUc1 := constructors.NewEmployeeMeUseCase(
		repository.GetEmployeeListRepository(client),
		repository.GetEmployeeUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	eUc2 := constructors.NewEmployeeAuthUseCase(
		repository.GetEmployeeListRepository(client),
		repository.GetEmployeeUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		cfg.AuthUseCase.AccessTTL,
		cfg.AuthUseCase.RefreshTTL,
	)
	eUc3 := constructors.NewEmployeeCustomerUseCase(
		repository.GetCustomerListRepository(client),
		repository.GetCustomerCreateRepository(client),
		repository.GetCustomerIsNextRepository(client),
		repository.GetEmployeeListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	eUc4 := constructors.NewEmployeeBankAccountUseCase(
		repository.GetBankAccountUpdateRepository(client),
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountIsNextRepository(client),
		repository.GetBankAccountDeleteRepository(client),
		repository.GetEmployeeListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	eUc5 := constructors.NewEmployeeTransactionUseCase(
		repository.GetTransactionListRepository(client),
		repository.GetTransactionIsNextRepository(client),
		repository.GetEmployeeListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	// Admin UseCase
	aUc1 := constructors.NewAdminMeUseCase(
		repository.GetAdminListRepository(client),
		repository.GetAdminUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	aUc2 := constructors.NewAdminAuthUseCase(
		repository.GetAdminListRepository(client),
		repository.GetAdminUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		cfg.AuthUseCase.AccessTTL,
		cfg.AuthUseCase.RefreshTTL,
	)
	aUc3 := constructors.NewAdminTransactionUseCase(
		repository.GetTransactionListRepository(client),
		repository.GetTransactionIsNextRepository(client),
		repository.GetAdminListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	aUc4 := constructors.NewAdminEmployeeUseCase(
		repository.GetEmployeeListRepository(client),
		repository.GetEmployeeCreateRepository(client),
		repository.GetEmployeeUpdateRepository(client),
		repository.GetEmployeeDeleteRepository(client),
		repository.GetEmployeeIsNextRepository(client),
		repository.GetAdminListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	// partner
	pUc1 := constructors.NewPartnerAuthUseCase(
		repository.GetPartnerListRepository(client),
		&cfg.App.SecretKey,
		cfg.AuthUseCase.AccessTTL,
		cfg.AuthUseCase.RefreshTTL,
	)
	pUc2 := constructors.NewPartnerTransactionUseCase(
		repository.GetBankAccountListRepository(client),
		repository.GetCustomerListRepository(client),
		repository.GetPartnerTransactionCreateRepository(client),
		repository.GetPartnerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.TransactionUseCase.Layout,
		&cfg.TransactionUseCase.FeeAmount,
	)
	pUc3 := constructors.NewPartnerBankAccountUseCase(
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountIsNextRepository(client),
		repository.GetPartnerListRepository(client),
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)

	v1.RegisterV1HTTPServices(
		handler,
		CMeUc,
		CAuthUc,
		cBankAccountUc,
		cStreamUc,
		cTxcUc,
		cDUc,
		cCUc,
		cCOUc,
		cUC1,
		eUc2,
		eUc1,
		eUc3,
		eUc4,
		eUc5,
		aUc1,
		aUc2,
		aUc3,
		aUc4,
		pUc1,
		pUc2,
		pUc3,
		b,
		l,
	)
	handler.Listen(fmt.Sprintf(":%s", cfg.HTTP.Port))
}
