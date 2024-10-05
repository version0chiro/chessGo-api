// api handlers for login singup and logout
package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	if u.Username == "Chess" && u.Password == "Chess" {
		tokenString, err := createToken(u.Username)
		fmt.Println("Token: ", tokenString)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error: ", err)

			fmt.Fprintf(w, "Failed to create token")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Token: %s", tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Invalid username or password")
	}
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
