package helper

import (
	"errors"
	"github.com/elsyarif/pms-api/pkg/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserId   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	jwt.RegisteredClaims
}

var (
	AccessToken  = "AccessToken"
	RefreshToken = "RefreshToken"
)

func GenerateToken(id string, username string, code string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	jwtKey := ""
	var age time.Duration
	switch code {
	case AccessToken:
		age = config.Conf.AccessTokenAge
		expirationTime = time.Now().Add(age * time.Second)
		jwtKey = config.Conf.AccessTokenKey
	case RefreshToken:
		expirationTime = time.Now().Add(72 * time.Hour)
		jwtKey = config.Conf.RefreshTokenKey
	}

	claim := &Claims{
		UserId:   id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func VerifyToken(tokenString string, code string) (*Claims, error) {
	jwtKey := ""
	switch code {
	case AccessToken:
		jwtKey = config.Conf.AccessTokenKey
	case RefreshToken:
		jwtKey = config.Conf.RefreshTokenKey
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}

		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	expirationTime, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	if expirationTime.Unix() < time.Now().Unix() {
		return nil, jwt.ErrTokenExpired
	}

	return claims, nil
}
