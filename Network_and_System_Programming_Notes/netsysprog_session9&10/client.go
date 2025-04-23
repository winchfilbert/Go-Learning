package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct{
	Name string
	Age int
}

func main(){

	buffer := new(bytes.Buffer)
	user := User{
		Name: "Filbert", 
		Age: 20,
	}
	json.NewEncoder(buffer).Encode(&user)
	
	// this is example for method post, this is intended to test
	resp, err := http.Post("http://localhost:5555/sendData", "application/json", buffer)

	if err != nil{
		fmt.Println(err);
		return
	}

	// if you did somehow uncomment this line below, it just give raw type.,. 
	// data.readall(io)
	
	data, err := io.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err);
		return
	}

	// to read the data is better to parse with string, it's like typecase, but here is a bit smoother
	fmt.Printf("%s\n",string(data))
}


