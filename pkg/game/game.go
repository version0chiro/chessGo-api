// core game logic for the chess-go-server
package game

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/version0chiro/chessGo-api/internal/models"
)

type GameSession struct {
	Player1 models.Player
	Player2 models.Player
}

type Message struct {
	Type    string `json:"type"` // e.g., "start_game", "move"
	Content string `json:"content"`
}

func StartGame(p1, p2 models.Player) {
	fmt.Println("Starting game between: ", p1.Username, "and", p2.Username)
	gameSession := GameSession{
		Player1: p1,
		Player2: p2,
	}
	go gameSession.Start()
}

func (gs *GameSession) Start() {
	log.Println("Game session started between: ", gs.Player1.Username, "and", gs.Player2.Username)
	startGameMessage := Message{
		Type:    "startGame",
		Content: "Game session started between: " + gs.Player1.Username + " and " + gs.Player2.Username,
	}
	message, err := json.Marshal(startGameMessage)
	if err != nil {
		log.Println("Error marshalling start game message:", err)
		return
	}
	gs.Player1.Conn.WriteMessage(1, message)
	gs.Player2.Conn.WriteMessage(1, message)
	go handlePlayerMessages(gs.Player1, gs.Player2)
	go handlePlayerMessages(gs.Player2, gs.Player1)
}

func handlePlayerMessages(player models.Player, opponent models.Player) {
	for {
		_, message, err := player.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			player.Conn.Close()
			return
		}

		log.Printf("received message for %s: %s\n", player.ID, message)
		moveMessage := Message{
			Type:    "move",
			Content: string(message),
		}
		json_message, m_err := json.Marshal(moveMessage)
		if m_err != nil {
			log.Println("Error marshalling move message:", m_err)
			return
		}
		err = opponent.Conn.WriteMessage(1, json_message)
		if err != nil {
			log.Println("Error sending message:", err)
			opponent.Conn.Close()
			return
		}
	}
}
