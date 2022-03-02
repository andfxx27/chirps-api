package entity

import (
	"time"

	"github.com/jackc/pgx/v4"
)

type UserDBEntity struct {
	ID        int
	FirstName string
	LastName  *string
	Username  string
	Email     string
	Password  string
	Status    *string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (user *UserDBEntity) ScanRow(row pgx.Row) error {
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)

	return err
}
