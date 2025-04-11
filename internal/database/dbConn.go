package database

import (
	"database/sql"
	"fmt"
	"log"
	"test-sharing-vision/internal/config"
	"test-sharing-vision/service/model"
	"time"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var (
	ArticleDB *gorm.DB
)

func init() {
	log.Println("[DATABASE] Connecting start")

	var err error
	articleConn := fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Cfg.Database.User,
		config.Cfg.Database.Password,
		config.Cfg.Database.Host,
		config.Cfg.Database.Port,
		config.Cfg.Database.Name)
	ArticleDB, err = gorm.Open(mysql.Open(articleConn))
	if err != nil {
		log.Fatalf("could not connect to the database article: %v", err)
	}

	var mysqlDb *sql.DB
	mysqlDb, err = ArticleDB.DB()
	if err != nil {
		log.Fatalf("could not return sql.DB database article: %v", err)
	}

	mysqlDb.SetMaxOpenConns(100)
	mysqlDb.SetMaxIdleConns(10)
	mysqlDb.SetConnMaxIdleTime(time.Hour)
	mysqlDb.SetConnMaxLifetime(time.Hour)
	log.Println("[DATABASE] Connected")
}

func SetupDB() error {
	log.Println("[DATABASE] AutoMigrate Start")
	err := ArticleDB.AutoMigrate(&model.Posts{})
	if err != nil {
		return err
	}
	log.Println("[DATABASE] AutoMigrate End")

	return nil
}
