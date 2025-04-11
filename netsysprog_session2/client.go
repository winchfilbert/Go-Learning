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

  //for method 1 and method 2, please know that the conn.write works
  // // the data sending for request simulation
  // conn.Write([]byte("filbert"))

  // but for the method 3, we're gonna simulate this in below lines of code
  payload := Binary("Filbert Christian Winch") // i'm picking the binary type dude
  _, err = payload.WriteTo(conn)
  
  //so what that payload.writeto doing in the background is that
  // determine what type
  // determine the size
  // so on, got either error or return me the response

  if err != nil{
    fmt.Println(err)
    return
  }

  // for the method 3 works: go run client.go && payload.go cuz the writeto is the payload.go function

  // closing connection
  defer conn.Close() // after all the code, ran close the conn
}
