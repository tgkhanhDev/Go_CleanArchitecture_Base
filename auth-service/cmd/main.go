package main

import (
	"fmt"
	"gin/internal/test"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	var a string = "123"
	router := gin.Default()
	fmt.Println("a ne: ", a)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	firstRs, err := test.NewRabbitMQTest("1")
	if err != nil {
		log.Fatal("firstRs error: ", err)
	}
	if firstRs != nil {
		log.Println("firstRs:", *firstRs)
	} else {
		log.Println("firstRs is nil")
	}

	secondRs, err := test.NewRabbitMQTest("2")
	if err != nil {
		log.Fatal("secondRs error: ", err)
	}
	if secondRs != nil {
		log.Println("secondRs:", *secondRs)
	} else {
		log.Println("secondRs is nil")
	}

	router.Run() // listen and serve on 0.0.0.0:8080
}
