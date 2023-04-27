package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/controllers"
	"github.com/kenykendf/go-restful/internal/app/repository"
	"github.com/kenykendf/go-restful/internal/app/service"
	"github.com/kenykendf/go-restful/internal/pkg/config"
	"github.com/kenykendf/go-restful/internal/pkg/db"
	"github.com/kenykendf/go-restful/internal/pkg/middleware"

	log "github.com/sirupsen/logrus"
)

var cfg config.Config
var DBConn *sqlx.DB

func init() {

	// read configuration,
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load app config")
		return
	}
	cfg = configLoad

	// connect database
	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		fmt.Println("db unavailable")
		return
	}
	DBConn = db

	// setup logrus
	logLevel, err := log.ParseLevel("debug")
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)                 // apply log level
	log.SetFormatter(&log.JSONFormatter{}) // define format using json

}

func main() {
	// using default gin logger
	// r := gin.Default()

	// using default gin logger
	r := gin.New()

	// enable middleware
	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ping"})
	})

	categoryRepo := repository.NewCategoryRepo(DBConn)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	// categories
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories", categoryController.BrowseCategory)
	r.GET("/categories/:id", categoryController.DetailCategory)
	r.PUT("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	productRepo := repository.NewProductRepo(DBConn)
	productService := service.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// products
	r.POST("/products", productController.Create)
	r.GET("/products", productController.BrowseProduct)
	// r.GET("/products/:id", productController)
	// r.PUT("/products/:id", productController)
	// r.DELETE("/products/:id", productController)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
