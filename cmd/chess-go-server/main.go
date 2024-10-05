// entry point of the chess-go-server

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/version0chiro/chessGo-api/pkg/auth"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/login", http.HandlerFunc(auth.LoginHandler)).Methods("POST")
	r.Handle("/protected", http.HandlerFunc(auth.ProtectedHandler)).Methods("GET")
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
