package dto

type UserInComment struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PostInComment struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type Comment struct {
	ID   uint           `json:"id"`
	Body string         `json:"body"`
	User *UserInComment `json:"user"`
	Post *PostInComment `json:"post"`
}

type CommentQueryParams struct {
	Q string `json:"q" query:"q"`
	PaginationParams
}

type CommentPagination struct {
	List       []*Comment  `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type CommentCreateRequest struct {
	Body   string `json:"body" validate:"required,min=10,max=500"`
	PostID uint   `json:"post_id" validate:"required,gte=0"`
}

type CommentCreateResponse struct {
	ID     uint   `json:"id"`
	Body   string `json:"body"`
	PostID uint   `json:"post_id"`
}

type CommentUpdateRequest struct {
	Body string `json:"body" validate:"min=10,max=500"`
}

type CommentUpdateResponse struct {
	ID     uint   `json:"id"`
	Body   string `json:"body"`
	PostID uint   `json:"post_id"`
}
