package dto

type CreatePostRequest struct {
	Title    string `json:"title" binding:"required,min=5"`
	Content  string `json:"content" binding:"required,min=10"`
	AuthorId uint   `json:"authorId" binding:"required,min=1"`
}

type UpdatePostRequest struct {
	Title   string `json:"title" binding:"min=5"`
	Content string `json:"content" binding:"min=10"`
}
