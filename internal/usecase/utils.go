package usecase

import (
	"context"
	"time"

	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
)

func GetUserAsCustomer(ctx context.Context) *model.Customer {
	uAny := ctx.Value("user")
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
