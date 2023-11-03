package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main.go/location"
	"main.go/user"
	"main.go/websocket"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register", user.Signup)
	public.POST("/login", user.Login)
	public.POST("/location", location.UpdateLocation)
	public.GET("/location", location.GetLocation)
	public.POST("/user", user.UpdateProfile)
	public.POST("/request_promotion", user.RequestPromotion)
	public.GET("/promote_user", user.PromoteUser)

	go func() {
		if err := r.Run(":8001"); err != nil {
			log.Fatal("HTTP server failed to start: ", err)
		}
	}()

	fmt.Println("HTTP server started on port 8001")
	router := mux.NewRouter()
	router.HandleFunc("/location", websocket.WsEndpoint)

	if err := http.ListenAndServe(":9200", router); err != nil {
		log.Fatal("WebSocket server failed to start: ", err)
	}
	fmt.Println("WS server started on port 9200")

}
