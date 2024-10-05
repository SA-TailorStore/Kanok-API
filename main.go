package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/adapter/mysql"
	"github.com/SA-TailorStore/Kanok-API/domain/controllers"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
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

	userRepo := mysql.NewUserMySQL(db)
	userService := services.NewUserService(userRepo, cfg)
	userController := controllers.NewUserController(userService)

	orderRepo := mysql.NewOrderMySQL(db)
	orderService := services.NewOrderService(orderRepo, cfg)
	orderController := controllers.NewOrderController(orderService)

	productRepo := mysql.NewProductMySQL(db)
	productService := services.NewProductService(productRepo, cfg)
	productController := controllers.NewProductController(productService)

	// api routes post
	// User
	app.Post("/register", userController.Register)
	app.Post("/login", userController.Login)
	app.Post("/login-token", userController.GetUserByJWT)
	// Order
	app.Post("/create-order", orderController.CreateOrder)
	// Product
	app.Post("/create-product", productController.CreateProduct)

	// api routes get
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// User
	app.Get("/users", userController.FindAllUser)

	// Order

	// Product

	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}
