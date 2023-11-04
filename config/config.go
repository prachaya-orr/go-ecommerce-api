package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}
	return &config{
		app: &app{},
		db:  &db{},
		jwt: &jwt{},
	}

}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type IAppConfig interface {
}
type app struct {
	host         string
	port         string
	name         string
	version      string
	readTimeOut  time.Duration
	writeTimeOut time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpbucket    string
}

func (c *config) App() IAppConfig {
	return nil

}

func (c *config) Db() IDbConfig {
	return nil

}

type IDbConfig interface {
}
type db struct {
	host           string
	port           string
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections string
}

type IJwtConfig interface {
}
type jwt struct {
	adminKey         string
	secretKey        string
	apiKey           string
	accessExpiresAt  string
	refreshExpiresAt string
}

func (c *config) Jwt() IJwtConfig {
	return nil

}
