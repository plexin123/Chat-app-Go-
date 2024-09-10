package ws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"gopractice2/server/internal/message"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//create a new wesocket server so then it can be used in different parts of the application

//first create the construct

type WebsocketServer struct {
	Upgrader    websocket.Upgrader
	Clients     map[string]*Client
	MessageRepo message.Repository
}
type Client struct {
	Conn     *websocket.Conn
	Message  chan *message.Message
	ID       string
	Username string
}

// type Message struct {
// 	Content  string `json:"content"`
// 	Sender   string `json:"sender"`
// 	Receiver string `json:"receiver"`
// }

func NewWebsocketServer(repo message.Repository) *WebsocketServer {
	return &WebsocketServer{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		Clients:     make(map[string]*Client),
		MessageRepo: repo,
	}
}

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// var clients = make(map[string]*Client)

func (ws *WebsocketServer) HandleWebSocket(c *gin.Context) {
	conn, err := ws.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error while updating connection", err)
		return
	}

	client := &Client{
		Conn:    conn,
		Message: make(chan *message.Message),
		ID:      c.Request.URL.Query().Get("id"),
	}
	ws.Clients[client.ID] = client

	go client.readMessage(ws)
	go client.sendMessage()
}

func (ws *WebsocketServer) GetMessageHistory(c *gin.Context) {
	sender := c.Param("sender")
	receiver := c.Param("receiver")

	messages, err := ws.MessageRepo.GetMessageHistory(c.Request.Context(), sender, receiver)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retreive message"})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (c *Client) sendMessage() {
	defer func() {

	}()
	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) readMessage(ws *WebsocketServer) {
	defer func() {

		c.Conn.Close()
	}()

	for {

		_, m, err := c.Conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err) // Log any unexpected errors.
			}
			break // Exit the loop if there's an error.
		}
		//every time its going to put the json in a Message Object, its going to call the funcion SaveMessage
		var msg message.Message
		err = json.Unmarshal(m, &msg)

		log.Printf("Client Username: %s", c.Username)
		log.Printf("Message Received: %+v", msg)
		if err != nil {
			log.Printf("Error unmarshalling, %v", err)
			continue
		}
		ctx := context.Background()
		if err := ws.MessageRepo.SaveMessage(ctx, &msg); err != nil {
			log.Printf("Error saving message %v", err)
			continue
		}

		msg.Sender = c.Username
		if receiver, ok := ws.Clients[msg.Receiver]; ok {
			receiver.Message <- &msg
			c.Message <- &message.Message{
				Content:  "Message delivered",
				Sender:   "Server",
				Receiver: msg.Receiver,
			}
		} else {
			c.Message <- &message.Message{
				Content:  "Receiver not connected",
				Sender:   "Server",
				Receiver: msg.Receiver,
			}
		}
	}

}
