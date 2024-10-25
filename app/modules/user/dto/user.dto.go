package dto

import (
	"time"

	"github.com/google/uuid"
)

//	type UserDTOResponse struct {
//		ID        uuid.UUID `json:"id"`
//		FirstName string    `json:"first_name"`
//		LastName  string    `json:"last_name"`
//		Username  string    `json:"user_name"`
//		Password  string    `json:"password"`
//		CreatedAt int64     `json:"created_at"`
//		UpdatedAt int64     `json:"updated_at"`
//		DeletedAt int64     `json:"deleted_at"`
//	}
type UserDTOResponse struct {
	json      struct{} `naming:"snake_case"`
	ID        uuid.UUID
	FirstName string
	LastName  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserCreateRequest struct {
	FirstName string
	LastName  string
	Username  string
	Password  string
}

type UserResetPasswordReq struct {
	ID          uuid.UUID
	Username    string
	Password    string
	NewPassword string
}
