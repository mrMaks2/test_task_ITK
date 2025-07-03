package envs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var ServerEnvs Envs

type Envs struct {
	POSTGRES_PASSWORD string
	POSTGRES_USER     string
	POSTGRES_PORT     string
	POSTGRES_NAME     string
	POSTGRES_HOST     string
	POSTGRES_USE_SSL  string
	CONN_HOST         string
}

func LoadEnvs() error {

	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Ошибка загрузки файла config.env:", err)
	}

	ServerEnvs.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	ServerEnvs.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	ServerEnvs.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	ServerEnvs.POSTGRES_NAME = os.Getenv("POSTGRES_NAME")
	ServerEnvs.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	ServerEnvs.POSTGRES_USE_SSL = os.Getenv("POSTGRES_USE_SSL")

	ServerEnvs.CONN_HOST = os.Getenv("CONN_HOST")

	return nil

}
