// queue for the chess-go-server
package queue

import (
	"fmt"
	"log"
	"sync"

	"github.com/version0chiro/chessGo-api/internal/models"
	"github.com/version0chiro/chessGo-api/pkg/game"
)

type QueueManager struct {
	queue []models.Player
	mutex sync.Mutex
}

func NewQueueManager() *QueueManager {
	return &QueueManager{
		queue: make([]models.Player, 0),
	}
}

func (qm *QueueManager) AddPlayer(p models.Player) {
	qm.mutex.Lock()
	defer qm.mutex.Unlock()

	qm.queue = append(qm.queue, p)
	log.Println("Player added to queue: ", p.Username)
	fmt.Println("Queue: ", qm.queue)

	if len(qm.queue) >= 2 {
		go qm.MatchPlayers()
	}
}

func (qm *QueueManager) MatchPlayers() {
	if len(qm.queue) < 2 {
		return
	}
	p1 := qm.queue[0]
	p2 := qm.queue[1]

	qm.queue = qm.queue[2:]

	log.Println("Player 1: ", p1.Username)
	log.Println("Player 2: ", p2.Username)

	game.StartGame(p1, p2)

}
