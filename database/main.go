package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	InitMigration()
	u := User{
		Name:  "hoge",
		Email: "hoge@jejeje.com",
	}
	h := newUser(u)
	fmt.Println(h)
}

type User struct {
	gorm.Model
	Name  string
	Email string
}

func OpenDatabase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./test.db")
	checkConnectError(err)
	return db
}

func checkConnectError(err error) {
	if err != nil {
		panic("DBの接続エラー!!")
	}
}

func InitMigration() {
	db := OpenDatabase()
	defer db.Close()
	db.AutoMigrate(&User{})
}

func newUser(u User) *User {
	db := OpenDatabase()
	defer db.Close()
	db.Create(&u)
	user := new(User)
	db.Where("name=?", u.Name).Find(&user)
	return user
}
