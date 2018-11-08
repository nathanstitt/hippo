package hippo;

import (
	"os"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB(c *cli.Context) *gorm.DB {
	conn := c.String("db_conn_url")
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		fmt.Printf("error connecting using: %s: %s\n", conn, err)
		os.Exit(1)
	}
	return db
}
