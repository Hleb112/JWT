package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = User{
	Username: "1",
	Password: "1",
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", handlePage)
	r.HandleFunc("/login", login).Methods("POST")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(request.Body).Decode(&message)
	if err != nil {
		return
	}
	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Println("user =", u)
	checkLogin(u)
	//io.WriteString(w, checkLogin(u))
}

func checkLogin(u User) string {
	if user.Username != u.Username || user.Password != u.Password {
		return "incorrect login/password"
	}
	validToken, err := generateJWT()

	if err != nil {
		fmt.Println("validGen=", err)
	}
	return validToken
}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["expiration"] = time.Now().Add(1 * time.Hour)

	tokenString, err := token.SignedString([]byte("dfdfjfd"))

	if err != nil {
		fmt.Println("tokenString=", err)
	}
	return tokenString, nil
}
