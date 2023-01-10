package gendata

import (
	"context"
	"log"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func GenData() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Info("creating data...")
	client, err := datastore.NewClient(cfg.DB.URL, cfg.DB.PoolMax, cfg.App.Debug)
	if err != nil {
		log.Fatalf("failed opening database client: %v", err)
	}
	defer client.Close()
	genData(client, cfg)
	l.Info("finish create data")
}
func genData(client *ent.Client, cfg *config.Config) {
	// flush db
	ctx := context.Background()
	err := client.Flush(ctx)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	user, err := ent.CreateFakeCustomer(ctx, client, nil,
		ent.Opt{
			Key:   "Email",
			Value: "dinhphat611@gmail.com",
		},
	)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	bA, err := ent.CreateFakeBankAccount(ctx, client, nil,
		ent.Opt{
			Key:   "CustomerID",
			Value: user.ID,
		},
		ent.Opt{
			Key:   "CashIn",
			Value: float64(100000000),
		},
		ent.Opt{
			Key:   "IsForPayment",
			Value: generic.GetPointer(true),
		},
	)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	// transactions
	for i := 1; i < 3; i++ {
		_, err = ent.CreateFakeTransaction(ctx, client, nil,
			ent.Opt{
				Key:   "SenderID",
				Value: bA.ID,
			},
		)
		if err != nil {
			log.Fatalf("failed generate data: %v", err)
		}
	}
	// contacts
	for i := 1; i < 3; i++ {
		_, err = ent.CreateFakeContact(ctx, client, nil,
			ent.Opt{
				Key:   "BankName",
				Value: cfg.App.Name,
			},
			ent.Opt{
				Key:   "OwnerID",
				Value: user.ID,
			},
		)
		if err != nil {
			log.Fatalf("failed generate data: %v", err)
		}
	}
	// debts
	for i := 1; i < 3; i++ {
		ent.CreateFakeDebt(ctx, client, nil,
			ent.Opt{
				Key:   "OwnerID",
				Value: bA.ID,
			},
			ent.Opt{
				Key:   "OwnerBankName",
				Value: cfg.App.Name,
			},
			ent.Opt{
				Key:   "ReceiverBankName",
				Value: cfg.App.Name,
			},
		)
	}
	// customers
	u1, err := ent.CreateFakeCustomer(ctx, client, nil,
		ent.Opt{
			Key:   "Username",
			Value: "alanwalker",
		},
	)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	_, err = ent.CreateFakeBankAccount(ctx, client, nil,
		ent.Opt{
			Key:   "IsForPayment",
			Value: generic.GetPointer(true),
		},
		ent.Opt{
			Key:   "CustomerID",
			Value: u1.ID,
		},
	)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	// admins
	_, err = ent.CreateFakeAdmin(ctx, client, nil,
		ent.Opt{
			Key:   "Username",
			Value: "iamadmin",
		},
	)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	// employees
	_, err = ent.CreateFakeEmployee(ctx, client, nil,
		ent.Opt{
			Key:   "Username",
			Value: "iamemployee",
		},
	)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
}
