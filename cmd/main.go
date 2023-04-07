package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/controllers"
	"github.com/kenykendf/go-restful/internal/app/repository"
	"github.com/kenykendf/go-restful/internal/app/service"
	"github.com/kenykendf/go-restful/internal/pkg/config"
	"github.com/kenykendf/go-restful/internal/pkg/db"
)

var cfg config.Config
var DBConn *sqlx.DB

func init() {

	// read configuration,
	configLoad, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("cannot load app config")
	}
	cfg = configLoad

	// connect database
	db, err := db.ConnectDB(cfg.DBDriver, cfg.DBConnection)
	if err != nil {
		log.Panic("db unavailable")
	}
	DBConn = db
}

func main() {

	fmt.Println("PRINT CFG DBConnection ", cfg.DBConnection)
	fmt.Println("PRINT CFG DBDriver ", cfg.DBDriver)
	fmt.Println("PRINT CFG ServerPort ", cfg.ServerPort)
	r := gin.Default()

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

	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	r.Run(appPort)
}
