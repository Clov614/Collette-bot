package ws

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Clients struct {
	conn *websocket.Conn
}
type Hub struct {
	Chmsg        chan []byte
	Sendmsg      chan interface{}
	LoginSuccess chan bool
}

var (
	Client = &Clients{}
	// 数据处理中心
	Hhub = &Hub{
		Chmsg:        make(chan []byte, 20),
		Sendmsg:      make(chan interface{}, 20),
		LoginSuccess: make(chan bool),
	}
)

var up = &websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
}

func WsHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := up.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalln("获取链接失败：", err)
	}
	// 将conn 写入client
	Client = &Clients{
		conn: conn,
	}
	go Hhub.Read(Client)
}

func (hub *Hub) Read(client *Clients) {
	//从连接中循环读取信息
	for {
		_, msg, err := client.conn.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}
		hub.Chmsg <- msg
	}
}

//func (hub *Hub) Write(client *Clients, sendType interface{}) {
//	err := client.conn.WriteJSON(sendType)
//	if err != nil {
//		log.Fatalf("%v \n", fmt.Sprintf("WriteJson Error :%v", err))
//	}
//
//}

func GetClient() *websocket.Conn {
	return Client.conn
}

//func Write(conn *websocket.Conn, sendType interface{}) {
//	err := conn.WriteJSON(sendType)
//	if err != nil {
//		log.Fatalf("%v \n", fmt.Sprintf("WriteJson Error :%v", err))
//	}
//
//}
