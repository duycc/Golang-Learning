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

func (f Food) TableName() string {
	return "foods"
}

func main() {
	// fmt.Println("hello gorm...")
	db, err := getDB()
	if err != nil {
		panic("conn mysql failed, error: " + err.Error())
	}

	// insertData(db)

	// queryData(db)

	// updateRecords(db)

	deleteRecords(db)

	defer db.Close()
}

func deleteRecords(db *gorm.DB) {
	db.Where("type = ?", 11).Delete(&Food{})
}

func updateRecords(db *gorm.DB) {
	fmt.Println("updateRecords test...")
	// food := Food{}
	// db.Where("id = ?", 2).Take(&food)
	// food.Price = 100
	// db.Save(&food)

	// if err := db.Model(&Food{}).Where("id = ?", 2).Update("price", 25).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }

	// db.Model(&Food{}).Update("price", 100)

	// updateFood := Food{
	// 	Price: 120,
	// 	Title: "柠檬雪碧",
	// }

	// if err := db.Model(&Food{}).Updates(&updateFood).Error; err != nil {
	// 	fmt.Printf("update failed, err: %+v\n", err)
	// 	return
	// }

	upd := make(map[string]interface{})
	upd["stock"] = 0
	upd["price"] = 200

	db.Model(&Food{}).Where("id = ?", 2).Updates(upd)
	db.Model(&Food{}).Where("id = ?", 2).Update("stock", gorm.Expr("stock + 1"))
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

	//

	// err := db.Where("id = ?", 2).Take(&food).Error
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	fmt.Println("record not found.")
	// 	return
	// } else if err != nil {
	// 	fmt.Println("query failed.")
	// 	return
	// }
	// if err := db.Where("id = ?", 2).Take(&food).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }

	// if err := db.Where("id in (?)", []int{1, 2}).Take(&food).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }

	// foods := make([]Food, 10)

	// if err := db.Where("title like ?", "%an%").Find(&foods).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }
	// fmt.Println("query ok.")
	// fmt.Println(foods)

	// food := Food{}
	// // if err := db.Select("id, title").Where("id = ?", 1).Take(&food).Error; err != nil {
	// if err := db.Select([]string{"id", "title"}).Where("id = ?", 1).Take(&food).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }
	// fmt.Println(food)

	// total := []int{}
	// if err := db.Model(&Food{}).Select("count(*) as total").Pluck("total", &total).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }
	// fmt.Println(total)

	// foods := []Food{}
	// if err := db.Where("id in (?)", []int{1, 2}).Order("create_time desc").Find(&foods).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }

	// if err := db.Order("create_time desc").Limit(1).Offset(0).Find(&foods).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }
	// fmt.Println(foods)

	// var total int64
	// db.Model(&Food{}).Count(&total)
	// fmt.Println("count: ", total)

	results := []Result{}
	// if err := db.Model(&Food{}).Select("type, count(*) as total").Group("type").Having("total > 0").Scan(&results).Error; err != nil {
	// 	fmt.Printf("query failed, err: %+v\n", err)
	// 	return
	// }
	// fmt.Println(results)

	// 直接执行sql语句
	sql := "select type, count(*) as total from `foods` where create_time > ? group by type having (total > 0)"
	db.Raw(sql, "2022-02-21 01:04:07").Scan(&results)
	fmt.Println(results)
}

type Result struct {
	Type  int
	Total int
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
