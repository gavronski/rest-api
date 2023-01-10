package dbrepo

import (
	"app/internal/repository"
	"database/sql"
)

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB) repository.DatabaseRepo {
	return &postgresDBRepo{
		DB: conn,
	}
}

// testDBREpo created to test hanlders
type testDBRepo struct {
}

// NewTestingRepo sets testDBRepo
func NewTestingRepo() repository.DatabaseRepo {
	return &testDBRepo{}
}
