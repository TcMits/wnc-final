package main

import (
	"context"
	"log"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Info("creating user...")
	client, err := datastore.NewClient(cfg.Sqlite.URL, cfg.Sqlite.PoolMax)
	if err != nil {
		log.Fatalf("failed opening sqlite client: %v", err)
	}
	defer client.Close()
	createUser(client)
	l.Info("finish create user")
}
func createUser(client *ent.Client) {
	hashPw, err := password.GetHashPassword("123456789")
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	if len(client.Customer.Query().Where(customer.Username("superuser")).AllX(context.Background())) == 0 {
		if _, err := client.Customer.Create().SetUsername("superuser").SetPassword(hashPw).SetEmail("su@gmail.com").SetPhoneNumber("+84923456789").SetIsActive(true).Save(context.Background()); err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
	}
}
