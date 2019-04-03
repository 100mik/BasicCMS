package app

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var Configuration Settings

func (mycms *App) Initialise() {
	ConfigureInstance()

	Configuration.DBConn = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		Configuration.DBUsername, Configuration.DBPassword, Configuration.DBName)

	fmt.Println(Configuration.DBConn)

	mycms.LoadRoutes()
}

func (mycms *App) migrate() {
	mycms.DB.CreateTable(&User{})
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