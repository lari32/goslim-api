package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	gin.SetMode(viper.GetString("runmode"))

	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Use(gin.Recovery())

	r.NoRoute(func (c *gin.Context) {})

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	r.Run()
}

func pingServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get(fmt.Sprintf("%s:%d:/ping", viper.GetString("url"), viper.GetInt("port")))
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		log.Println("Waiting for the router, retry in 1 second")
		time.Sleep(1 * time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
