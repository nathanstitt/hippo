package hippo;

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/nathanstitt/hippo/models"
	"github.com/volatiletech/sqlboiler/boil"
)

type DB = *sql.Tx

// as a convience re-export models/func from models
// so that other users of Hippo don't have to import models
type Tenant = hm.Tenant
var Tenants = hm.Tenants
var FindTenantP = hm.FindTenantP

type User = hm.User
var Users = hm.Users
var FindUserP = hm.FindUserP

type Subscription = hm.Subscription
var Subscriptions = hm.Subscriptions

func ConnectDB(c Configuration) *sql.DB {
	conn := c.String("db_connection_url")
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(fmt.Sprintf("invalid syntx for db_conn_url %s: %s\n", conn, err))
	}
	pingErr := db.Ping()
	if pingErr != nil {
		panic(fmt.Sprintf("unable to connect to DB using %s: %s\n", conn, pingErr))
	}

	hm.AddUserHook(boil.BeforeDeleteHook, ensureOwnerAndGuest)

	return db
}
