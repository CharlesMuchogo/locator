package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	wsConn, err := wsUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Printf("could not upgrade: %s \n", err.Error())
	}
	defer wsConn.Close()

	for {

		var location Location
		err := wsConn.ReadJSON(&location)

		if err != nil {
			fmt.Printf("error reading json: %s \n", err.Error())
		}

		// Echo the received message back to the client
		err = wsConn.WriteJSON(location)
		if err != nil {
			fmt.Printf("error writing json: %s \n", err.Error())
			break
		}

	}
}
