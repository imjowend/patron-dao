package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/imjowend/patron-dao/utilities"
)

func get() *sql.DB {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
