package handlers

import (
	"gin-api-blog/api/dto"
	"gin-api-blog/api/helpers"
	"gin-api-blog/config"
	"gin-api-blog/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	service := service.NewUserService(cfg)
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) LoginByUsername(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, helpers.ValidationError, err),
		)
		return
	}

	token, err := h.service.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err),
		)
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(token, true, helpers.Success))
}

func (h *AuthHandler) Register(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, helpers.ValidationError, err),
		)
		return
	}
	err = h.service.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err),
		)
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, helpers.Success))
}
