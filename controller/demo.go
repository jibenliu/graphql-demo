package controller

import (
	"fmt"
	"graphql-demo/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var now = time.Now()

func Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
		"data":       model.NewsList[id],
		"code":       200,
		"serverTime": now,
	})
}

func List(c *gin.Context) {
	c.JSON(200, gin.H{
		"data":       model.NewsList,
		"code":       200,
		"serverTime": now,
	})
}

func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "I am Post",
	})
}

func Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "I am Put",
	})
}

func Destroy(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "I am Delete",
	})
}
