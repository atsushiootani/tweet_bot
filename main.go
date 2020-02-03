package main

import (
	"./app/domain/tweet"
	"./app/infrastructure"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	// infrastructure.DbInit()
	db := infrastructure.DbOpenConnection()
	defer db.Close()
	db.AutoMigrate(&tweet.Tweet{})

	// worker
	go worker()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		tweets := tweet.All()
		needsToTweets := tweet.GetNeedsToTweets(time.Now())
		ctx.HTML(200, "index.html", gin.H{
			"tweets": tweets,
			"needsToTweets": needsToTweets,
			"needsToTweetsCount": len(needsToTweets),
			"workerAvailable": workerAvailable,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")

		tweetAtString := ctx.PostForm("tweetAt")
		tweetAtTime, err := tweet.DateFormat(tweetAtString)

		if err == nil {
			tweet.Create(text, tweetAtTime)
			ctx.Redirect(302, "/")
		} else{
			ctx.Status(400)
		}
	})

	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		t := tweet.Get(id)
		ctx.HTML(200, "detail.html", gin.H{"tweet": t, "Statuses": tweet.OrderedStatuses})
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		tweetAtString := ctx.PostForm("tweetAt")
		tweetAtDate, err := tweet.DateFormat(tweetAtString)

		if err == nil {
			tw := tweet.Get(id)
			tw.Text = text
			tw.Status = tweet.StringToStatus(status)
			tw.TweetAt = tweetAtDate
			tw.Save()
			ctx.Redirect(302, "/")
		} else{
			ctx.Redirect(400, "/")
		}
	})

	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		tw := tweet.Get(id)
		ctx.HTML(200, "delete.html", gin.H{"tweet": tw})
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		tweet.Delete(id)
		ctx.Redirect(302, "/")
	})

	router.POST("/do_tweet", func(ctx *gin.Context) {
		tweet.TweetNowAll(time.Now())
		ctx.Redirect(302, "/")
	})

	router.POST("/worker_pause", func(ctx *gin.Context) {
		workerAvailable = false
		ctx.Redirect(302, "/")
	})

	router.POST("/worker_restart", func(ctx *gin.Context) {
		workerAvailable = true
		ctx.Redirect(302, "/")
	})

	router.Run()
}

// 定時ワーカー
const workerInterval = 10 * time.Second
var workerAvailable = true

func worker() {
	for {
		if workerAvailable {
			fmt.Println("worker process at ", time.Now().Format(time.RFC3339))

			tweet.TweetNowAll(time.Now())
		}

		time.Sleep(workerInterval)
	}
}
