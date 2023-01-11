package configpartner

import (
	"encoding/base64"
	"fmt"
	"log"

	"crypto/rand"
	"crypto/rsa"

	"github.com/Pallinder/go-randomdata"
	"github.com/TcMits/wnc-final/config"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("failed generate config vendor: %v", err)
	}
	public := base64.StdEncoding.EncodeToString(privateKey.N.Bytes())
	private := base64.StdEncoding.EncodeToString(privateKey.D.Bytes())
	s.RSAPrivateK = private
	s.RSAPublicK = public
	l.Info("finish generate config vendor!")
	fmt.Println(s)
}
