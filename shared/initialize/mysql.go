package initialize

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitMysql(userName, password, dbName string, tables... interface{}) *gorm.DB {
	dsn := "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	conn := mysql.Open(fmt.Sprintf(dsn, userName, password, dbName)) 
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	for _, t := range tables {
		err = db.AutoMigrate(t)
		if err != nil {
			panic(err)
		}
	}
	return db
}
