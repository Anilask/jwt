package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/home", isAuthorized(HomePage))
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:9000",
	}
	log.Fatal(srv.ListenAndServe())

}
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			endpoint(w, r)
		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authorized")
}
