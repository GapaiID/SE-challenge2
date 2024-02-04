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
}

type CommentCreateResponse struct {
}

type CommentUpdateRequest struct {
}

type CommentUpdateResponse struct {
}
