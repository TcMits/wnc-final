package datastore

import (
	"database/sql"
	"testing"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

func OpenClient(databaseUrl string, maxPoolSize int, debug bool) (*ent.Client, error) {
	var db *sql.DB
	var err error
	var drv *entsql.Driver
	if debug {
		db, err = sql.Open("pgx", databaseUrl)
	} else {
		db, err = sql.Open("sqlite3", databaseUrl)
	}
	if err != nil {
		return nil, err
	}
	if debug {
		drv = entsql.OpenDB(dialect.Postgres, db)
	} else {
		drv = entsql.OpenDB(dialect.SQLite, db)
	}
	db.SetMaxOpenConns(maxPoolSize)

	// Create an ent.Driver from `db`.
	return ent.NewClient(ent.Driver(drv)), nil
}

func NewClient(url string, poolMax int, debug bool) (*ent.Client, error) {
	return OpenClient(url, poolMax, debug)
}

func openTestConnection(t *testing.T) (*ent.Client, error) {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
	}
	return enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...), nil
}

func NewClientTestConnection(t *testing.T) (*ent.Client, error) {
	return openTestConnection(t)
}
