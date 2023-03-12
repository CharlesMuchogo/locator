package location

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/databasehandler"
	"main.go/myStructs"
	"net/http"
)

func UpdateLocation(c *gin.Context) {
	var location myStructs.LocationUpdate

	if err := c.ShouldBindJSON(&location); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	status, response := databasehandler.UpdateLocation(location.UserId, location.CurrentLatitude, location.CurrentLongitude, location.MaxDistance, location.OriginLatitude, location.OriginLongitude)

	if status == 200 {
		c.JSON(http.StatusOK, gin.H{"message": response})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": response})
	}
}

func GetLocation(c *gin.Context) {

	userLocation, status := databasehandler.GetUsersLocation()

	if status == 200 {
		c.JSON(http.StatusOK, gin.H{
			"message": "get devices location success",
			"devices": userLocation,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "server error",
		})
	}
}
