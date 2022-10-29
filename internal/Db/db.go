package Db

import (
	"database/sql"
	"fmt"
	"github.com/bmsandoval/flagr/internal/Utilities/AppContext"
	"github.com/bmsandoval/flagr/pkg/DbContext"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

var DB *sql.DB

func Start(ctx AppContext.Context) (*DbContext.Connection, error) {
	// Open and ping database
	if err := open(ctx); err != nil {
		return nil, err
	}
	// Run UP migrations
	if err := migrate(ctx); err != nil {
		return nil, err
	}

	return &DbContext.Connection{
		DB: DB,
	}, nil
}

func Stop() error {
	return DB.Close()
}

func open(ctx AppContext.Context) error {
	var err error

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		ctx.Config.DbUser,
		ctx.Config.DbPass,
		ctx.Config.DbHost,
		ctx.Config.DbPort,
		ctx.Config.DbSchema)

	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	if err := ping(); err != nil {
		return err
	}

	return nil
}

func ping() error {
	return DB.Ping()
}

func migrate(ctx AppContext.Context) error {
	if err := goose.SetDialect("mysql"); err != nil {
		return err
	}
	if err := goose.Up(DB, ctx.Config.DbMigrationLocation); err != nil {
		return err
	}
	return nil
}
