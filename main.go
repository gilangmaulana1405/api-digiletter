package main

import (
	"github.com/api-digiletter/go-restapi-gin/controllers/mahasiswacontroller"
	"github.com/api-digiletter/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/mahasiswa", mahasiswacontroller.Index)

	r.Run()
}