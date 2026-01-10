package database

import (
	"database/sql"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
)

func EntClient() *ent.Client {
	dsn := os.Getenv("DB_URL")
	db, _ := sql.Open("pgx", dsn)
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}
