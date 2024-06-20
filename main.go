package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
	usecase2 "glintecoTask/module/token/usecase"
	"glintecoTask/module/user/delivery"
	"glintecoTask/module/user/repository"
	"glintecoTask/module/user/usecase"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {

	// todo read config instead of hardcode
	mysqlConfig := entity.MysqlConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "glinteco",
		Password: "glinteco@123",
		DBName:   "glinteco_db",
	}

	serverConfig := entity.ServerConfig{
		Host:      "localhost",
		Port:      3030,
		JWTSecret: "",
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)
	gormConfig := gorm.Config{}
	gormConfig.DisableNestedTransaction = true
	gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	db, err := gorm.Open(mysql.Open(dns), &gormConfig)
	if err != nil {
		log.Fatalf("Connect to MySQL server failed. %v", err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Cannot config MySQL connection. %v", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	app := gin.New()
	userRepository := repository.NewUserRepository(db)
	tokenUseCase := usecase2.NewTokenUseCase(serverConfig.JWTSecret, userRepository)
	userUseCase := usecase.NewUserUseCase(userRepository, tokenUseCase)
	userHandler := delivery.NewUserHandler(app, userUseCase)
	userHandler.RegisterHandler()

	addr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	err = app.Run(addr)
	if err != nil {
		panic(err)
	}
}
