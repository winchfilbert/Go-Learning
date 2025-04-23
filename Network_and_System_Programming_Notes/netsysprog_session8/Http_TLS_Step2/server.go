package main

import (
	"main/handlers"
	"fmt"
	"net/http"
)

func middleware(method string, handlerFunc http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request){
			if r.Method != method{
				fmt.Fprintln(w, "Method wrong")
				return
			}
			handlerFunc(w, r)
		}
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/index", middleware(http.MethodGet, handlers.PrintHelloHandler))


	server := http.Server{
		Addr: "localhost:5556",
		Handler: mux,
	}
	
	// putting the TLS here, to differentiate the sessions
	// have to download additional package here, named mkcert (this is for make certificate)
	// link to mkcert download : https://github.com/FiloSottile/mkcert
	// rename it (plz refer to README.md)
	// then put those as arguments	
	err := server.ListenAndServeTLS("cert.pem", "key.pem")
	if err != nil{
		fmt.Println("error occured")
		return
	}
}