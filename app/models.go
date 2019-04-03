package app

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `gorm:"not null" json:"username"`
	Password string	`gorm:"not null" json:"password"`
	Email string	`gorm:"not null" json:"email"`
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
