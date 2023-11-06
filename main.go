package main

import (
	"fmt"
	"os"

	"github.com/prachaya-or/eCommerce-api/config"
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
	fmt.Printf("%#v\n", cfg.App())
	fmt.Printf("%#v\n", cfg.Db())
	fmt.Printf("%#v\n", cfg.Jwt())
}
