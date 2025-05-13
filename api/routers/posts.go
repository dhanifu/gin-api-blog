package routers

import (
	"gin-api-blog/api/handlers"
	"gin-api-blog/api/middlewares"
	"gin-api-blog/config"

	"github.com/gin-gonic/gin"
)

func Posts(router *gin.RouterGroup, cfg *config.Config) {
	p := handlers.NewPostsHandler(cfg)
	router.GET("/", middlewares.Authentication(cfg), p.GetAllPost)
}
