package main

import (

	"video-chat-app/server"
	"net/http"
	"log"

)

func main() {
  http.HandleFunc("/create", server.CreateRoomRequestHandler)
  http.HandleFunc("/join", server.JoinRoomRequestHandler )
  log.Println("Starting Server on Port 8080")
  err := http.ListenAndServe(":8080", nil)
  if(err != nil){
	  log.Fatal(err)
  }
}