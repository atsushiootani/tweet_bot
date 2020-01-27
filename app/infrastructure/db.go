package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Tweet struct {
	gorm.Model
	Text   string // 文面
	Status string // 投稿したか
	TweetAt time.Time // 投稿する日時
}

func dbOpenConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/test.sqlite3")
	if err != nil {
		panic("failed to connect db")
	}
	return db
}

func DbInit() {
	db := dbOpenConnection()
	defer db.Close()

	db.AutoMigrate(&Tweet{})
}

func DbInsert(text string, status string, tweetAt time.Time) {
	db := dbOpenConnection()
	defer db.Close()

	db.Create(&Tweet{Text: text, Status: status, TweetAt: tweetAt})
}

func DbUpdate(id int, text string, status string, tweetAt time.Time) {
	db := dbOpenConnection()
	defer db.Close()

	var tweet Tweet
	db.First(&tweet, id)
	tweet.Text = text
	tweet.Status = status
	tweet.TweetAt = tweetAt
	db.Save(&tweet)
}

func DbDelete(id int) {
	db := dbOpenConnection()
	defer db.Close()

	var tweet Tweet
	db.First(&tweet, id)
	db.Delete(&tweet)
}

func DbGetAll() []Tweet {
	db := dbOpenConnection()
	defer db.Close()

	var tweets []Tweet
	db.Order("created_at desc").Find(&tweets)
	return tweets
}

func DbGetOne(id int) Tweet {
	db := dbOpenConnection()
	defer db.Close()

	var tweet Tweet
	db.First(&tweet, id)
	return tweet
}
