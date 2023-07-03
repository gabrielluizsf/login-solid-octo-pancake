package v1

import (
	"net/http"
	"log"
 	"github.com/gabrielluizsf/login-solid-octo-pancake/controllers"
)

func Start(PORT string){
	http.Handle("/ip",http.HandlerFunc(controllers.WhoIP))
	http.Handle("/signup",http.HandlerFunc(controllers.NewLogin))
	http.Handle("/usernames",http.HandlerFunc(controllers.Find_Usernames))
	log.Fatal(http.ListenAndServe(PORT,nil))
}