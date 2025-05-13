package middlewares

import (
	"errors"
	"gin-api-blog/api/helpers"
	"gin-api-blog/config"
	"gin-api-blog/constants"
	"gin-api-blog/pkg/logging/service_errors"
	"gin-api-blog/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	tokenService := service.NewTokenService(cfg)

	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			err := &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helpers.GenerateBaseResponseWithError(nil, false, helpers.AuthError, err),
			)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			err := &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helpers.GenerateBaseResponseWithError(nil, false, helpers.AuthError, err),
			)
			return
		}

		claimMap, err := tokenService.GetClaims(parts[1])
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenMalformed}
			} else if errors.Is(err, jwt.ErrSignatureInvalid) {
				err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
			} else {
				err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized,
				helpers.GenerateBaseResponseWithError(nil, false, helpers.AuthError, err),
			)
			return
		}

		c.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		c.Set(constants.NameKey, claimMap[constants.NameKey])
		c.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		c.Set(constants.EmailKey, claimMap[constants.EmailKey])
		c.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		c.Next()
	}
}
