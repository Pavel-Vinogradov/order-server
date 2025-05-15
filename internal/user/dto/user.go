package dto

type CreateUserRequest struct {
	Phone string `json:"phone" validate:"required,phone"`
}
