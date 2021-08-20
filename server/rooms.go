package server

import (
	//"debug/dwarf"
	
	"log"
	//"strings"
	"sync"
	"math/rand"
	"time"
	"github.com/gorilla/websocket"
)
// Participant describes a single Entity in the HashMap
type Participant struct {
	Host bool
	Conn *websocket.Conn
}
type  RoomMap struct {
	Mutex sync.RWMutex
	Map map[string][]Participant
}
// initialize riom struct
func (r *RoomMap) Init(){
r.Mutex.Lock()
defer r.Mutex.Unlock()
r.Map = make(map[string][]Participant)
}
// Return Participant
func (r *RoomMap) Get(roomId string)[]Participant{
r.Mutex.RLock()
defer r.Mutex.RUnlock()
return r.Map[roomId]
}
//CreateRoom generate a unique room Id and return it -> insert 
func (r *RoomMap) CreateRoom()string {
// generate room id and insert to hash map

r.Mutex.Lock()
defer r.Mutex.Unlock()
rand.Seed(time.Now().UnixNano())
var letters = []rune("abcdefghijklmnopqrstuvwxyz")
b := make([]rune, 8)

for i := range b {
	b[i] = letters[rand.Intn((len(letters)))]
}
roomId := string(b)
log.Println(r)
log.Println("We have acreated the roomId")
r.Map= make(map[string][]Participant)
r.Map[roomId] = []Participant{}
log.Println(r)
return roomId
}

func (r *RoomMap) DeleteRoom(roomId string){
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomId)
}
//InsertIntoRoom will create a participant and add it in the hashmap 
func (r *RoomMap) InsertIntoRoom(roomId string, host bool, conn *websocket.Conn){
r.Mutex.Lock()
defer r.Mutex.Unlock()

p := Participant{host,conn}
log.Println("Inserting into Room with RoomId", roomId)
r.Map[roomId] = append(r.Map[roomId], p)
}
