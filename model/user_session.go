package model

import (
	"database/sql"
	"time"
)

type UserSession struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Token     string    `json:"token"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUserSession(tx *sql.Tx, session *UserSession) error {
	query := `INSERT INTO user_sessions (user_id, token) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err := tx.QueryRow(query, session.UserID, session.Token).Scan(&session.ID, &session.CreatedAt, &session.UpdatedAt)
	return err
}

func GetUserSessionByToken(tx *sql.Tx, token string) (*UserSession, error) {
	session := &UserSession{}
	query := `SELECT id, user_id, created_at, updated_at FROM user_sessions WHERE token = $1`
	err := tx.QueryRow(query, token).Scan(&session.ID, &session.UserID, &session.CreatedAt, &session.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Session not found
		}
		return nil, err
	}
	session.Token = token
	return session, nil
}

func GetLatestUserSessionByUserId(tx *sql.Tx, userId int) (*UserSession, error) {
	sessions := []*UserSession{}
	query := `SELECT id, token, created_at, updated_at FROM user_sessions WHERE user_id = $1 ORDER BY created_at DESC`
	rows, err := tx.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		session := &UserSession{}
		err := rows.Scan(&session.ID, &session.Token, &session.CreatedAt, &session.UpdatedAt)
		if err != nil {
			return nil, err
		}
		session.UserID = userId
		sessions = append(sessions, session)
	}

	if len(sessions) == 0 {
		return nil, nil // No sessions found
	}

	return sessions[0], nil // Return the first session (latest created_at)
}

func UpdateUserSession(tx *sql.Tx, session *UserSession) error {
	query := `UPDATE user_sessions SET updated_at = $1 WHERE id = $2`
	_, err := tx.Exec(query, time.Now(), session.ID)
	return err
}

func DeleteUserSession(tx *sql.Tx, token string) error {
	query := `DELETE FROM user_sessions WHERE token = $1`
	_, err := tx.Exec(query, token)
	return err
}

func DeleteUserSessionByUserId(tx *sql.Tx, userId int) error {
	query := `DELETE FROM user_sessions WHERE user_id = $1`
	_, err := tx.Exec(query, userId)
	return err
}

func DeactivateUserSessionByToken(tx *sql.Tx, token string) error {
	query := `UPDATE user_sessions SET is_active = $1 WHERE token = $2`
	_, err := tx.Exec(query, false, token)
	return err
}
