package config

import (
	"github.com/joho/godotenv"
	"log"
	"strconv"
)

type Configuration struct {
	Server       SetupServer
	Database     SetupDatabase
	AuthKey      SetupAccessKey
	SuperAccount SetupSuperAdmin
}

type SetupServer struct {
	Port string
}

type SetupDatabase struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort string
}

type SetupSuperAdmin struct {
	SuperName     string
	SuperPassword string
}

type SetupAccessKey struct {
	SecretKey string
	ExpiresAt int
}

var (
	Config *Configuration
)

func SetupConfiguration() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Config = new(Configuration)
	Config.Server.Port = envs["PORT"]
	Config.Database.DBHost = envs["DB_HOST"]
	Config.Database.DBPort = envs["DB_PORT"]
	Config.Database.DBName = envs["DB_NAME"]
	Config.Database.DBUser = envs["DB_USER"]
	Config.Database.DBPass = envs["DB_PASS"]
	Config.AuthKey.SecretKey = envs["SECRET_KEY"]
	Config.AuthKey.ExpiresAt, _ = strconv.Atoi(envs["EXPIRED_AT"])
	Config.SuperAccount.SuperName = envs["SUPER_NAME"]
	Config.SuperAccount.SuperPassword = envs["SUPER_PASSWORD"]
}
