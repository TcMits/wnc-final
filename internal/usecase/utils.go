package usecase

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/TcMits/wnc-final/pkg/tool/password"
)

type UserCtxType string

const (
	UserCtxKey UserCtxType = "user"
	UserCtxVal UserCtxType = "user-value"
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
	return fmt.Sprintf("%v-%v-%v-%v-%v-%v", user.ID.String(), user.JwtTokenKey, user.IsActive, user.PhoneNumber, user.Email, user.Password)
}
func MakeOTPValue(ctx context.Context, otp string, extra ...string) string {
	tk := MakeValue(ctx)
	base := fmt.Sprintf("%v-%v", otp, tk)
	for _, e := range extra {
		base = fmt.Sprintf("%v-%v", base, e)
	}
	return base
}

func GetUserAsCustomer(ctx context.Context) *model.Customer {
	user, ok := ctx.Value(string(UserCtxKey)).(*model.Customer)
	if !ok {
		return nil
	}
	return user
}
func GetUserAsEmployee(ctx context.Context) *model.Employee {
	user, ok := ctx.Value(string(UserCtxKey)).(*model.Employee)
	if !ok {
		return nil
	}
	return user
}
func GetUserAsAdmin(ctx context.Context) *model.Admin {
	user, ok := ctx.Value(string(UserCtxKey)).(*model.Admin)
	if !ok {
		return nil
	}
	return user
}
func GetUserAsPartner(ctx context.Context) *model.Partner {
	user, ok := ctx.Value(string(UserCtxKey)).(*model.Partner)
	if !ok {
		return nil
	}
	return user
}
func EmbedUser(ctx context.Context, u *model.Customer) context.Context {
	newCtx := context.WithValue(ctx, UserCtxKey, UserCtxVal)
	newCtx = context.WithValue(newCtx, UserCtxVal, u)
	return newCtx
}
func GenerateConfirmTxcToken(
	ctx context.Context,
	token,
	signingKey string,
	isFeePaidByMe bool,
	secondsDuration time.Duration,
) (string, error) {
	tk, err := jwt.NewToken(map[string]any{
		"token":             token,
		"is_fee_paid_by_me": isFeePaidByMe,
	}, signingKey, secondsDuration)
	if err != nil {
		return "", WrapError(fmt.Errorf("internal.usecase.utils.GenerateConfirmTxcToken: %s", err))
	}
	return tk, nil
}
func GenerateForgetPwdToken(
	ctx context.Context,
	token,
	email,
	signingKey string,
	secondsDuration time.Duration,
) (string, error) {
	tk, err := jwt.NewToken(map[string]any{
		"token": token,
		"email": email,
	}, signingKey, secondsDuration)
	if err != nil {
		return "", WrapError(fmt.Errorf("internal.usecase.utils.GenerateConfirmTxcToken: %s", err))
	}
	return tk, nil
}
func GenerateFulfillToken(
	ctx context.Context,
	token,
	signingKey string,
	secondsDuration time.Duration,
) (string, error) {
	tk, err := jwt.NewToken(map[string]any{
		"token": token,
	}, signingKey, secondsDuration)
	if err != nil {
		return "", WrapError(fmt.Errorf("internal.usecase.utils.GenerateFulfillToken: %s", err))
	}
	return tk, nil
}

func ParseToken(
	ctx context.Context,
	token string,
	signingKey string,
) (map[string]any, error) {
	pl, err := jwt.ParseJWT(token, signingKey)
	if err != nil {
		return nil, WrapError(fmt.Errorf("internal.usecase.utils.ParseConfirmTxcToken: %s", err))
	}
	return pl, nil
}

func GenerateHashInfo(
	v string,
) (string, error) {
	hashPwd, err := password.GetHashPassword(v)
	if err != nil {
		return "", WrapError(fmt.Errorf("internal.usecase.utils.GenerateHashInfo: %s", err))
	}
	return hashPwd, nil
}

func ValidateHashInfo(
	raw,
	hash string,
) error {
	err := password.ValidatePassword(hash, raw)
	if err != nil {
		return WrapError(fmt.Errorf("internal.usecase.utils.ValidateHashInfo: %s", err))
	}
	return nil
}

func AuthenticateCtx(ctx *context.Context, c *ent.Client, user *model.Customer) {
	if user == nil {
		user, _ = ent.MustCustomerFactory().CreateWithClient(*ctx, c)
	}
	*ctx = context.WithValue(*ctx, string(UserCtxKey), user)
}
