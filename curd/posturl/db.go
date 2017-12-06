package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

//DB 返回 *gorm.DB
func DB() *gorm.DB {
	if db == nil {

		newDb, err := newDB()
		if err != nil {
			panic(err)
		}
		newDb.DB().SetMaxIdleConns(10)
		newDb.DB().SetMaxOpenConns(100)

		newDb.LogMode(false)
		db = newDb
	}

	return db
}

func newDB() (*gorm.DB, error) {

	// sqlConnection := fmt.Sprintf("host=%v user=%v port=%v dbname=%v sslmode=disable password=%v", conf.Conf.DB.Host, conf.Conf.DB.User, conf.Conf.DB.Port, conf.Conf.DB.DBName, conf.Conf.DB.Password)
	// db, err := gorm.Open(conf.Conf.DB.Type, sqlConnection)
	db, err := gorm.Open("sqlite3", "notice.db")

	// db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=wxrank sslmode=disable password=123456")
	// db, err := gorm.Open("postgres", "host=119.23.237.215 user=xiaoyi dbname=wxrank sslmode=disable password=123456")

	if err != nil {
		return nil, err
	}
	return db, nil
}

// Article has and belongs to many languages, use `Article_languages` as join table
type Article struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Author    string
	AppName   string
	AppID     string
	Cover     string
	Intro     string
	PubAt     time.Time
	Like      int64   `gorm:"type:default(0)"`
	Hate      int64   `gorm:"type:default(0)"`
	URL       string  `gorm:"type:varchar(100);unique_index"`
	Rank      float64 `sql:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func main() {
	var articles []Article
	DB().Order("id DESC").Find(&articles)
	for k, article := range articles {
		fmt.Println(k, article.URL)
		pushLink(article.URL)
	}
}

func pushLink(u string) {

	link := url.QueryEscape(u)
	doc, err := goquery.NewDocument(fmt.Sprintf("http://localhost:8005/fetch?url=%v", link))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Text())
	fmt.Println(u)
}
