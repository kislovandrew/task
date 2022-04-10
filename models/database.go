// Описание методов работы с БД
package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServerConnectionData struct {
	Host string
	Port string
	Name string
	User string
	Pass string
}

var Server ServerConnectionData

var DB *gorm.DB
var Err_connect error

// Возвращает указатель на БД и ошибку в случае неудачного подключения.
// При ошибке подключения пользователь корректно увидит ошибку.
func DBConnect() {
	dsn := "host=" + Server.Host + " port=" + Server.Port + " user=" + Server.User + " password=" + Server.Pass + " dbname=" + Server.Name
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, Err_connect = db, err
}
