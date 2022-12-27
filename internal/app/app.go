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
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/infrastructure/backgroundserver"
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
	// Task client
	taskClient := backgroundserver.NewClient(cfg.Redis.URL, cfg.Redis.Password, cfg.Redis.DB)
	defer taskClient.Close()

	// HTTP Server
	handler := v1.NewHandler()

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

	v1.RegisterV1HTTPServices(handler,
		CMeUc,
		CAuthUc,
		cBankAccountUc,
		l,
	)

	if err := handler.Build(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - handler.Build: %w", err))
	}
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	l.Info("Listening and serving HTTP on %s", httpServer.Addr())

	// Task Worker
	workerHandler := task.NewHandler()
	// Register Tasks
	task.RegisterTask(workerHandler, l, cfg.Mail.Host, cfg.Mail.User, cfg.Mail.Password, cfg.Mail.SenderName, cfg.Mail.Port)

	workerServer := backgroundserver.NewWorkerServer(workerHandler, cfg.Redis.URL, cfg.Redis.Password, cfg.Redis.DB)

	// Starting worker server
	if err = workerServer.Run(); err != nil {
		l.Fatal(fmt.Errorf("app - Run - workerServer.Run: %w", err))
	}

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
