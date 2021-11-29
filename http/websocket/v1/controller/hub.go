package controller

import (
	"math/rand"
	"net/http"
	"nowqs/frame/language"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Hub struct {
	clients    map[*Client]bool   // clients all client
	clientLock sync.RWMutex       // client read/write lock
	users      map[string]*Client // users all login user client,key: userType + "_" + device + "_" + userID
	userLock   sync.RWMutex       // userLock read/write lock
	broadcast  chan []byte        // broadcast send message to all client
	register   chan *Client       // register connection handle
	unregister chan *Client       // unregister connection close handle
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		users:      make(map[string]*Client),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.EventRegister(conn)
		case conn := <-h.unregister:
			h.EventUnregister(conn)
		}
	}
}

// createTag is careate client tag
func createTag() string {
	now := time.Now().UnixNano()
	rand.Seed(now)
	tag := strconv.FormatInt(now+(1000+rand.Int63n(4)), 10)
	return string(tag)
}

// WsHandle is handle websocket client
func WsHandle(h *Hub, c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  language.GetErrorMsg(1),
		})
		return
	}

	client := NewClient(h, conn.RemoteAddr().String(), conn)
	h.register <- client
	go client.Read()
	go client.Write()
}
