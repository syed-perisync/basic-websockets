package hub

import (
	"log"

	"github.com/gorilla/websocket"
)

type IClient interface {
	Read(hub IHub)
	Write(data interface{}) error
}

func (c *Client) Read(hub IHub) {

	for {
		mt, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("->", err)
		}

		if mt == -1 {
			log.Println("client got disconnected..")
			break
		}

		log.Println(mt, string(msg))
	}
}

func (c *Client) Write(data interface{}) error {
	return c.conn.WriteJSON(data)
}

type Client struct {
	conn   *websocket.Conn
	userID string
}

func NewClient(conn *websocket.Conn, userID string) *Client {
	return &Client{
		conn:   conn,
		userID: userID,
	}
}
