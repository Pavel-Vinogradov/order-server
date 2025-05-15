package repository

import (
	"database/sql"
	entityAuth "order-server/internal/auth/entity"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) Create(auth *entityAuth.Auth) (*entityAuth.Auth, error) {
	query := `INSERT INTO auth (user_id, session) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRow(query, auth.UserId, auth.Session).
		Scan(&auth.Id)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (r *AuthRepository) FindByUserId(userID int64) (*entityAuth.Auth, error) {
	query := `SELECT * FROM auth WHERE user_id = $1`
	var auth entityAuth.Auth
	err := r.db.QueryRow(query, userID).
		Scan(&auth.Id, &auth.UserId, &auth.Session)
	return &auth, err
}

func (r *AuthRepository) UpdateSession(userId int64, session string) (*entityAuth.Auth, error) {
	query := `
         UPDATE auth
         SET session = $1
         WHERE user_id = $2
         RETURNING id, user_id, session
	`
	var updated entityAuth.Auth
	err := r.db.QueryRow(query, session, userId).
		Scan(&updated.Id, &updated.UserId, &updated.Session)
	return &updated, err
}

func (r *AuthRepository) FindBySession(sessionId string) (*entityAuth.Auth, error) {
	query := `SELECT * FROM auth WHERE session = $1`
	var auth entityAuth.Auth
	err := r.db.QueryRow(query, sessionId).
		Scan(&auth.Id, &auth.UserId, &auth.Session)
	return &auth, err
}
