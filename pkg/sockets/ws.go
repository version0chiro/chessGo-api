package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/version0chiro/chessGo-api/internal/models"
	"github.com/version0chiro/chessGo-api/pkg/queue"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(qm *queue.QueueManager, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade to websocket:", err)
		return
	}

	playerID := r.URL.Query().Get("player_id")

	player := models.Player{
		ID:       playerID,
		Username: playerID,
		Conn:     conn,
	}

	qm.AddPlayer(player)
}

func handlePlayerConnection(conn *websocket.Conn, playerID string) {
	defer conn.Close()
	for {
		// Read message from WebSocket
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Handle incoming moves or events
		log.Printf("Received message from player %s: %s", playerID, message)

		// Example: Broadcast move to opponent
		broadcastToOpponent(playerID, message)
	}
}

func broadcastToOpponent(playerID string, message []byte) {
	panic("unimplemented")
}
