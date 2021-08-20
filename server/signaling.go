package server

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type broadcastMsg struct {
	Message map[string]interface{}
	RoomId string
	Client *websocket.Conn
}

var braodcast = make(chan broadcastMsg)
func broadcaster(){
	for {
		msg := <- braodcast


		for _, client := range AllRooms.Map[msg.RoomId]{
			if(client.Conn != msg.Client){
				err := msg.Client.WriteJSON(msg.Message)
			    if err != nil{
					log.Fatal("Unable to send message",err)
					client.Conn.Close()
				}
			}
		}
	}
}

// Global  HashMap is the gloobal hashmap for the server
var AllRooms RoomMap
// CreateRoomRequestHandler  Create room and return RoomID
func  CreateRoomRequestHandler(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	roomId := AllRooms.CreateRoom()
 type resp struct {
	 RoomId string `json:"room_id"`
 }
 log.Println(AllRooms.Map)
 json.NewEncoder(w).Encode(resp{RoomId:roomId})
}

var upgrager = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}
// JoinRoomRequestHandler will join the client in a particular room
func JoinRoomRequestHandler( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	roomId,ok := r.URL.Query()["roomId"]
    if !ok {
		log.Println("roomId is missing in Parameters")
	    return
	}
    ws,err := upgrager.Upgrade(w,r,nil)
    if err != nil{
		log.Fatal("Web socket upgrade error", err)
	}
	AllRooms.InsertIntoRoom(roomId[0],false,ws)
    go broadcaster()

	for{
		var msg broadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil{
			log.Fatal("Could not send mesage", err)
		}
		msg.Client = ws
		msg.RoomId = roomId[0]
		log.Println(msg.Message)
		braodcast <- msg
	}
}