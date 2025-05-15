package dto

type AuthRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type AuthResponse struct {
	SessionID string `json:"sessionId"`
}

type AuthVerifyRequest struct {
	SessionID string `json:"sessionId" validate:"required"`
	Code      int    `json:"code" validate:"min=4,required"`
}

type AuthVerifyResponse struct {
	Token string `json:"token"`
}
