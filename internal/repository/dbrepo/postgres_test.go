package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/ory/dockertest/v3"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "root"
	dbName   = "players_test"
	port     = "5433"
	dsn      = "host=%s port=%s user=%s password=%d dbname=%s sslmode=disable"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB

func TestMain(m *testing.M) {
	// connect to docker
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	pool = p

	// set up docker options
	options := dockertest.RunOptions{
		Repository: "postgres",
		Env: []string{
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_USER=" + user,
			"POSTGRES_DB=" + dbName,
			"listen_addresses = '*'",
		},
	}

	// get a resource (instance of an docker image)
	resource, err = pool.RunWithOptions(&options)

	if err != nil {
		_ = pool.Purge(resource)
		log.Fatal("could not start resource", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, hostAndPort, dbName)

	// start the image wait until it's ready
	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", databaseUrl)
		if err != nil {
			log.Println("Error:", err)
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to database: %s", err)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("error while creating tables: %s", err)
	}

	// run tests
	code := m.Run()
	log.Println("hello world")
	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/players.sql")

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
