package main

import (
	"./app/domain/tweet"
	"./app/infrastructure"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	// infrastructure.DbInit()
	db := infrastructure.DbOpenConnection()
	defer db.Close()
	db.AutoMigrate(&tweet.Tweet{})


	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		tweets := tweet.All()
		ctx.HTML(200, "index.html", gin.H{
			"tweets": tweets,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		// status := ctx.PostForm("status")
		// tweetAt := ctx.PostForm("tweetAt") // TODO
		tweet.Create(text, time.Now())
		ctx.Redirect(302, "/")
	})

	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		t := tweet.Get(id)
		ctx.HTML(200, "detail.html", gin.H{"tweet": t})
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		// tweetAt := ctx.PostForm("tweetAt") //TODO
		tweet := tweet.Get(id)
		tweet.Text = text
		tweet.Status = status
		tweet.TweetAt = time.Now()
		tweet.Save()
		ctx.Redirect(302, "/")
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

	router.Run()
}
