package main

import (
	"main/handlers"
	"fmt"
	"net/http"
)

// if you somehow want to make relative path to / and which method
// which called middleware

func middleware(method string, handlerFunc http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request){
			if r.Method != method{
				// give response with w variable if method is wrong
				fmt.Fprintln(w, "Method wrong")
				return
			}
			handlerFunc(w, r)
		}
}

func main(){
	// making the multiplexer
	mux := http.NewServeMux()

	// this is example if you didn't use the middleware func
	// mux.HandleFunc("/index", handlers.PrintHelloHandler)


	// this if u implement the middleware
	mux.HandleFunc("/index", middleware(http.MethodGet, handlers.PrintHelloHandler))


	server := http.Server{
		Addr: "localhost:5555",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		fmt.Println("error occured")
		return
	}
}