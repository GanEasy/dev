package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
)

//User 用户
type User struct {
	gorm.Model
	TID string
	Tag pq.Int64Array `gorm:"type:integer[]"`
}

func main() {
	db, _ := gorm.Open("postgres", "host=localhost user=postgres dbname=gorm sslmode=disable password=123456")
	defer db.Close()

	db.LogMode(true)

	db.AutoMigrate(&User{})

	// user := User{TID: "Jinzhu", Tag: pq.Int64Array{1234, 2312313, 864123, 1231}}
	// db.Create(&user)
	// user2 := User{TID: "zz", Tag: pq.Int64Array{1234, 12345, 1122, 111115}}

	// db.Create(&user2)
	// user3 := User{TID: "zz", Tag: pq.Int64Array{12345, 1122, 85423, 231142}}

	// db.Create(&user3)

	// var user User
	// db.Where(&User{Tag: pq.Int64Array{1234, 2312313, 864123, 1231}}).First(&user)

	// fmt.Println(user)

	// var users []User
	// db.Where(&User{Tag: pq.Int64Array{1234, 2312313, 864123, 1234}}).Or(&User{Tag: pq.Int64Array{2312313}}).Find(&users)

	// fmt.Println(users)

	// 含有以下一个标签
	var users2 []User
	db.Where("tag && '{12345, 2312313, 864123}'").Find(&users2)

	fmt.Println(users2)

	// 含有以所有标签
	var users3 []User
	db.Where("tag @> '{1234, 2312313, 864123}'").Find(&users3)
	// db.Find(&users2)

	fmt.Println(users3)
}
