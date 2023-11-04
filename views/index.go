package views

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexView(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "If you are seeing this maybe we should hire you!"})
}
