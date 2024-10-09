// queue handlers for the chess-go-server
package queue

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/version0chiro/chessGo-api/internal/models"
)

var queueManager = NewQueueManager()

func QueueHandler(w http.ResponseWriter, r *http.Request, db *dynamodb.Client) {
	w.Header().Set("Content-Type", "application/json")
	var u models.Player
	json.NewDecoder(r.Body).Decode(&u)
	username := u.Username
	fmt.Println("Username: ", username)
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	player := models.Player{
		Username: username,
	}
	queueManager.AddPlayer(player)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Player added to queue"))
}
