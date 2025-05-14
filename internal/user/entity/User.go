package entity

import "time"

type User struct {
	Id        int64
	Name      *string
	Password  *string
	Email     *string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
