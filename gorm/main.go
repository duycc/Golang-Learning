package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 商品
type Food struct {
	Id         int
	Title      string
	Price      float32
	Stock      int
	Type       int
	CreateTime time.Time
}

func (v Food) TableName() string {
	return "foods"
}

func main() {
	fmt.Println("hello gorm...")
	db, err := getDB()
	if err != nil {
		panic("conn mysql failed, error: " + err.Error())
	}

	// insertData(db)

	queryData(db)

	defer db.Close()
}

func queryData(db *gorm.DB) {
	// food := Food{}
	// if err := db.Take(&food).Error; err != nil {
	// 	fmt.Println("query failed.")
	// 	return
	// }

	// if err := db.First(&food).Error; err != nil {
	// 	fmt.Println("query failed.")
	// 	return
	// }

	// if err := db.Last(&food).Error; err != nil {
	// 	fmt.Println("query failed.")
	// 	return
	// }

	// var foods []Food
	// if err := db.Find(&foods).Error; err != nil {
	// 	fmt.Println("query failed.")
	// 	return
	// }

	var titles []string
	if err := db.Model(&Food{}).Pluck("title", &titles).Error; err != nil {
		fmt.Println("query failed.")
		return
	}

	fmt.Println("query ok.")
	fmt.Println(titles)
}

func insertData(db *gorm.DB) {
	food := Food{
		Title:      "liangpi",
		Price:      12.34,
		Stock:      2,
		Type:       22,
		CreateTime: time.Now(),
	}
	if err := db.Create(&food).Error; err != nil {
		fmt.Println("insert failed, err: ", err.Error())
		return
	}

	fmt.Println("insert success.")
}

func getDB() (*gorm.DB, error) {
	username := "root"
	password := "duyong"
	address := "localhost"
	dbName := "gorm"
	timeout := "10s"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, address, dbName, timeout)
	db, err := gorm.Open("mysql", dsn)

	return db, err
}
