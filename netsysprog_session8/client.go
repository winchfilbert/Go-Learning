package main

import (
	"fmt"
	"io"
	"net/http"
)


func main(){
	// this is for method get
	resp, err := http.Get("http://localhost:5555/index")

	// this is example for method post, this is intended to test
	// if the method of server is wrong, it outputs "Method Wrong"
	// resp, err := http.Post("http://localhost:5555/index", "application/json", nil)

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
	fmt.Printf("%s",string(data))
}


