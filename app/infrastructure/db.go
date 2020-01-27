package infrastructure

import "github.com/jinzhu/gorm"


func dbOpenConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/test.sqlite3")
	if err != nil {
		panic("failed to connect db")
	}
	return db
}

//DB初期化
func DbInit() {
	db := dbOpenConnection()
	defer db.Close()

	//db.AutoMigrate(&Todo{})
}
