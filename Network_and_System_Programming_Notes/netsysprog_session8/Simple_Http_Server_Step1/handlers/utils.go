package handlers;

import (
	"fmt"
	"net/http"
)

func PrintHelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome...")	
}

