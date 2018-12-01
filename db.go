package hippo;

import (
	"os"
	"fmt"
	"database/sql"
)

type DB = *sql.Tx

func ConnectDB(c Configuration) *sql.DB {
	conn := c.String("db_conn_url")
	db, err := sql.Open("postgres", conn)

	if err != nil {
		fmt.Printf("error connecting using: %s: %s\n", conn, err)
		os.Exit(1)
	}
	return db
}
