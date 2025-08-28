package dto

type AuthRequest struct {
	Phone string `json:"phone" validate:"required,e164"`
}

type AuthResponse struct {
	SessionID string `json:"sessionId"`
}

type AuthVerifyRequest struct {
	SessionID string `json:"sessionId" validate:"required"`
	Code      string `json:"code" validate:"required,min=4,max=6"`
}

type AuthVerifyResponse struct {
	Token string `json:"token"`
}
