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

// Open new connection.
func Open(databaseUrl string, maxPoolSize int) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxPoolSize)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

// NewClient returns an orm client.
func NewClient(url string, poolMax int) (*ent.Client, error) {
	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	return Open(url, poolMax)
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
