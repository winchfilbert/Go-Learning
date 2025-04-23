package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type User struct{
	Name string
	Age int
}

func PrintHelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome...")	
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request){
	var data User
	json.NewDecoder(r.Body).Decode(&data);

	fmt.Println(data)


	w.Write([]byte("Success Get Data"))
}


