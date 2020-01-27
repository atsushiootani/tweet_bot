package tweet

import (
	"../../infrastructure"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"unicode/utf8"
)

type Tweet struct {
	gorm.Model // TODO: ドメイン層としては外したい
	Text   string // 文面
	Status string // 投稿したか
	TweetAt time.Time // 投稿する日時
}

func Create(text string, tweetAt time.Time) *Tweet{
	result := Tweet{
		Text: text,
		Status: "New",
		TweetAt: tweetAt,
	}

	infrastructure.DbCreate(&result)

	return &result
}

func Delete(id int) bool {
	infrastructure.DbDelete(Get(id))
	return true
}

func All() (results []*Tweet) {
	db := infrastructure.DbOpenConnection()
	defer db.Close()

	db.Find(&results).Order("id asc")
	return
}

func Get(id int) *Tweet {
	db := infrastructure.DbOpenConnection()
	defer db.Close()

	var result Tweet
	db.First(&result, id)
	return &result
}

func (tweet *Tweet) TextLength() int{
	return utf8.RuneCountInString(tweet.Text)
}

func (tweet *Tweet) Save() bool{
	infrastructure.DbSave(tweet)
	return true
}

// ツイート実行
func (tweet *Tweet) DoTweet() bool {
	fmt.Println(tweet.Text)
	tweet.Status = "Done"
	return true
}

func (tweet *Tweet) HasTweeted() bool {
	return tweet.Status != "Done"
}
