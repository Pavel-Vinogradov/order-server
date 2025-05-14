package dto

type AuthRequest struct {
	Phone string `json:"phone"`
}

type AuthResponse struct {
	SessionID string `json:"sessionId"`
}
