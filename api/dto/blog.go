package dto

type BlogUser struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BlogPost struct {
	ID    uint     `json:"id"`
	Title string   `json:"title"`
	Body  string   `json:"body"`
	User  BlogUser `json:"user"`
}

type BlogPostPagination struct {
	List       []BlogPost  `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type BlogPostCreateRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogPostCreateResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
