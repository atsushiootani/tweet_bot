package main

import (
	"./app/infrastructure"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	_ "strconv"
)

func main() {
	infrastructure.DbInit()

	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.String(200, "Hello, world!")
	})
	r.Run()

}
