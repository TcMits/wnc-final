// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TcMits/wnc-final/config"
	v1 "github.com/TcMits/wnc-final/internal/controller/http/v1"
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/contact"
	"github.com/TcMits/wnc-final/internal/usecase/debt"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/internal/usecase/stream"
	"github.com/TcMits/wnc-final/internal/usecase/transaction"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/httpserver"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	client, err := datastore.NewClient(cfg.PG.URL, cfg.PG.PoolMax)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer client.Close()

	// HTTP Server
	handler := v1.NewHandler()

	// Broker Server sent event
	b := sse.NewBroker(l)

	// Usecase
	CMeUc := me.NewCustomerMeUseCase(
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	CAuthUc := auth.NewCustomerAuthUseCase(
		repository.GetCustomerListRepository(client),
		repository.GetCustomerUpdateRepository(client),
		&cfg.App.SecretKey,
		cfg.AuthUseCase.AccessTTL,
		cfg.AuthUseCase.RefreshTTL,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeAmount,
		&cfg.TransactionUseCase.FeeDesc,
	)
	cBankAccountUc := bankaccount.NewCustomerBankAccountUseCase(
		repository.GetBankAccountUpdateRepository(client),
		repository.GetBankAccountListRepository(client),
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
		repository.GetCustomerListRepository(client),
		repository.GetBankAccountListRepository(client),
		repository.GetBankAccountUpdateRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.Mail.ConfirmEmailSubject,
		&cfg.Mail.FrontendURL,
		&cfg.Mail.ConfirmEmailTemplate,
		&cfg.TransactionUseCase.FeeAmount,
		cfg.Mail.OTPTimeout,
	)
	cDUc := debt.NewCustomerDebtUseCase(
		repository.GetDebtListRepository(client),
		repository.GetDebtCreateRepository(client),
		repository.GetDebtUpdateRepository(client),
		repository.GetDebtFulfillRepository(client),
		repository.GetCustomerListRepository(client),
		repository.GetBankAccountListRepository(client),
		task.GetDebtTaskExecutor(b, l),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.TransactionUseCase.FeeAmount,
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
		repository.GetCustomerListRepository(client),
		&cfg.App.SecretKey,
		&cfg.App.Name,
		&cfg.TransactionUseCase.FeeDesc,
		&cfg.TransactionUseCase.FeeAmount,
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
