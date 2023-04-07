package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(DBDriver string, DBConn string) (*sqlx.DB, error) {
	db, err := sqlx.Connect(DBDriver, DBConn)
	if err != nil {
		fmt.Println("IF ERR HERE ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Database connection status up.")
	return db, nil
}
