package main

import (
	"github.com/gin-gonic/gin"

	"WishListApi/Services"
)

func main() {
	router := gin.Default()

	router.GET("/test", Services.GetWishes)

	router.Run("localhost:3003")
}
