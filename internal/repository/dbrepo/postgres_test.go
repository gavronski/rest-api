package dbrepo

import (
	"app/internal/models"
	"app/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

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
var testRepo repository.DatabaseRepo

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
		// delete container
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
		// delete container
		_ = pool.Purge(resource)
		log.Fatalf("could not connect to database: %s", err)
	}

	// populate the database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("error while creating tables: %s", err)
	}

	testRepo = &postgresDBRepo{DB: testDB}

	// run tests
	code := m.Run()

	// delete container after tests
	_ = pool.Purge(resource)

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

func TestPostgresDBRepoInsertPlayer(t *testing.T) {
	player := models.Player{
		ID:        1,
		FirstName: "Tom",
		LastName:  "Stones",
		Age:       21,
		Country:   "Wales",
		Club:      "Chelsea FC",
		Position:  "Right Back",
		Goals:     2,
		Assists:   13,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := testRepo.InsertPlayer(player)

	if err != nil {
		t.Errorf("insert player reports an error %s", err)
	}
}

func TestPostgresDBRepoGetPlayers(t *testing.T) {
	players, err := testRepo.GetPlayers()

	if err != nil {
		t.Errorf("insert player reports an error %s", err)
	}

	if len(players) != 1 {
		t.Errorf("get players reports wrong length; expected 1 but got - %d", len(players))
	}

	player := models.Player{
		ID:        1,
		FirstName: "Tom",
		LastName:  "Stones",
		Age:       21,
		Country:   "Wales",
		Club:      "Chelsea FC",
		Position:  "Right Back",
		Goals:     2,
		Assists:   13,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_ = testRepo.InsertPlayer(player)

	players, err = testRepo.GetPlayers()

	if len(players) != 2 {
		t.Errorf("wrong players length after inserting a player;expected 2 but got %d", len(players))
	}
}

func TestPostgresDBRepoGetMaxID(t *testing.T) {

	id, err := testRepo.GetMaxID()

	if err != nil {
		t.Errorf("max id reports bad retu an error %s", err)
	}

	if id != 2 {
		t.Errorf("max id reports bad return;expected 2 but got - %d", id)
	}
}

func TestPostgresDBRepoGetPlayerByID(t *testing.T) {
	player, err := testRepo.GetPlayerByID(2)

	if err != nil {
		t.Errorf("get player reports an error %s", err)
	}

	if player.ID != 2 {
		t.Errorf("get player reports wrong player id;expected 2 but got %d", player.ID)
	}

	player, err = testRepo.GetPlayerByID(3)

	// max id is 2 so should be an error
	if err == nil {
		t.Error("get player should have reported an error;max id is 2")
	}

	if player.ID != 0 {
		t.Errorf("get player should have returned 0, but got %d", player.ID)
	}
}

func TestPostgresDBRepoUpdatePlayer(t *testing.T) {
	player := models.Player{
		ID:        1,
		FirstName: "Tom",
		LastName:  "Stones",
		Age:       23,
		Country:   "Wales",
		Club:      "Arsenal FC",
		Position:  "Right Back",
		Goals:     2,
		Assists:   13,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := testRepo.UpdatePlayer(player)

	if err != nil {
		t.Errorf("update player reports an error %s", err)
	}

	player, _ = testRepo.GetPlayerByID(player.ID)

	if player.Age != 23 || player.Club != "Arsenal FC" {
		t.Errorf("update player reports an error; player is expected to have age 23 and club Arsenal FC, but got %d and %s", player.Age, player.Club)
	}

	player = models.Player{
		ID:        4,
		FirstName: "Tom",
		LastName:  "Stones",
		Age:       23,
		Country:   "Wales",
		Club:      "Arsenal FC",
		Position:  "Right Back",
		Goals:     2,
		Assists:   13,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = testRepo.UpdatePlayer(player)

	if err == nil {
		t.Errorf("update player report an error; id [%d] is out of the range", player.ID)
	}
}

func TestPostgresDBRepDeletePlayer(t *testing.T) {
	err := testRepo.DeletePlayer(2)

	if err != nil {
		t.Errorf("delete player reports an error %s", err)
	}

	err = testRepo.DeletePlayer(3)

	if err == nil {
		t.Error("delete player should have reported an error; there is no id [3]")
	}
}

func TestPostgresDBRepoAuthenticate(t *testing.T) {
	
}