package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"glintecoTask/entity"
	delivery2 "glintecoTask/module/contract/delivery"
	repository2 "glintecoTask/module/contract/repository"
	usecase2 "glintecoTask/module/contract/usecase"
	tokenDelivery "glintecoTask/module/token/delivery"
	tokenUC "glintecoTask/module/token/usecase"
	"glintecoTask/module/user/delivery"
	"glintecoTask/module/user/repository"
	"glintecoTask/module/user/usecase"
	apiLog "glintecoTask/utils/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var logConfig apiLog.Config
var mysqlConfig entity.MysqlConfig
var serverConfig entity.ServerConfig

func config() {
	err := cleanenv.ReadEnv(&logConfig)
	if err != nil {
		log.Fatalf("Read log config failed. %v", err)
	}

	err = apiLog.InitLog(logConfig)
	if err != nil {
		log.Fatalf("Init log failed. %v", err)
	}

	// todo no hardcode
	mysqlConfig = entity.MysqlConfig{
		Host:     "localhost",
		Port:     3307,
		User:     "glinteco",
		Password: "glinteco@123",
		DBName:   "glintecodb",
	}

	serverConfig = entity.ServerConfig{
		Host:      "localhost",
		Port:      3030,
		JWTSecret: "5b0b18dc37004b97946367ca5d82673918a6c6e7a817bf84236abe1c0907b9bf",
	}
}

func main() {

	config()

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

	jwtSecret, _ := hex.DecodeString(serverConfig.JWTSecret)
	tokenUseCase := tokenUC.NewTokenUseCase(jwtSecret)

	app := gin.New()
	app.Use(apiLog.AccessLog)

	// user
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, tokenUseCase)

	tokenHandler := tokenDelivery.NewTokenHandler(tokenUseCase, userUseCase, jwtSecret)

	userHandler := delivery.NewUserHandler(app, userUseCase)
	userHandler.RegisterHandler(tokenHandler.Authenticate, tokenHandler.AdminAuthorize)

	// contract
	cRepository := repository2.NewContractRepository(db)
	cUseCase := usecase2.NewContractUseCase(cRepository)
	cHandler := delivery2.NewContractHandler(app, cUseCase)
	cHandler.RegisterHandler(tokenHandler.Authenticate)

	addr := fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port)
	apiLog.Info().Msg("Server start at: http://" + addr)
	err = app.Run(addr)

	if err != nil {
		panic(err)
	}

}
