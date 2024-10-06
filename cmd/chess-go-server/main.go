// entry point of the chess-go-server

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gorilla/mux"
	"github.com/version0chiro/chessGo-api/pkg/auth"
	"github.com/version0chiro/chessGo-api/pkg/queue"
)

type App struct {
	DB *dynamodb.Client
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("chess-go-local"),
		config.WithRegion("ap-south-1"))
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}
	ddb := dynamodb.NewFromConfig(cfg)
	app := App{DB: ddb}
	r := mux.NewRouter()
	r.Handle("/login", http.HandlerFunc(app.LoginHandler)).Methods("POST")
	r.Handle("/protected", http.HandlerFunc(auth.ProtectedHandler)).Methods("GET")
	r.Handle("/signup", http.HandlerFunc(app.SignupHandler)).Methods("POST")
	r.Handle("/queue", http.HandlerFunc(app.QueueHandler)).Methods("POST")
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func (app *App) SignupHandler(w http.ResponseWriter, r *http.Request) {
	auth.SignupHandler(w, r, app.DB)
}

func (app *App) LoginHandler(w http.ResponseWriter, r *http.Request) {
	auth.LoginHandler(w, r, app.DB)
}

func (app *App) QueueHandler(w http.ResponseWriter, r *http.Request) {
	queue.QueueHandler(w, r, app.DB)
}
