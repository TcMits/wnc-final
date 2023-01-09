// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TcMits/wnc-final/cmd/createuser/createuser"
	"github.com/TcMits/wnc-final/cmd/migrate/migrate"
	"github.com/TcMits/wnc-final/config"
	v1 "github.com/TcMits/wnc-final/internal/controller/http/v1"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/contact"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/debt"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/internal/usecase/option"
	"github.com/TcMits/wnc-final/internal/usecase/stream"
	"github.com/TcMits/wnc-final/internal/usecase/transaction"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/httpserver"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// setup
	if !cfg.App.Debug {
		migrate.Migrate()
		createuser.CreateUser()
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
	CMeUc := me.NewCustomerMeUseCase(
		repository.GetCustomerListRepository(client),
		repository.GetCustomerUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	CAuthUc := auth.NewCustomerAuthUseCase(
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
	cBankAccountUc := bankaccount.NewCustomerBankAccountUseCase(
		repository.GetBankAccountUpdateRepository(client),
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountIsNextRepository(client),
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	cTxcUc := transaction.NewCustomerTransactionUseCase(
		task.GetEmailTaskExecutor(cfg.Mail.Host, cfg.Mail.User, cfg.Mail.Password, cfg.Mail.SenderName, cfg.Mail.Port, l),
		repository.GetTransactionConfirmSuccessRepository(client),
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
		&cfg.TransactionUseCase.FeeAmount,
		cfg.Mail.OTPTimeout,
	)
	cDUc := debt.NewCustomerDebtUseCase(
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
	cStreamUc := stream.NewCustomerStreamUseCase(
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	cCUc := contact.NewCustomerContactUseCase(
		repository.GetContactListRepository(client),
		repository.GetContactUpdateRepository(client),
		repository.GetContactCreateRepository(client),
		repository.GetContactDeleteRepository(client),
		repository.GetContactIsNextRepository(client),
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.TransactionUseCase.FeeAmount,
	)
	cCOUc := option.NewOptionUseCase(
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)

	// Employee UseCase
	eUc1 := me.NewEmployeeMeUseCase(
		repository.GetEmployeeListRepository(client),
		repository.GetEmployeeUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	eUc2 := auth.NewEmployeeAuthUseCase(
		repository.GetEmployeeListRepository(client),
		repository.GetEmployeeUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		cfg.AuthUseCase.AccessTTL,
		cfg.AuthUseCase.RefreshTTL,
	)
	eUc3 := customer.NewEmployeeCustomerUseCase(
		repository.GetCustomerListRepository(client),
		repository.GetCustomerCreateRepository(client),
		repository.GetCustomerIsNextRepository(client),
		repository.GetEmployeeListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
	)
	eUc4 := bankaccount.NewEmployeeBankAccountUseCase(
		repository.GetBankAccountUpdateRepository(client),
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountIsNextRepository(client),
		repository.GetEmployeeListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
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
		eUc2,
		eUc1,
		eUc3,
		eUc4,
		b,
		l,
	)

	if err := handler.Build(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Build: %w", err))
	}
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Info("Listening and serving HTTP on %s", httpServer.Addr())

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
