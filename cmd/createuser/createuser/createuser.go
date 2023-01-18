package createuser

import (
	"context"
	"log"

	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/customer"
	"github.com/TcMits/wnc-final/ent/employee"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

func CreateUser() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Info("creating user...")
	client, err := datastore.NewClient(cfg.DB.URL, cfg.DB.PoolMax, cfg.App.Debug)
	if err != nil {
		log.Fatalf("failed opening database client: %v", err)
	}
	defer client.Close()
	createUser(client)
	l.Info("finish create user")
}
func createUser(client *ent.Client) {
	ctx := context.Background()
	hashPw, err := password.GetHashPassword("123456789")
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	if len(client.Customer.Query().Where(customer.Username("superuser")).AllX(ctx)) == 0 {
		_, err = ent.MustCustomerFactory(
			ent.Opt{
				Key:   "Username",
				Value: "superuser",
			},
			ent.Opt{
				Key:   "Password",
				Value: generic.GetPointer(hashPw),
			},
			ent.Opt{
				Key:   "Email",
				Value: "su@gmail.com",
			},
			ent.Opt{
				Key:   "PhoneNumber",
				Value: "+84923456789",
			},
			ent.Opt{
				Key:   "IsActive",
				Value: generic.GetPointer(true),
			},
		).CreateWithClient(ctx, client)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
	}
	if len(client.Employee.Query().Where(employee.Username("superuser")).AllX(ctx)) == 0 {
		_, err = ent.MustEmployeeFactory(
			ent.Opt{
				Key:   "Username",
				Value: "superuser",
			},
			ent.Opt{
				Key:   "Password",
				Value: generic.GetPointer(hashPw),
			},
			ent.Opt{
				Key:   "IsActive",
				Value: generic.GetPointer(true),
			},
		).CreateWithClient(ctx, client)
		if err != nil {
			log.Fatalf("failed creating user: %v", err)
		}
	}
}
