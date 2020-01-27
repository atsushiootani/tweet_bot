package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func dbOpenConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/test.sqlite3")
	if err != nil {
		panic("failed to connect db")
	}
	return db
}

func DbOpenConnection() *gorm.DB {
	return dbOpenConnection()
}

func DbInit() {
	db := dbOpenConnection()
	defer db.Close()

	// db.AutoMigrate(&Tweet{})
}

func DbCreate(record interface{}) {
	db := dbOpenConnection()
	defer db.Close()

	db.Create(record)
}

func DbSave(record interface{}) {
	db := dbOpenConnection()
	defer db.Close()

	db.Save(record)
}

func DbDelete(record interface{}) {
	db := dbOpenConnection()
	defer db.Close()

	db.Delete(record)
}

func DbFindAll(records []*gorm.Model) {
	db := dbOpenConnection()
	defer db.Close()

	db.Find(&records)
}

func DbFind(id int) (record *gorm.DB) {
	db := dbOpenConnection()
	defer db.Close()

	db.First(&record, id)
	return
}
