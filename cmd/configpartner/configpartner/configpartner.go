package configpartner

import (
	"fmt"
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type ConfigVendor struct {
	ApiKey      string
	SecretKey   string
	RSAPrivateK string
	RSAPublicK  string
}

func (s *ConfigVendor) String() string {
	return fmt.Sprintf("\nApiKey: \n%s\nSecretKey: \n%s\nRSAPrivateKey: \n%s\nRSAPublicKey: \n%s", s.ApiKey, s.SecretKey, s.RSAPrivateK, s.RSAPublicK)
}

func GenConfigVendor() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	l := logger.New(cfg.Log.Level)
	l.Info("creating config...")
	s := new(ConfigVendor)
	s.ApiKey = randomdata.Alphanumeric(30)
	s.SecretKey = randomdata.Alphanumeric(30)
	pair, err := password.GenerateRSAKeyPair()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	s.RSAPrivateK = pair.PrivateKey
	s.RSAPublicK = pair.PublicKey
	l.Info("finish generate config vendor!")
	fmt.Println(s)
}
