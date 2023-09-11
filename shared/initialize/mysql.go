package initialize

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func InitMysql(tables... interface{}) *gorm.DB {
	userName := Config.GetString("db_username")
	password := Config.GetString("db_password")
	dbName := Config.GetString("db_name")
	addr := Config.GetString("db_addr")
	dsnfmt := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnfmt, userName, password, addr, dbName)
	klog.Info("connect to ", dsn)
	conn := mysql.Open(dsn) 
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
