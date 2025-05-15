package entity

type Auth struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"user_id"`
	Session string `json:"session"`
}
