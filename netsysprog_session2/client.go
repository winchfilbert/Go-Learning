package main

import (
	"fmt"
	"net"
)

func main(){
  conn, err := net.Dial("tcp", "localhost:8080") // map the data and err, as usual
  if err != nil{
    fmt.Println(err);
    return
  }
  // the data sending for request simulation
  conn.Write([]byte("filbert"))
  defer conn.Close() // after all the code, ran close the conn
}
