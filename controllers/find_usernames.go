package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"log"
    "github.com/gabrielluizsf/login-solid-octo-pancake/models"
	"github.com/gabrielluizsf/login-solid-octo-pancake/database"
	"go.mongodb.org/mongo-driver/bson"
)

func Find_Usernames(w http.ResponseWriter, r *http.Request){
	connect, client := database.Connect();
	filter := bson.M{}
	
	cur, err := connect.Find(context.Background(),filter)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(context.Background())

	var users []models.UserPublic

	for cur.Next(context.Background()){
		var user models.UserPublic
		err := cur.Decode(&user)
		if err != nil{
			log.Printf(err)
		    return
		}
		users = append(users,user)	
	}
	err = client.Disconnect(context.Background())
	if err != nil{
		log.Printf(err)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(users)

}

