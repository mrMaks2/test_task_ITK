package server

import (
	"log"
	"test_task_ITK/database"
	"test_task_ITK/envs"
	"test_task_ITK/models"
)

func InitServer() {

	errEnvs := envs.LoadEnvs()
	if errEnvs != nil {
		log.Fatal("Ошибка загрузки ENV: ", errEnvs)
	} else {
		log.Println("Успешное получение ENV")
	}

	errDatabase := database.InitDatabase()
	if errDatabase != nil {
		log.Fatal("Ошибка подключения к базе данных: ", errDatabase)
	} else {
		log.Println("Успешное подключение к базе данных")
		database.DB.AutoMigrate(&model.Wallet{})
	}

}

func StartServer() {
	InitRoutes()
}
