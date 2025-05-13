package routers

import (
	"gin-api-blog/api/handlers"
	"gin-api-blog/config"

	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewAuthHandler(cfg)
	router.POST("/register", h.Register)
	router.POST("/login", h.LoginByUsername)
}
