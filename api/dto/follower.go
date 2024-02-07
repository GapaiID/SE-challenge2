package dto

type UserInFollower struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type FollowerCreateRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}

type UnFollowerCreateRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}
