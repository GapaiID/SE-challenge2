package dto

type CommentUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	User *CommentUser `json:"user"`
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
