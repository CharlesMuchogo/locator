package main

import (
	"fmt"
	"main.go/location"
	"main.go/user"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	fmt.Printf("time now is %s \n", time.Now().Format("2006-01-02T15:04:05.999999Z"))
	fmt.Printf("time now is %s \n", time.Now())
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", user.Register)
	public.POST("/login", user.Login)
	public.POST("/location", location.UpdateLocation)
	public.GET("/location", location.GetLocation)

	r.Run(":8001")
	fmt.Println("Server started on port 8080")
}
