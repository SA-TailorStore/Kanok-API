package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SA-TailorStore/Kanok-API/configs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	app := fiber.New()

	ctx := context.Background()

	cfg := configs.NewConfig()
	fmt.Printf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	fmt.Println()
	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// api routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}
