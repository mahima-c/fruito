package model

import "database/sql"

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"` // Omit from JSON output
}

func GetUserById(tx *sql.Tx, userID int) (*User, error) {
	row := tx.QueryRow("SELECT id, name, phone, email FROM users WHERE id = $1", userID)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}

func GetUserByPhone(tx *sql.Tx, phone string) (*User, error) {
	row := tx.QueryRow("SELECT id, name, phone, email FROM users WHERE phone = $1", phone)
	user := &User{}
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return user, nil
}
