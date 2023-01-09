package gendata

import (
	"context"
	"log"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/customer"
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
	genData(client)
	l.Info("finish create data")
}
func genData(client *ent.Client) {
	ctx := context.Background()
	_, err := client.Customer.Query().Where(customer.Email("dinhphat611@gmail.com")).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Fatalf("failed generate data: %v", err)
	}
	if ent.IsNotFound(err) {
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
	}
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
}
