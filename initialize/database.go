package initialize

import (
	"fmt"
	"foxtail/global"
	"foxtail/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	mysqlInfo := global.Settings.Mysqlinfo
	sqliteInfo := global.Settings.Sqliteinfo

	if sqliteInfo.Path != "" {
		fmt.Println("now sqlite")
		db, _ := gorm.Open(sqlite.Open(sqliteInfo.Path), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		global.DB = db
	} else {
		// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			mysqlInfo.Name, mysqlInfo.Password, mysqlInfo.Host,
			mysqlInfo.Port, mysqlInfo.DBName)
		db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		global.DB = db
	}

	global.DB.AutoMigrate(&model.User{},
		&model.Project{},
		&model.Company{},
		&model.UserProject{},
		&model.PersonalAccessToken{})

	//migrate(db)
}
