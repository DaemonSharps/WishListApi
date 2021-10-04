package Services

import (
	"WishListApi/Types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var wishes = []Types.WishRecord{
	{ID: 1, Name: "Хочу ватружку"},
	{ID: 2, Name: "хочу пакушать"},
}

func GetWishes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, wishes)
}
