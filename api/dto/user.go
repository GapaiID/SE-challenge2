package dto

type UserQueryParams struct {
	SearchFilter
	OrderFilter
	PaginationParams
}

func (q UserQueryParams) SearchFields() []string {
	return []string{"email", "name"}
}

func (q UserQueryParams) OrderFields() []string {
	return []string{"id"}
}

func (q UserQueryParams) DefaultPageSize() int {
	return 10
}

type UserPagination struct {
	List       []*User     `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type MeUpdateRequest struct {
	Name string `json:"name"`
}

type MeResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MePasswordRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=20"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=20"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=20"`
}

func (d MePasswordRequest) PasswordMatches() bool {
	return d.NewPassword == d.ConfirmPassword
}

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
