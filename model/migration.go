package model

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Leave{}, &Scan{}, &Article{}, &Dorm{})
	if err != nil {
		return
	}
	//DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}
