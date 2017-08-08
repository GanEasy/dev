package serve_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID         int
	Email      string
	GroupUsers []*GroupUser
}

type Group struct {
	ID         int
	Name       string
	GroupUsers []*GroupUser
}

type GroupUser struct {
	ID      int
	Group   *Group
	GroupID int
	User    *User
	UserID  int
	Owner   bool
}

func (*GroupUser) TableName() string {
	return "group_user"
}

// Test_IntMapToString []int to []string
func Test_URLAddParam(t *testing.T) {
	db, err := gorm.Open("sqlite3", "tmp.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Group{}, &GroupUser{})

	db.LogMode(true)

	db.Debug().Save(&User{
		Email: "1@1.com",
		GroupUsers: []*GroupUser{
			{
				Group: &Group{
					Name: "ccc",
				},
				Owner: true,
			},
		},
	})

	var group Group
	if err := db.Preload("GroupUsers.User").First(&group).Error; err != nil {
		panic(err)
	}

	spew.Dump(group)

}
