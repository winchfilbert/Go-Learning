package main

import (
	"fmt"
	"io"
	"net/http"
)


func main(){
	// the link to request needs https
	resp, err := http.Get("https://localhost:5556/index")

	if err != nil{
		fmt.Println(err);
		return
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil{
		fmt.Println(err);
		return
	}

	fmt.Printf("%s",string(data))
}


