package controllers

import "net/http"

func WhoIP(w http.ResponseWriter, r *http.Request){
	reqIP := r.RemoteAddr;
	w.Write([]byte(reqIP))
}