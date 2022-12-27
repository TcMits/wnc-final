package usecase

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type UserCtxType string

const (
	UserCtx UserCtxType = "user"
)

func EncodeToString(max int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func GenerateOTP(max int) string {
	return EncodeToString(max)
}
func MakeValue(ctx context.Context) string {
	user := GetUserAsCustomer(ctx)
	fmt.Println(user)
	return fmt.Sprintf("%v-%v-%v-%v-%v-%v", user.ID.String(), user.JwtTokenKey, user.IsActive, user.PhoneNumber, user.Email, user.Password)
}
func MakeOTPValue(ctx context.Context, otp string) string {
	tk := MakeValue(ctx)
	return fmt.Sprintf("%v-%v", otp, tk)
}

func GetUserAsCustomer(ctx context.Context) *model.Customer {
	uAny := ctx.Value(UserCtx)
	user, ok := uAny.(*model.Customer)
	if !ok {
		return nil
	}
	return user
}
func GenerateConfirmTxcToken(
	ctx context.Context,
	payload map[string]any,
	signingKey string,
	secondsDuration time.Duration,
) (string, error) {
	return jwt.NewToken(payload, signingKey, secondsDuration)
}

func ParseConfirmTxcToken(
	ctx context.Context,
	token string,
	signingKey string,
) (map[string]any, error) {
	return jwt.ParseJWT(token, signingKey)
}

func GenerateHashInfo(
	v string,
) (string, error) {
	return password.GetHashPassword(v)
}

func ValidateHashInfo(
	raw,
	hash string,
) error {
	return password.ValidatePassword(hash, raw)
}
