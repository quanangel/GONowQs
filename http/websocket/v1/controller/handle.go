package controller

import (
	"fmt"
	"nowqs/frame/config"
)

// EventRegister is client connection event
func (h *Hub) EventRegister(client *Client) {
	h.clientLock.Lock()
	defer h.clientLock.Unlock()

	if config.AppConfig.Debug {
		fmt.Println("EventRegister")
	}

	h.clients[client] = true
}

// EventUnregister is client disconnect event
func (h *Hub) EventUnregister(client *Client) {
	h.clientLock.Lock()
	defer h.clientLock.Unlock()

	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
	}

	if config.AppConfig.Debug {
		fmt.Println("EventUnregister")
	}

	h.DelUser(client)
}

// AddUser is add user client, which exists after login
func (h *Hub) AddUser(client *Client) {
	h.userLock.Lock()
	defer h.userLock.Unlock()
	key := GetUserKey(client.userType, client.device, client.userID)
	h.users[key] = client
}

// DelUser is delete user client, which exists after login
func (h *Hub) DelUser(client *Client) (result bool) {
	h.userLock.Lock()
	defer h.userLock.Unlock()
	key := GetUserKey(client.userType, client.device, client.userID)
	if _, ok := h.users[key]; ok {
		delete(h.users, key)
		result = true
	}
	return
}

// GetUserKey is get user client key
func GetUserKey(userType string, device string, userID string) (key string) {
	key = userType + "_" + device + "_" + userID
	return
}

// GetUserAllKey is get user all client key
func GetUserAllKey(userType string, userID string) (list []string) {
	for _, val := range deviceList {
		list = append(list, GetUserKey(userType, val, userID))
	}
	return
}
