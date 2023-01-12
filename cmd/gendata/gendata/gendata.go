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
		ent.Opt{
			Key:   "AccountNumber",
			Value: generic.GetPointer("11112222333344445"),
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
	// partners
	_, err = ent.CreateFakePartner(ctx, client, &ent.PartnerCreateInput{
		APIKey:    "8JnDlw1CyEpr372uZL5S3OUoLARZgh",
		SecretKey: "QwZHAcABNd98ehV1Y1qkmlJTsDJjox",
		PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
		MIIEpAIBAAKCAQEApoL43bl4FCVmHJpsHzdxGiaMIxcsogjsBGryvERaZonQwj1K
		9rQ1oJds5uUvLBFhNqPC1DkvhvF1JO/5fgIXv9XF+PHjpIaPn81l0Lfg3vZWDynC
		MbuQhOzKFXlO8mJ5nRNmAxe+iLwSBlPEtjAe38E1XTaurenwLUHSD6NtH3Us0hu5
		N/LomlXpX4p6BTtfCwVYQGV7rh+pbKt4D5Ck4If0QwwHUz5UWBo8p0Rz7gFTYnUc
		RHAblt+Aos93rfWocsAgTIIM+hd9PoyIpT07YbkzvmuScqLuptNl3p2iUPDik+G3
		NpEW67bKVg1U190qQV38x6jhwGFkUCl4wT3rdwIDAQABAoIBABJlvCt7V4oaS7rz
		UOjuDkAObENgRx0sES+3xVQJ1Vco+PjNHuA9LwOsV2r3fYvH687GrzsVimvFd+T5
		4r4KaydV7kbAKx+9glbsscgq2NypqP6J5ZvrFl+pbfaLba6YNMmnIAlNcU7YQMGU
		NdLPZEnipgvDg+9RPqkxcY0qyF7GPOfMlNsT6koITBbKny26W3cxJ2XPj3emoplZ
		s+Ho/FkTjvPWLvYXujDprdNd0EAPAyPvMBBo/pECbK5SIG0aGR7X+mMKbZ7b1rET
		hfZRRjTrOG4r8QEVNFRPtaKiW0XokqNqjOj/CrG/DjNYtkoDJuBAqK8RfIME5zN2
		TP1V62kCgYEAyczH1j7DkUahgynlOREP3K71ppBXXxgAEWkCwlSdbCHJpYrjOaJc
		YQQyxeEWtxPyANWzj2Ms23wxFXsYw84U66qr9ifM4kmI52NyF8WRIblrBI4YEaZS
		Rc3VhE4DqIQS3zTO3BjpPq5peMEpKEsTja+jruCgVYxB9rM18ag26IUCgYEA0zvc
		rzEEcYa8R2myei4PrbhHUiBggzZYXSZmzSYB4nF2XFEY1JxRRVy/T5CPHedYAPRC
		+3JjxvH0HNrBGxDBFmgWqqfFuXmrxL0UdvXUvTMFM2rcEf9AzmkzRJR0kwQhoGFC
		Si30pR/udxQHORM5XsqF6V1bpRxHgxh2Ey8ggssCgYEAh5N7FudoAKxRSovVIrfh
		zkQFafhDmvHG6euBHQo4ETPKA9wBuPDp24w0iFknJ6zQw9rPHiBlPLh9SY9AHhzQ
		VVx+14v5zHXW8o4PUwU68kteKNtGNGVnlNoq/w5iys6g1zDlYV2jJqeK2MP9YMK/
		ykscTxs+3Pq0Pog8T7TR3vkCgYAlJYDQmjki0cPodxD45YLCVQbNzX3LdVIix28K
		oqVwMe3TnDtWoEq2fPHzxwM0Cgvy1wG0gFBFmyUHsfyFivO5wgJCbpCZ5mirh2jC
		5sZLo15FxYP/8jhuVBe89rJtbCuRrajfrKc6Jpxj+nSut8+9+LWF7XIBXjDBQBr7
		kq0P8wKBgQCD5h3sTdYlb8pvYf+HCE3OhYWz6qJ4pYXnwx9PR+rNIrMJGjLikxEM
		FF+a70g2/QbqHwX1yZD5IVoCJaLwpggcIUQNkYrZfrtW0MPKRNnX0HPYqtu8xKVB
		FXNsKFx4RGTykyS/sNJewxQvkP3QoBUhSsFGeZbz+w8e0qrn2JVOaA==
		-----END RSA PRIVATE KEY-----`,
		PublicKey: `-----BEGIN RSA PUBLIC KEY-----
		MIIBCgKCAQEApoL43bl4FCVmHJpsHzdxGiaMIxcsogjsBGryvERaZonQwj1K9rQ1
		oJds5uUvLBFhNqPC1DkvhvF1JO/5fgIXv9XF+PHjpIaPn81l0Lfg3vZWDynCMbuQ
		hOzKFXlO8mJ5nRNmAxe+iLwSBlPEtjAe38E1XTaurenwLUHSD6NtH3Us0hu5N/Lo
		mlXpX4p6BTtfCwVYQGV7rh+pbKt4D5Ck4If0QwwHUz5UWBo8p0Rz7gFTYnUcRHAb
		lt+Aos93rfWocsAgTIIM+hd9PoyIpT07YbkzvmuScqLuptNl3p2iUPDik+G3NpEW
		67bKVg1U190qQV38x6jhwGFkUCl4wT3rdwIDAQAB
		-----END RSA PUBLIC KEY-----`,
		Name: generic.GetPointer("Sacombank"),
	})
	if err != nil {
		log.Fatalf("failed generate data: %v", err)
	}
}
