package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	pgUsername = "PG_USERNAME"
	pgPassword = "PG_PASSWORD"
	pgHost     = "PG_HOST"
	pgSchema   = "PG_SCHEMA"
)

var (
	Client *sql.DB

	username string
	password string
	host     string
	schema   string
)

func loadEnvironment() {
	err := godotenv.Load("./datasources/postgresql/users_db/.env")

	if err != nil {
		log.Println("Couldn't load environment variables")
		panic(err)
	}
}

func init() {
	//var err error

	loadEnvironment()

	username = os.Getenv(pgUsername)
	password = os.Getenv(pgPassword)
	host = os.Getenv(pgHost)
	schema = os.Getenv(pgSchema)

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
}
