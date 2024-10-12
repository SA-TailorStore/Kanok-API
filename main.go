package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SA-TailorStore/Kanok-API/configs"
	"github.com/SA-TailorStore/Kanok-API/database/adapter/mysql"
	"github.com/SA-TailorStore/Kanok-API/domain/controllers"
	"github.com/SA-TailorStore/Kanok-API/domain/services"
	"github.com/cloudinary/cloudinary-go/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func main() {
	app := fiber.New()

	// Cloudinary
	cfg := configs.NewConfig()
	cld, err := cloudinary.NewFromURL(cfg.Cloudinary_url)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	fmt.Printf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sqlx.ConnectContext(ctx, "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepo := mysql.NewUserMySQL(db)
	userService := services.NewUserService(userRepo, cfg, cld)
	userController := controllers.NewUserController(userService)

	orderRepo := mysql.NewOrderMySQL(db)
	orderService := services.NewOrderService(orderRepo, cfg)
	orderController := controllers.NewOrderController(orderService)

	productRepo := mysql.NewProductMySQL(db)
	productService := services.NewProductService(productRepo, cfg)
	productController := controllers.NewProductController(productService)

	designRepo := mysql.NewDesignMySQL(db)
	designService := services.NewDesignService(designRepo, cfg, cld)
	designController := controllers.NewDesignController(designService)

	fabricRepo := mysql.NewFabricMySQL(db)
	fabricService := services.NewFabricService(fabricRepo, cfg, cld)
	fabricController := controllers.NewFabricController(fabricService)

	materialRepo := mysql.NewMaterialMySQL(db)
	materialService := services.NewMaterialService(materialRepo, cfg)
	materialController := controllers.NewMaterialController(materialService)

	prefix := "/api"
	// api routes post

	// User
	app.Post(prefix+"/register", userController.Register)
	app.Post(prefix+"/login", userController.Login)
	app.Post(prefix+"/login/token", userController.LoginToken)
	app.Post(prefix+"/user/token", userController.GetUserByJWT)
	app.Post(prefix+"/user/update/address", userController.UpdateAddress)
	app.Post(prefix+"/user/profile/upload", userController.UpdateImage)

	// Order
	app.Post(prefix+"/order/create", orderController.CreateOrder)
	app.Post(prefix+"/order/get", orderController.GetOrderByID)
	app.Post(prefix+"/order/user", orderController.GetOrderByJWT)
	app.Post(prefix+"/order/update/status", orderController.UpdateStatus)
	app.Post(prefix+"/order/update/payment", orderController.UpdatePayment)

	// Product
	app.Post(prefix+"/product/create", productController.CreateProduct)
	app.Post(prefix+"/product/get", productController.GetProductByID)
	app.Post(prefix+"/product/get/order", productController.GetProductByOrderID)

	// Design
	app.Post(prefix+"/design/add", designController.AddDesign)
	app.Post(prefix+"/design/update", designController.UpdateDesign)
	app.Post(prefix+"/design/delete", designController.DeleteDesign)
	app.Post(prefix+"/design/get", designController.GetDesignByID)

	// Fabric
	app.Post(prefix+"/fabric/add", fabricController.AddFabric)
	app.Post(prefix+"/fabric/update", fabricController.UpdateFabric)
	app.Post(prefix+"/fabric/updates", fabricController.UpdateFabrics)
	app.Post(prefix+"/fabric/delete", fabricController.DeleteFabric)
	app.Post(prefix+"/fabric/get", fabricController.GetFabricByID)

	// Material
	app.Post(prefix+"/material/add", materialController.AddMaterial)
	app.Post(prefix+"/material/update", materialController.UpdateMaterial)
	app.Post(prefix+"/material/delete", materialController.DeleteMaterial)
	app.Post(prefix+"/material/get", materialController.GetMaterialByID)

	// api routes get
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// User
	app.Get(prefix+"/users", userController.FindAllUser)

	// Order
	app.Get(prefix+"/orders", orderController.GetAllOrders)

	// Product

	// Design
	app.Get(prefix+"/designs", designController.GetAllDesigns)

	// Fabric
	app.Get(prefix+"/fabrics", fabricController.GetAllFabrics)

	// Material
	app.Get(prefix+"/materials", materialController.GetAllMaterials)

	if err := app.Listen(":9000"); err != nil {
		log.Fatal(err)
	}
}
