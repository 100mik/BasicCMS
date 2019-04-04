package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID int `gorm:"primary_key;auto_increment" json:"userID"`
	Username string `gorm:"not null" json:"username"`
	Password string	`gorm:"not null;unique" json:"password"`
	Email string	`gorm:"not null" json:"email"`
}

type Post struct {
	ID int `gorm:"primary_key;auto_increment" json:"postID"`
	User int `gorm:"not null" json:"userID"`
	Title string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
}

type CurrUser struct {
	ID int `json:"userID"`
	Username string `json:"username"`
}

type Settings struct {
	ServerAddress string
	DBUsername string
	DBPassword string
	DBName string
	DBConn string
}

type App struct {
	DB       *gorm.DB
	router   *mux.Router
}
