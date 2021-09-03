package database

import (
	"fmt"

	model "go-test-checkout/graph/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type dbConfig struct {
	host     string
	port     string
	dbname   string
	user     string
	password string
}

func getDatabaseUrl() string {

	var config = dbConfig{"localhost", "3306", "be_candidate_test", "root", ""}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.user, config.password, config.host, config.port, config.dbname)
}

func GetDatabase() *gorm.DB {

	db, err := gorm.Open("mysql", getDatabaseUrl())

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	db.LogMode(true)

	db.AutoMigrate(
		&model.User{},
		&model.Item{},
		&model.Cart{},
		&model.Order{},
		&model.OrderDetail{},
	)

	return db
}
