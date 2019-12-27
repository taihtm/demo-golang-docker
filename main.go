package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	hosts      = "localhost:27017"
	database   = "sendo-test"
	username   = ""
	password   = ""
	collection = "test_1"
)

func mongoConnection() (*mgo.Session, error) {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  3 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return session, err
	}
	return session, nil
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		_, err := mongoConnection()
		if err != nil {
			c.JSON(400, gin.H{
				"status": 0,
				"message": "Can't connect db",
			})
			return
		}
		c.JSON(200, gin.H{
			"status": 1,
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}