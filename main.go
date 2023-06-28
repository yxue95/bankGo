package main

import (
	"database/sql"
	"log"

	"github.com/yxue95/bankGo/util"

	_ "github.com/lib/pq"
	"github.com/yxue95/bankGo/api"
	db "github.com/yxue95/bankGo/db/sqlc"
)

func main() {
	// load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
