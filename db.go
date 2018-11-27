package hippo;

import (
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB = *gorm.DB

func ConnectDB(c Configuration) DB {
	conn := c.String("db_conn_url")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		fmt.Printf("error connecting using: %s: %s\n", conn, err)
		os.Exit(1)
	}
	return db
}
