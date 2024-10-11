package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	db *sqlx.DB
}

func NewMySQL() *sqlx.DB {
	cfg := configs.NewConfig()
	ctx := context.Background()

	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return db
}
