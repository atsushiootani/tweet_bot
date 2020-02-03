package tweet

import (
	"../../infrastructure"
	"fmt"
	"github.com/Songmu/go-httpdate"
	"github.com/jinzhu/gorm"
	"time"
	"unicode/utf8"
)

type Tweet struct {
	gorm.Model // TODO: ドメイン層としては外したい
	Text   string // 文面
	Status Status // 投稿したか
	TweetAt time.Time // 投稿する日時
}

type Status string
const(
	New Status = "New"
	Done Status = "Done"
	Failed Status = "Failed"
)

var Statuses = map[string]Status{
	"New": New,
	"Done": Done,
	"Failed": Failed,
}

var OrderedStatuses = []Status{
	New, Done, Failed,
}

func StringToStatus(str string) Status{
	return Statuses[str]
}

func StatusToString(status Status) string{
	return string(status)
}

func Create(text string, tweetAt time.Time) *Tweet{
	result := Tweet{
		Text: text,
		Status: New,
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

	db.Order("tweet_at asc").Find(&results)
	return
}

func Get(id int) *Tweet {
	db := infrastructure.DbOpenConnection()
	defer db.Close()

	var result Tweet
	db.First(&result, id)
	return &result
}

// 今ツイートすべきもの一覧
func GetNeedsToTweets(now time.Time) (results []*Tweet) {
	db := infrastructure.DbOpenConnection()
	defer db.Close()

	db.Where("status = ? AND tweet_at < ?", New, now).Find(&results)
	return
}

// 今ツイートすべきものをツイート
func TweetNowAll(now time.Time) {
	tweets := GetNeedsToTweets(now)

	for _, tweet := range tweets {
		tweet.DoTweet()
	}
}

const dateFormat = "2006/01/02 15:04:05"
var jst *time.Location

func getJst() *time.Location {
	if jst == nil {
		localJst, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			panic("failed to load location 'Asia/Tokyo'")
		}
		jst = localJst
	}
	return jst
}

func DateFormat(dateString string) (time.Time, error) {
	time, err := httpdate.Str2Time(dateString, getJst())
	return time, err
}

func (tweet *Tweet) TextLength() int{
	return utf8.RuneCountInString(tweet.Text)
}

func (tweet *Tweet) SetDate(dateString string) error{
	t, err := DateFormat(dateString)
	if err != nil {
		tweet.TweetAt = t
	}
	return err
}

func (tweet *Tweet) TweetAtString() string{
	return tweet.TweetAt.Format(dateFormat)
}

func (tweet *Tweet) Save() bool{
	infrastructure.DbSave(tweet)
	return true
}

// ツイート実行
func (tweet *Tweet) DoTweet() bool {

	// TODO: ツイート実行
	fmt.Println(tweet.Text)

	tweet.Status = Done
	tweet.Save()

	return true
}

func (tweet *Tweet) HasTweeted() bool {
	return tweet.Status != Done
}
