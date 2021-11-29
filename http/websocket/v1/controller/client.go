package controller

import (
	"encoding/json"
	"fmt"
	"nowqs/frame/config"
	"nowqs/frame/language"
	"time"

	"github.com/gorilla/websocket"
)

const (
	expireTime = 50 // expiretTime is connection expiret time(second)
)

var deviceList = []string{"ios", "android", "web"}

// Client is client struct
type Client struct {
	hub           *Hub            // hub
	conn          *websocket.Conn // connection
	device        string          // device is client device ios/android/web
	userType      string          // userType is user type message, which exists after login
	userID        string          // userID is user id, which exists after login
	address       string          // address, is ip address
	firstTime     uint64          // firstTime is first connection time
	heartbeatTime uint64          // heartbeatTime is last heartbeat time
	loginTime     uint64          // loginTime is last login time, which exists after login
	send          chan []byte     // send message
}

// NewClient is add new client to hub
func NewClient(hub *Hub, address string, conn *websocket.Conn) (client *Client) {
	now := uint64(time.Now().Unix())
	client = &Client{
		hub:           hub,
		conn:          conn,
		address:       address,
		firstTime:     now,
		heartbeatTime: now,
		send:          make(chan []byte),
	}
	return
}

// Heartbeat is connection heartbeat
func (c *Client) Heartbeat() {
	c.heartbeatTime = uint64(time.Now().Unix())
}

// IsConnectionTimeOut is judgment time out
func (c *Client) IsConnectionTimeOut() (is bool) {
	now := uint64(time.Now().Unix())
	if c.heartbeatTime+expireTime <= now {
		is = true
	}
	c.Close()
	return
}

// Close is close c.send channel
func (c *Client) Close() {
	close(c.send)
}

// Read is read client message
func (c *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			config.WriteLog(fmt.Errorf("%v", r))
		}
	}()

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			config.WriteLog(err)
			c.hub.unregister <- c
			c.conn.Close()
			break
		}
		c.send <- message
	}
}

func (c *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			config.WriteLog(fmt.Errorf("%v", r))
		}
	}()
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, c.JSON(&Message{
					Action: "Unknown",
					Code:   1,
					Msg:    language.GetErrorMsg(1),
				}))
				return
			}
			response := c.HandleMessage(message)

			c.conn.WriteMessage(websocket.TextMessage, response)
		}
	}
}

func (c *Client) HandleMessage(message []byte) []byte {
	base := &_base{}
	err := json.Unmarshal(message, base)
	response := &Message{
		Action: "Unknown",
		Code:   1,
	}
	if err != nil {
		response.Msg = err.Error()
		return c.JSON(response)
	}
	switch base.Action {
	case "login":
		request := &_Login{}
		err = json.Unmarshal(message, request)
		// reslut := Login(request.UserType, request.Token)
		response.Code = 0
		response.Msg = language.GetMsg("success")
	default:
		response.Msg = language.GetErrorMsg(7)
	}

	return c.JSON(response)
}

func (c *Client) JSON(msg *Message) []byte {
	message, _ := json.Marshal(msg)
	return message
}

// TODO: Login
func Login(userType string, Token string) {
	return
}
