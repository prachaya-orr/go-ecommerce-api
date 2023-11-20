package main

import (
	"os"

	"github.com/prachaya-or/eCommerce-api/config"
	"github.com/prachaya-or/eCommerce-api/modules/servers"
	"github.com/prachaya-or/eCommerce-api/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}
