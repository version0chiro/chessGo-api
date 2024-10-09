// User, game and session structs
package models

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/websocket"
)

// App struct holds the DynamoDB client and other shared dependencies.
type App struct {
	DB *dynamodb.Client
}

// NewApp initializes the App struct with dependencies like DynamoDB client
func NewApp(ddb *dynamodb.Client) *App {
	return &App{
		DB: ddb,
	}
}

type Player struct {
	ID       string
	Username string
	Conn     *websocket.Conn
}
