package password

import (
	"context"
	"crypto"
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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
	tk := GenerateHashData(ctx, data, secretKey)
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

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pKBuitIn, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("expect rsa.PublicKey but got %t", pub)
	}
	return pKBuitIn, nil
}

func VerifySignature(ctx context.Context, sig, msg, publicK string) error {
	hashed := sha256.Sum256([]byte(msg))
	pub, err := ParsePublicKey(publicK)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], []byte(sig))
	return err
}
