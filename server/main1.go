package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)
var mySecretkey = []byte("My secrete key not secure @don't revel")

func main() {
	router := mux.NewRouter()
	router.Handle("/home", isAuthorized(HomePage))
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:9000",
	}
	log.Fatal(srv.ListenAndServe())

}
func isAuthorized(endpoint func(http.ResponseWriter,*http.Request))http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){

		if r.Header["Token"]!=nil{
			tokenString:=r.Header["Token"][0]

			token,err:=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
				
				if _,ok:=token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("we got some error: %v",token.Header["alg"])
				}
				return mySecretkey,nil
			})
			if err!=nil{
				fmt.Fprintf(w,err.Error())
			}
			if token.Valid{
				endpoint(w,r)
			}

			
			
			
		}else{
			fmt.Fprintf(w,"Not Authorized")
		}
	})
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authorized")
}
