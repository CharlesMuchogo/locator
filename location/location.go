package location

import (
	"fmt"
	"net/http"

	"github.com/NaySoftware/go-fcm"
	"github.com/gin-gonic/gin"
	"main.go/databasehandler"
	"main.go/myStructs"
)

const (
	serverKey = "AAAAUxlnNFg:APA91bHqvt-A-ujXVg_N40DLacfkn6heZoBDI3o_1bPp1hxDX30TXSb5zsTrySc87FdmuTtx2gv-ajf3o2fRs4t7cgGlJDtRX1ampl-c5YwlYlO-tik9wFiRHKD7J4UiNCnpz0Nid7vW"
	topic     = "fJpRUcuBT72VAUepVdPNio:APA91bHETvO-aGmJo7xKDMLvRyHNHdlqemqN9wDicD9KqpkxmYdUBL5KgPX50Ki0qhiskTlpdMzcP4BFMjgCD82s7eceupe0wk1ATw1RNTTVhFBJKPA2Ro3XWHAlaBkzPMX2TwvQYmS7vz"
)

func UpdateLocation(c *gin.Context) {
	var location myStructs.LocationUpdate

	fcm := fcm.NewFcmClient(serverKey)

	fcm.NewFcmMsgTo(topic, "new message")

	if err := c.ShouldBindJSON(&location); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	status, response := databasehandler.UpdateLocation(location.UserId, location.CurrentLatitude, location.CurrentLongitude, location.MaxDistance, location.OriginLatitude, location.OriginLongitude)

	devices, devicesStatus := databasehandler.GetUsersLocation()

	if devicesStatus == 200 {
		if status == 200 {
			c.JSON(http.StatusOK, gin.H{"message": response, "users": devices})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": response})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "server error"})
	}
}

func GetLocation(c *gin.Context) {

	userLocation, status := databasehandler.GetUsersLocation()

	fcm := fcm.NewFcmClient(serverKey)
	fmt.Printf(" should be sending the message now \n")

	fcm.NewFcmMsgTo(topic, "new message")

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
