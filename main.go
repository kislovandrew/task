// Решение тестовой задачи для получения данных из указанной БД
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kislovandrew/task/controllers"
	m "github.com/kislovandrew/task/models"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Загрузка данных для подключения из env
	m.Server.Host = os.Getenv("DB_HOST")
	m.Server.Port = os.Getenv("DB_PORT")
	m.Server.Name = os.Getenv("DB_NAME")
	m.Server.User = os.Getenv("DB_USER")
	m.Server.Pass = os.Getenv("DB_PASS")

	m.DBConnect()

	router := gin.Default()
	router.GET("api/v1/item6/", controllers.GetDBData)
	router.Run(":8080")
}
