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

	"github.com/casbin/casbin/v2"
	log "github.com/sirupsen/logrus"
)

var (
	cfg      config.Config
	DBConn   *sqlx.DB
	Enforcer *casbin.Enforcer
)

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
		log.Panic(err)
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

	// setup casbin
	e, err := casbin.NewEnforcer("config/model.conf", "config/policy.csv")
	if err != nil {
		log.Panic("cannot init casbin")
	}
	Enforcer = e
}

// nolint
func main() {
	// using default gin logger
	// r := gin.Default()
	// using default gin logger
	r := gin.New()

	// enable middleware
	r.Use(
		middleware.LoggingMiddleware(),
		// handle panic return
		middleware.RecoveryMiddleware(),
	)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ping"})
	})

	categoryRepo := repository.NewCategoryRepo(DBConn)
	productRepo := repository.NewProductRepo(DBConn)
	userRepo := repository.NewUserRepository(DBConn)
	authRepository := repository.NewAuthRepo(DBConn)

	tokenMaker := service.NewGenerateToken(
		cfg.AccessTokenKey,
		cfg.RefreshTokenKey,
		cfg.AccessTokenDuration,
		cfg.RefreshTokenDuration,
	)

	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(userRepo, authRepository, tokenMaker)

	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)
	userController := controllers.NewUserController(userService)
	sessionController := controllers.NewSessionController(sessionService)

	r.POST("/auth/login", sessionController.Login)

	r.Use(middleware.AuthMiddleware(tokenMaker))
	// categories
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories", categoryController.BrowseCategory)
	r.GET("/categories/:id", categoryController.DetailCategory)
	r.PUT("/categories/:id", categoryController.UpdateCategory)
	r.DELETE("/categories/:id", categoryController.DeleteCategory)

	// products
	r.POST("/products", productController.Create)
	r.GET("/products", productController.BrowseProduct)
	// r.GET("/products/:id", productController)
	// r.PUT("/products/:id", productController)
	// r.DELETE("/products/:id", productController)

	// users
	r.POST("/user", userController.Create)
	// r.GET("/users", userController.Create)
	// r.GET("/users/:id", userController.Create)
	// r.PUT("/users/:id", userController.Create)
	// r.DELETE("/users/:id", userController.Create)

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
