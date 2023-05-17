package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/controllers"
	"github.com/kenykendf/go-restful/internal/app/repository"
	"github.com/kenykendf/go-restful/internal/app/service"
	"github.com/kenykendf/go-restful/internal/pkg/config"
	"github.com/kenykendf/go-restful/internal/pkg/middleware"
)

type Server struct {
	cfg    config.Config
	dbConn *sqlx.DB
	router *gin.Engine
}

func NewServer(cfg config.Config, DBConn *sqlx.DB) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		dbConn: DBConn,
	}

	// setup router
	server.setupRouter()

	return server, nil
}

func (svr *Server) setupRouter() {
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

	uploaderService := service.NewUploadService(
		cfg.CloudinaryName,
		cfg.CloudinaryAPIKey,
		cfg.CloudinaryAPISecret,
		cfg.CloudinaryDir,
	)
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo, categoryRepo, uploaderService)
	userService := service.NewUserService(userRepo)
	sessionService := service.NewSessionService(userRepo, authRepository, tokenMaker)

	categoryController := controllers.NewCategoryController(categoryService)
	productController := controllers.NewProductController(productService)
	userController := controllers.NewUserController(userService)
	sessionController := controllers.NewSessionController(sessionService, tokenMaker)

	r.POST("/auth/login", sessionController.Login)

	r.Use(middleware.AuthMiddleware(tokenMaker))
	// categories
	r.POST("/categories", categoryController.CreateCategory)
	r.GET("/categories", middleware.AuthorizationMiddleware("alice", "data1", "read", Enforcer), categoryController.BrowseCategory)
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

	svr.router = r
}

func (svr *Server) Start(address string) error {
	return svr.router.Run(address)
}
