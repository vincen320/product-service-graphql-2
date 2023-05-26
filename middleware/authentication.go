package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/vincen320/product-service-graphql-2/state"
)

var JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTPayload struct {
	jwt.StandardClaims
}

func JWTAuthentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return state.NewResponse(http.StatusUnauthorized, "authorization is empty").JSON(c.Response())
			}

			tokenStr := authHeader[len("Bearer "):]
			var claims JWTPayload

			token, err := jwt.ParseWithClaims(tokenStr, &claims,
				func(t *jwt.Token) (interface{}, error) {
					_, ok := t.Method.(*jwt.SigningMethodHMAC)
					if !ok {
						return nil, errors.New("invalid method")
					}
					return JWT_SECRET_KEY, nil
				})

			if err != nil {
				if errors, ok := err.(*jwt.ValidationError); ok && errors.Errors == jwt.ValidationErrorExpired {
					return state.NewResponse(http.StatusUnauthorized, "token is expired").JSON(c.Response())
				}

				if err == jwt.ErrSignatureInvalid {
					return state.NewResponse(http.StatusUnauthorized, jwt.ErrSignatureInvalid.Error()).JSON(c.Response())
				}
			}

			if !token.Valid {
				return state.NewResponse(http.StatusUnauthorized, "invalid token").JSON(c.Response())
			}
			c.Set("user-id", claims.Id)
			return next(c)
		}
	}
}
