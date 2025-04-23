package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type User struct{
	Name string
	Age int
}

func PrintHelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome...")	
	log.Println("User Access Index")
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request){
	var data User
	err := json.NewDecoder(r.Body).Decode(&data);
	if err != nil{
		// this line below shuts the system
		// log.Fatal(err)
		log.Println(err)
		return
	}

	//put a log in here
	log.Println("User Sent a data")
	fmt.Println(data)


	w.Write([]byte("Success Get Data"))
}


