package driver

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	PgConn *pgx.Conn
}

var Driver = &DB{
	PgConn: &pgx.Conn{},
}

// Function ConnectPg initialize DB connection.
func ConnectPg(connString string) error {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return err
	}
	Driver.PgConn = conn
	return nil
}
