package follow

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
)

type repository struct {
	db *pgx.Conn
}

type IRepository interface {
	CreateFollow(followerID, followedID int) error
	DeleteFollow(followerID, followedID int) error
}

func NewRepository(db *pgx.Conn) IRepository {
	return &repository{db: db}
}

func (repo *repository) CreateFollow(followerID, followedID int) error {
	_, err := repo.db.Exec(context.Background(), `
		INSERT INTO follows
		(follower_id, followed_id)
		VALUES
		($1, $2)
		ON CONFLICT ON CONSTRAINT follows_pkey
		DO
			UPDATE SET deleted_at = $3
	`, followerID, followedID, nil)
	return err
}

func (repo *repository) DeleteFollow(followerID, followedID int) error {
	_, err := repo.db.Exec(context.Background(), `
		UPDATE follows
		SET 
		deleted_at = $1
		WHERE
		follower_id = $2
		AND
		followed_id = $3
	`, time.Now().UTC(), followerID, followedID)
	return err
}
