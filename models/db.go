package models

import "mvc-gin/component/mysql"

func InitDb()  {
	db := mysql.GetMysqlDb()
	_ = db.AutoMigrate(&User{})
}
