package dto

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
