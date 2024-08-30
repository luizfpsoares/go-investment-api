package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizfpsoares/albums/model"
	"github.com/luizfpsoares/albums/storage"
)

func main() {

	router := gin.Default()
	storage.ConnectDatabase()
	router.GET("/healtz", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status": "running",
		})
	})
	router.GET("/user", model.GetUser)
	router.POST("/user", model.AddUser)

	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
