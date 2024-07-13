package main

import (
	"database/sql"
	"log"

	"github.com/ImArnav19/ecom/cmd/api"
	"github.com/ImArnav19/ecom/config"
	"github.com/ImArnav19/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
		Net:                  "tcp",
	})

	if err != nil {
		log.Fatal(err)
	}

	dbinit(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func dbinit(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)

	}
	log.Println("DB Connected")
}
