package password

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type RSAKeyPair struct {
	PublicKey  string
	PrivateKey string
}

// ValidatePassword validates a plain password against the model's password.
func ValidatePassword(passwordHash, password string) error {
	bytePassword := []byte(password)
	bytePasswordHash := []byte(passwordHash)

	// comparing the password with the hash
	return bcrypt.CompareHashAndPassword(bytePasswordHash, bytePassword)
}

// SetPassword sets cryptographically secure string to `model.Password`.
func GetHashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("The provided plain password is empty")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateHashData(ctx context.Context, secretKey, data string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
func ValidateHashData(ctx context.Context, data, secretKey, token string) error {
	tk := GenerateHashData(ctx, secretKey, data)
	if tk != token {
		return errors.New("invalid token")
	}
	return nil
}

func ParsePublicKey(publicK string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicK))
	if block == nil {
		return nil, errors.New("failed to parse")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub, nil
}
func ParsePrivateKey(privateK string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateK))
	if block == nil {
		return nil, errors.New("failed to parse")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func VerifySignature(ctx context.Context, sig, msg, publicK string) error {
	pub, err := ParsePublicKey(publicK)
	if err != nil {
		return err
	}
	hashed := sha256.Sum256([]byte(msg))
	sigDe, err := hex.DecodeString(sig)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], sigDe)
	return err
}
func GenerateSignature(ctx context.Context, msg, privateK string) (string, error) {
	msgByte := []byte(msg)
	msgHash := sha256.New()
	_, err := msgHash.Write(msgByte)
	if err != nil {
		return "", err
	}
	msgHashSum := msgHash.Sum(nil)

	priv, err := ParsePrivateKey(privateK)
	if err != nil {
		return "", err
	}
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, msgHashSum)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(signature), nil
}

func GenerateRSAKeyPair() (*RSAKeyPair, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	pub := key.Public()
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)
	return &RSAKeyPair{
		PublicKey:  string(pubPEM),
		PrivateKey: string(keyPEM),
	}, nil
}
