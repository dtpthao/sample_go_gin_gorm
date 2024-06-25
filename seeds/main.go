package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sample-go-server/entity"
	"sample-go-server/module/user/repository"
	"sample-go-server/utils"
)

func OpenConnection(config entity.MysqlConfig) *gorm.DB {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}

	return db
}

func main() {
	admin := repository.User{
		Uuid:     utils.NewUuid(),
		Username: os.Args[1],
		Password: os.Args[2],
		IsAdmin:  true,
	}

	config := entity.MysqlConfig{
		Host:     "localhost",
		Port:     3307,
		User:     "docker",
		Password: "docker@123",
		DBName:   "dockerdb",
	}
	db := OpenConnection(config)

	err := db.Save(&admin).Error
	if err != nil {
		fmt.Println(err)
	}
}
