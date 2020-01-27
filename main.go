package main

import (
	"./app/infrastructure"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func main() {
	infrastructure.DbInit()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	//Index
	router.GET("/", func(ctx *gin.Context) {
		tweets := infrastructure.DbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"tweets": tweets,
		})
	})

	//Create
	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		// tweetAt := ctx.PostForm("tweetAt") // TODO
		infrastructure.DbInsert(text, status, time.Now())
		ctx.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		tweet := infrastructure.DbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"tweet": tweet})
	})

	//Update
	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		// tweetAt := ctx.PostForm("tweetAt") //TODO
		infrastructure.DbUpdate(id, text, status, time.Now())
		ctx.Redirect(302, "/")
	})

	//削除確認
	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		tweet := infrastructure.DbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"tweet": tweet})
	})

	//Delete
	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		infrastructure.DbDelete(id)
		ctx.Redirect(302, "/")

	})

	router.Run()

}
