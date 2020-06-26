package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	mux "github.com/gorilla/mux"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Credentials used for receiving username and password
// from users
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims used for giving user it's claim if successfully signed in
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtkey = []byte("DwiKurniawan21")

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pegawai/getpegawai", returnAllPegawai).Methods("GET")
	router.HandleFunc("/pegawai/updatepegawai/{id}", updatePegawai).Methods("PUT")
	router.HandleFunc("/pegawai/newpegawai", createPegawai).Methods("POST")
	router.HandleFunc("/pegawai/deletepegawai/{id}", deletePegawai).Methods("DELETE")
	router.HandleFunc("/pegawai/signin", signinPegawai)

	fmt.Println("Connected")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signinPegawai(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPass, ok := users[creds.Username]

	if !ok || creds.Password != expectedPass {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
