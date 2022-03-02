package connection

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func InitializeDatabaseConnection() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalln("Could not establish connection to database host: " + err.Error())
	}

	return conn
}
