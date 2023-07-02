package main

import (
	server "github.com/gabrielluizsf/login-solid-octo-pancake/api/v1"
)
const PORT = ":8080"

func  main(){
	server.Start(PORT)
}