package postgres

import (
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SslMode  bool
	Verbose  bool
}

func Connect(config DBConfig) *bun.DB {
	ssl := "disable"
	if config.SslMode {
		ssl = "enable"
	}

	// Open a PostgresSQL database.
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		ssl,
	)
	pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgDb, pgdialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return db
}
