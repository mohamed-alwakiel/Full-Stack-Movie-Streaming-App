package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-alwakiel/Full-Stack-Movie-Streaming-App/server/controller"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("failed : ", err)
	}
}
