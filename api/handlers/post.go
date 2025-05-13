package handlers

import (
	"gin-api-blog/api/helpers"
	"gin-api-blog/config"
	"gin-api-blog/data/db"
	"gin-api-blog/data/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	cfg *config.Config
}

func NewPostsHandler(cfg *config.Config) *PostHandler {
	return &PostHandler{
		cfg: cfg,
	}
}

func (h *PostHandler) GetAllPost(c *gin.Context) {
	db := db.GetDB()

	var posts []models.Post
	query := `SELECT id, title, content, author_id, created_at, updated_at FROM post`
	err := db.Select(&posts, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err),
		)
		return
	}

	c.JSON(
		http.StatusOK,
		helpers.GenerateBaseResponse(posts, true, helpers.Success),
	)
}
