package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var mySecretkey = []byte("My secrete key not secure @don't revel")

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Homepage)
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}
	log.Fatal(srv.ListenAndServe())

}
func Homepage(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authoried"] = true
	claims["id"] = "1"
	claims["user"] = "anil"
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	tokenString, err := token.SignedString(mySecretkey)
	if err != nil {
		fmt.Errorf("Something went wrong : %s", err.Error())
	}

	fmt.Fprint(w, tokenString)
}
