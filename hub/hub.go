package hub

import (
	"log"

	"github.com/gorilla/websocket"
)

type IHub interface {
	Run()
	Stop()

	Register(conn *websocket.Conn, userID string) IClient
}

func (h *Hub) Register(conn *websocket.Conn, userID string) IClient {

	c := NewClient(conn, userID)
	h.clients[userID] = c

	return c
}

func (h *Hub) Run() {
	log.Println(h.clients)
}

func (h *Hub) Stop() {
}

type Hub struct {
	clients map[string]*Client
	Test    bool
}

func New() IHub {
	return &Hub{
		clients: make(map[string]*Client),
	}
}
