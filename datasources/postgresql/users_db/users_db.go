package users_db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
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
	Client   *pgxpool.Pool
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
	var err error
	loadEnvironment()

	username = os.Getenv(pgUsername)
	password = os.Getenv(pgPassword)
	host = os.Getenv(pgHost)
	schema = os.Getenv(pgSchema)

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s",
		username, password, host, schema,
	)

	Client, err = pgxpool.Connect(context.Background(), dataSourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(context.Background()); err != nil {
		panic(err)
	}

	log.Println("Database successfully set up")
}
