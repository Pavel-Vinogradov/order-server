package repository

import (
	"database/sql"
	"order-server/internal/user/entity"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(u entity.User) (entity.User, error) {
	q := `INSERT INTO users (name, password, phone) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(q, u.Name, u.Password, u.Phone).Scan(&u.Id, &u.CreatedAt, &u.UpdatedAt)
	return u, err
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByID(id int64) (entity.User, error) {
	var u entity.User
	q := `SELECT * FROM users WHERE id=$1`
	err := r.db.QueryRow(q, id).Scan(
		&u.Id,
		&u.Name,
		&u.Password,
		&u.Email,
		&u.Phone,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)
	return u, err

}

func (r *UserRepository) FindByPhone(phone string) (entity.User, error) {
	var u entity.User
	q := `SELECT * FROM users WHERE phone=$1`
	err := r.db.QueryRow(q, phone).Scan(
		&u.Id,
		&u.Name,
		&u.Password,
		&u.Email,
		&u.Phone,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)
	return u, err
}
