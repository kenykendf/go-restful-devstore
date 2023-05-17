package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/pkg/config"
	"github.com/kenykendf/go-restful/internal/pkg/db"

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
	// init server
	server, err := NewServer(cfg, DBConn)
	if err != nil {
		log.Panic("cannot init server")
	}

	// start server
	appPort := fmt.Sprintf(":%s", cfg.ServerPort)
	err = server.Start(appPort)
	if err != nil {
		log.Panic(fmt.Errorf("error cannot start app : %w", err))
	}
}
