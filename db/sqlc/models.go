// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	UserID    int32     `json:"user_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int32     `json:"account_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type MoneyRecord struct {
	ID        int32   `json:"id"`
	UserID    int32   `json:"user_id"`
	Reference string  `json:"reference"`
	Status    string  `json:"status"`
	Amount    float64 `json:"amount"`
}

type Transfer struct {
	ID            int64     `json:"id"`
	FromAccountID int32     `json:"from_account_id"`
	ToAccountID   int32     `json:"to_account_id"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type User struct {
	ID             int64          `json:"id"`
	Email          string         `json:"email"`
	HashedPassword string         `json:"hashed_password"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Username       sql.NullString `json:"username"`
}
