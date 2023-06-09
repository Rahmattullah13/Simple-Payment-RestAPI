package authenticator

import (
	"fmt"
	"log"
	"simple-payment/config"
	"simple-payment/model"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessToken interface {
	CreateAccessToken(cred *model.UserCredential) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type accessToken struct {
	cfg config.TokenConfig
}

func (t *accessToken) CreateAccessToken(cred *model.UserCredential) (string, error) {
	now := time.Now().UTC()
	end := now.Add(t.cfg.AccessTokenLifeTime)
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: t.cfg.ApplicationName,
		},
		UserId: cred.UserId,
		Email:  cred.Email,
	}
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = end.Unix()

	token := jwt.NewWithClaims(
		t.cfg.JwtSigningMethod,
		claims,
	)

	return token.SignedString([]byte(t.cfg.JwtSignatureKey))
}

func (t *accessToken) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing Method Invalid")
		} else if method != t.cfg.JwtSigningMethod {
			return nil, fmt.Errorf("Signing MEthod Invalid")
		}
		return []byte(t.cfg.JwtSignatureKey), nil
	})

	if err != nil {
		fmt.Println(err)
		fmt.Println(token)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["iss"] != t.cfg.ApplicationName {
		log.Println("Token Invalid")
		return nil, err
	}
	return claims, nil
}

func NewAccessToken(config config.Config) AccessToken {
	return &accessToken{
		cfg: config.TokenConfig,
	}
}
