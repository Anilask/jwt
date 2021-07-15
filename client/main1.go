package main

import(
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
	"time"
	"io/ioutil"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySecretKey= []byte("My Secret key is not secret@dont reveal ffff")
func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/",HomePage)
	srv:=&http.Server{
		Handler: router,
		Addr: "localhost:8080",
		
	}
	log.Fatal(srv.ListenAndServe())
	
}
func HomePage(w http.ResponseWriter,r *http.Request){
	
	token:=jwt.New(jwt.SigningMethodHS256)

	claims:=token.Claims.(jwt.MapClaims)
	claims["authorized"]=true
	claims["id"]="1"
	claims["user"]="Chaitanya"
	claims["exp"]=time.Now().Add(time.Minute*10).Unix()

	tokenString,err:=token.SignedString(mySecretKey)
	if err!=nil{
		fmt.Errorf("Something went wrong : %s",err.Error())
	}

	client:=&http.Client{}
	req,_:=http.NewRequest("GET","http://localhost:9000/home",nil)
	req.Header.Set("Token",tokenString)

	res,err:=client.Do(req)
	if err!=nil{
		fmt.Fprintf(w,"Error message : %s",err.Error())
	}
	body,err:=ioutil.ReadAll(res.Body)
	if err!=nil{
		fmt.Println(w,err.Error())
	}
	fmt.Fprintf(w,string(body))

	
}