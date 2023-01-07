package main

import (
	"context"
	"log"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
)

func main() {
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
	createUser(client)
	l.Info("finish create data")
}
func createUser(client *ent.Client) {
	ctx := context.Background()
	user, err := client.Customer.Query().Where(customer.Username("superuser")).First(ctx)
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
	if user == nil {
		log.Fatalf("failed generate data: user does not exist")
	}
	bA, err := ent.CreateFakeBankAccount(ctx, client, nil,
		ent.Opt{
			Key:   "CustomerID",
			Value: user.ID,
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
