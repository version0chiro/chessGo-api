// core game logic for the chess-go-server
package game

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/version0chiro/chessGo-api/internal/models"
)

type GameSession struct {
	Player1 models.Player
	Player2 models.Player
}

type Message struct {
	Type    string     `json:"type"` // e.g., "start_game", "move"
	Content string     `json:"content"`
	Board   [][]string `json:"board"`
	Turn    string     `json:"turn"`
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
	username1 := gs.Player1.Username
	username2 := gs.Player2.Username
	turn := username1

	if rand.Intn(2) == 0 {
		log.Println(username1, "goes first.")
		turn = username1
	} else {
		log.Println(username2, "goes first.")
		turn = username2
	}

	board := [][]string{
		{"r", "n", "b", "q", "k", "b", "n", "r"}, // Black back rank
		{"p", "p", "p", "p", "p", "p", "p", "p"}, // Black pawns
		{"", "", "", "", "", "", "", ""},         // Empty rank
		{"", "", "", "", "", "", "", ""},         // Empty rank
		{"", "", "", "", "", "", "", ""},         // Empty rank
		{"", "", "", "", "", "", "", ""},         // Empty rank
		{"P", "P", "P", "P", "P", "P", "P", "P"}, // White pawns
		{"R", "N", "B", "Q", "K", "B", "N", "R"}, // White back rank
	}
	startGameMessage := Message{
		Type:    "startGame",
		Content: "Game session started between: " + gs.Player1.Username + " and " + gs.Player2.Username,
		Board:   board,
		Turn:    turn,
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

		var move map[string]interface{}
		err = json.Unmarshal(message, &move)
		if err != nil {
			log.Println("Error unmarshalling move message:", err)
			return
		}
		var board [][]string
		boardData, err := json.Marshal(move["board"])
		if err != nil {
			log.Println("Error marshalling board data:", err)
			return
		}
		err = json.Unmarshal(boardData, &board)
		if err != nil {
			log.Println("Error unmarshalling board data:", err)
			return
		}
		from := move["from"].(map[string]interface{})
		to := move["to"].(map[string]interface{})
		fromRow := int(from["row"].(float64))
		fromCol := int(from["col"].(float64))
		toRow := int(to["row"].(float64))
		toCol := int(to["col"].(float64))
		fmt.Println("Move: ", fromRow, fromCol, toRow, toCol)
		isValid := IsValidMove(fromRow, fromCol, toRow, toCol, board)
		fmt.Println("Is valid move: ", isValid)

		if !isValid {
			log.Println("Invalid move by ", player.Username)
			invalidMoveMessage := Message{
				Type:    "invalidMove",
				Content: "Invalid move by " + player.Username,
				Board:   board,
				Turn:    opponent.ID,
			}
			invalidMoveJson, err := json.Marshal(invalidMoveMessage)
			if err != nil {
				log.Println("Error marshalling invalid move message:", err)
				return
			}
			err = opponent.Conn.WriteMessage(1, invalidMoveJson)
			if err != nil {
				log.Println("Error sending message:", err)
				opponent.Conn.Close()
				return
			}
			err = player.Conn.WriteMessage(1, invalidMoveJson)
			if err != nil {
				log.Println("Error sending message:", err)
				player.Conn.Close()
				return
			}

		} else {
			board[toRow][toCol] = board[fromRow][fromCol]
			board[fromRow][fromCol] = ""
			fmt.Println("Board after move: ", board)
			move["board"] = board
			move["type"] = "move"
			moveMessage := Message{
				Type:    "move",
				Content: "Move made by " + player.Username,
				Board:   board,
				Turn:    opponent.ID,
			}

			moveJson, err := json.Marshal(moveMessage)
			if err != nil {
				log.Println("Error marshalling move message:", err)
				return
			}

			err = opponent.Conn.WriteMessage(1, moveJson)

			if err != nil {
				log.Println("Error sending message:", err)
				opponent.Conn.Close()
				return
			}

			err = player.Conn.WriteMessage(1, moveJson)
			if err != nil {
				log.Println("Error sending message:", err)
				player.Conn.Close()
				return
			}
		}
	}
}
