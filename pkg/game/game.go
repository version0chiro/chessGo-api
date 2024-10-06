// core game logic for the chess-go-server
package game

import (
	"fmt"
	"log"

	"github.com/version0chiro/chessGo-api/internal/models"
)

type GameSession struct {
	Player1 models.Player
	Player2 models.Player
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
}
