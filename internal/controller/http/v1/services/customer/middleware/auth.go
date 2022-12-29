package middleware

import (
	goCtx "context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	goJWT "github.com/golang-jwt/jwt/v4"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	irisJWT "github.com/kataras/iris/v12/middleware/jwt"
)

const UserCtxKey usecase.UserCtxType = "iris.user"

func BaseAuthenticator(secretKey *string, userGetter func(goCtx.Context, map[string]any) (any, error), validators ...func(*context.Context, goJWT.MapClaims) error) iris.Handler {
	return func(ctx *context.Context) {
		extractors := []irisJWT.TokenExtractor{
			irisJWT.FromHeader,
			irisJWT.FromQuery,
		}
		token := ""
		for _, extract := range extractors {
			if token = extract(ctx); token != "" {
				break // ok we found it.
			}
		}
		if token == "" {
			ctx.StatusCode(iris.StatusTokenRequired)
			return
		}
		payload, err := jwt.ParseJWT(token, *secretKey)
		if err != nil {
			ctx.StatusCode(iris.StatusInvalidToken)
			return
		}
		user, err := userGetter(ctx, payload)
		if err != nil {
			ctx.StopWithError(iris.StatusInternalServerError, err)
			return
		}
		ctx.Values().Set(string(usecase.UserCtxKey), UserCtxKey)
		ctx.SetUser(user)
		for _, v := range validators {
			err = v(ctx, payload)
			if err != nil {
				ctx.StatusCode(iris.StatusInvalidToken)
				return
			}
		}
		ctx.Next()
	}
}

func validateToken(ctx *context.Context, claim goJWT.MapClaims) error {
	user := GetUserFromCtxAsCustomer(ctx)
	if user == nil {
		return fmt.Errorf("invalid token")
	}
	jwt_keyAny := claim["jwt_key"]
	jwt, ok := jwt_keyAny.(string)
	if !ok {
		return fmt.Errorf("invalid token")
	}
	if jwt != user.JwtTokenKey {
		return fmt.Errorf("unauthorized")
	}
	return nil
}

func GetUserFromCtxAsCustomer(ctx *context.Context) *model.Customer {
	userAny, err := ctx.User().GetRaw()
	if err != nil {
		return nil
	}
	user, ok := userAny.(*model.Customer)
	if !ok {
		return nil
	}
	return user
}

func Authenticator(
	secretKey *string,
	userGetter func(goCtx.Context, map[string]any) (any, error),
) iris.Handler {
	return BaseAuthenticator(secretKey, userGetter, validateToken)
}
