package app

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var(
	Configuration Settings
	CookieHandler *securecookie.SecureCookie
)

func (mycms *App) Initialise() {
	ConfigureInstance()

	Configuration.DBConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		Configuration.DBUsername, Configuration.DBPassword, Configuration.DBName)

	mycms.LoadRoutes()

	CookieHandler = securecookie.New(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))
}

func (mycms *App) migrate() {
	mycms.DB.CreateTable(&Post{}, &User{})
}

func (mycms *App) Run(Args []string) {
	var err error
	mycms.DB, err = gorm.Open("mysql", Configuration.DBConn)
	if err != nil {
		log.Fatalf("Could not establish database connection, shutting down\n\t" + err.Error())
		return
	}

	log.Println("Database connection established")
	defer mycms.DB.Close()

	for _, arg := range Args {
		switch arg {
		case "-m":
			fmt.Println("Migrating tables")
			mycms.migrate()
		}
	}

	err = http.ListenAndServe(Configuration.ServerAddress, mycms.router)
	if err != nil {
		log.Fatalf("Could not start server, shutting down")
	}
}