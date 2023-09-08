package main

import (
	"gin-onter/models"
	"gin-onter/router"
)

func main() {
	models.InitDb()
	router.InitRouter()

}
