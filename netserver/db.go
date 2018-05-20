package netserver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func OpenDatabase() {
	var err error
	DB, err = gorm.Open(Conf.Database.Driver, Conf.Database.Connection)

	if err != nil {
		panic(err)
	}

	fmt.Println("--- Open database ---")
	DB.LogMode(Conf.Debug)
	DB.SingularTable(true)
}
