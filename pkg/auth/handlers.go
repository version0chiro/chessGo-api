// api handlers for login singup and logout
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/version0chiro/chessGo-api/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request, ddb *dynamodb.Client) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	_, password, err := db.GetUser(ddb, u.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		fmt.Fprintf(w, "Failed to get user")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid username or password")
		return
	}
	tokenString, err := createToken(u.Username)
	fmt.Println("Token: ", tokenString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		fmt.Fprintf(w, "Failed to create token")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Token: %s", tokenString)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
		return
	}
	fmt.Println("Token: ", tokenString)
	tokenString = tokenString[len("Bearer "):]
	err := verifyToken(tokenString)
	if err != nil {
		fmt.Println("Error: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorized")
	} else {
		fmt.Fprintf(w, "Protected endpoint")
	}
}

func SignupHandler(w http.ResponseWriter, r *http.Request, ddb *dynamodb.Client) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	hashedPassword, err := HashPassword(u.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		fmt.Fprintf(w, "Failed to hash password")
		return
	}
	if u.Username == "" || u.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Username and password are required")
		return
	}
	tokenString, err := createToken(u.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		fmt.Fprintf(w, "Failed to create token")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Token: %s", tokenString)
	db.AddUser(ddb, u.Username, hashedPassword)
}
