package user

import (
	"context"

	"github.com/jackc/pgx/v4"

	userEntity "github.com/andfxx27/chirps-api/domain/user/entity"
)

type repository struct {
	db *pgx.Conn
}

type IRepository interface {
	CreateUser(user userEntity.UserDBEntity) error
	FindUserByUsernameOrEmail(username, email string) (userEntity.UserDBEntity, error)
}

func NewRepository(db *pgx.Conn) IRepository {
	return &repository{db: db}
}

func (repo *repository) CreateUser(user userEntity.UserDBEntity) error {
	_, err := repo.db.Exec(context.Background(), `
		INSERT INTO users
		(first_name, last_name, username, email, password)
		VALUES
		($1, $2, $3, $4, $5)
	`, user.FirstName, user.LastName, user.Username, user.Email, user.Password)
	return err
}

func (repo *repository) FindUserByUsernameOrEmail(username, email string) (userEntity.UserDBEntity, error) {
	row := repo.db.QueryRow(context.Background(), `
		SELECT * FROM users
		WHERE
		username = $1
		OR
		email = $2
		AND
		deleted_at IS NULL
		LIMIT 1
	`, username, email)

	user := userEntity.UserDBEntity{}
	err := user.ScanRow(row)
	return user, err
}
