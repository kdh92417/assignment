package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func getStage1(c * gin.Context) {
	value, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(value, err)

	i := 0
	for i < len(value) {
        value[i] += 1 
        i += 1
	}

	fmt.Println(value, string(value))
	c.Data(http.StatusOK, "text/html", value)
}

func getStage2(c * gin.Context) {
	value, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for i < len(value) {
		if value[i] != 10 {
			value[i] += 1
		}
		i += 1
	}
	c.Data(http.StatusOK, "text/html", value)
}

func getStage1Result(c * gin.Context) {
	value, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(value, string(value), err)
}

func main() {
	router := gin.Default()
	router.POST("/stage1", getStage1)
	router.POST("/stage2", getStage2)
	router.POST("/stage3", getStage2)
	router.POST("/stage/result", getStage1Result)
	router.Run("0.0.0.0:8090")
}

