package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wfm_api/internal/model"
	"wfm_api/internal/repo"
)


func main() {
	//dsn := "host=localhost user=wfm_user password=wfm_pa$$wd dbname=wfm_db port=5433"
	configPath := "./config/config.yaml"
	dsn := repo.GetDBConfig(configPath)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err !=  nil {
		log.Fatalln("Unable to connect database ", err)
	}
	err = db.AutoMigrate(&model.User{})
	err = db.AutoMigrate(&model.Cart{})
	err = db.AutoMigrate(&model.Good{})
	err = db.AutoMigrate(&model.CartGood{})

	if err !=  nil {
		log.Fatalln("Auto migration failed ", err)
	}

	router := gin.Default()



}
