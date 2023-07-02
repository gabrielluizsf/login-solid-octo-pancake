package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gabrielluizsf/login-solid-octo-pancake/models"
	"github.com/gabrielluizsf/login-solid-octo-pancake/database"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func NewLogin(w http.ResponseWriter, r *http.Request){
	const (
		successMessage = "Login cadastrado com sucesso"
		failMessage    = "Falha ao cadastrar Login"
	)
	var (
		existingAccount models.Login
		loginSignup models.Login
	)
	err := json.NewDecoder(r.Body).Decode(&loginSignup)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	passwordEncrypted, err := bcrypt.GenerateFromPassword([]byte(loginSignup.Password), 10)
	if err != nil{
		http.Error(w, "PASSWORD ENCRYPT ERROR", http.StatusInternalServerError)
		return
	}
	signupDocument := bson.M{
		"username":  loginSignup.Username,
		"passoword": string(passwordEncrypted),
	}
	connection, client := database.Connect();
	defer client.Disconnect(context.Background());

	// Verifica se já existe alguma conta com o mesmo username
	existingAccountDocument := bson.M{"username": loginSignup.Username}
	err = connection.FindOne(context.Background(), existingAccountDocument).Decode(&existingAccount)
	if err == nil {
		http.Error(w, "REQUEST ERROR", http.StatusBadRequest)
		return
	}
	result, err := connection.InsertOne(context.Background(), signupDocument)
	if err != nil || result == nil {
		http.Error(w, failMessage, http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}