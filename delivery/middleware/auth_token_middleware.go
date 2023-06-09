package middleware

import (
	"simple-payment/util/authenticator"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequiredToken() gin.HandlerFunc
	DecodeToken(string) (jwt.MapClaims, error)
}

type authTokenMiddleware struct {
	acctToken authenticator.AccessToken
}

func (a *authTokenMiddleware) RequiredToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{
			AuthorizationHeader: "",
		}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)

		tokenCookies, _ := ctx.Cookie("token")

		if tokenString == "" && tokenCookies == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized, Get yourself token first by login",
			})
			ctx.Abort()
			return
		} else if tokenString == "" {
			tokenString = tokenCookies
		}

		token, err := a.acctToken.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized, Get yourself a token first by login",
			})
			ctx.Abort()
			return
		}
		if token != nil {
			ctx.Next()
		} else {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized, Get yourself a token first by login",
			})
			ctx.Abort()
			return
		}
	}
}

func (a *authTokenMiddleware) DecodeToken(bearerToken string) (jwt.MapClaims, error) {
	tokenString := strings.Replace(bearerToken, "Bearer ", "", -1)

	jwtClaim, err := a.acctToken.VerifyAccessToken(tokenString)
	if err != nil {
		return nil, err
	}
	return jwtClaim, nil
}

func NewTokenValidator(acctToken authenticator.AccessToken) AuthTokenMiddleware {
	return &authTokenMiddleware{
		acctToken: acctToken,
	}
}
