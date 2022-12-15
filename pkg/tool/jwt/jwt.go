package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AccessToken  = "access"
	RefreshToken = "refresh"
	Algorithm    = "HS256"
)

type TokenPair struct {
	AccessToken  *string
	RefreshToken *string
}

var SigningMethod = func() jwt.SigningMethod {
	return jwt.SigningMethodHS256
}

// ParseUnverifiedJWT parses JWT token and returns its claims
// but DOES NOT verify the signature.
func ParseUnverifiedJWT(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}

	parser := &jwt.Parser{}
	_, _, err := parser.ParseUnverified(token, claims)

	if err == nil {
		err = claims.Valid()
	}

	return claims, err
}

// ParseJWT verifies and parses JWT token and returns its claims.
func ParseJWT(token, verificationKey string) (jwt.MapClaims, error) {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{Algorithm}))

	parsedToken, err := parser.Parse(token, func(t *jwt.Token) (any, error) {
		return []byte(verificationKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("Unable to parse token.")
}

// NewToken generates and returns new HS256 signed JWT token.
func NewToken(
	payload jwt.MapClaims,
	signingKey string,
	secondsDuration time.Duration,
) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"exp": now.Add(secondsDuration).Unix(),
		"iat": now.Unix(),
	}
	for k, v := range payload {
		claims[k] = v
	}

	return jwt.NewWithClaims(SigningMethod(), claims).SignedString([]byte(signingKey))
}

func NewAccessToken(
	payload jwt.MapClaims,
	signingKey string,
	secondsDuration time.Duration,
) (string, error) {
	claims := jwt.MapClaims{
		"type": AccessToken,
	}
	for k, v := range payload {
		claims[k] = v
	}
	return NewToken(claims, signingKey, secondsDuration)
}

func NewRefreshToken(
	payload jwt.MapClaims,
	signingKey string,
	secondsDuration time.Duration,
) (string, error) {
	claims := jwt.MapClaims{
		"type": RefreshToken,
	}
	for k, v := range payload {
		claims[k] = v
	}
	return NewToken(payload, signingKey, secondsDuration)
}

func NewTokenPair(
	signingKey string,
	plAccessToken jwt.MapClaims,
	plRefreshToken jwt.MapClaims,
	sdAccessToken time.Duration,
	sdRefreshToken time.Duration,
) (*TokenPair, error) {
	aT, err := NewAccessToken(
		plAccessToken,
		signingKey,
		sdAccessToken,
	)
	if err != nil {
		return nil, err
	}
	rT, err := NewRefreshToken(
		plRefreshToken,
		signingKey,
		sdRefreshToken,
	)
	if err != nil {
		return nil, err
	}
	return &TokenPair{
		AccessToken:  &aT,
		RefreshToken: &rT,
	}, nil
}
