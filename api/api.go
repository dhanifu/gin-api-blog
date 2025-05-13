package api

import (
	"fmt"
	"gin-api-blog/api/routers"
	"gin-api-blog/api/validations"
	"gin-api-blog/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	RegisterValidators()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Gin API Blog",
		})
	})
	v1 := r.Group("api/v1")
	{
		auth := v1.Group("auth")
		routers.Auth(auth, cfg)

		posts := v1.Group("posts")
		routers.Posts(posts, cfg)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			log.Fatal("validation failed")
		}
	}
}
