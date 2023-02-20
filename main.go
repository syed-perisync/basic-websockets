package main

import (
	"log"
	"net/http"

	"test-web-socket/hub"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	hub := hub.New()
	hub.Run()

	h := New(hub)
	http.HandleFunc("/ws", h.HandleWS)

	if err := http.ListenAndServe(":3051", nil); err != nil {
		log.Println(err)
	}
}

type IHandler interface {
	HandleWS(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) HandleWS(w http.ResponseWriter, r *http.Request) {

	upgrade := websocket.Upgrader{}
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("unable to upgrade the connection,", err)
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Println("unable to generate a uuid,", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID := id.String()
	c := h.hub.Register(conn, userID)

	go c.Read(h.hub)
}

type Handler struct {
	hub hub.IHub
}

func New(hub hub.IHub) IHandler {
	return &Handler{hub: hub}
}
