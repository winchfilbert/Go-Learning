package main

import (
	"context"
	"fmt"
	"net"
)


func main(){

  // if the ctx want to cancel, 
  ctx, cancel := context.WithCancel(context.Background())
  serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
  
  if err != nil{
    fmt.Println(err)
    return
  }

  // this is for, if the ctx is done and we want to cancel, use the defer cancel to cancel
  defer cancel()

  client, err := net.ListenPacket("udp", "127.0.0.1:")

  defer client.Close()

  

  // sent the buffer / write a message to server
  _, err = client.WriteTo([]byte("hallo"), serverAddr)

  if err != nil{
    fmt.Println(err)
    return
  }


  // read the buffer / read message from server (the response)
  buf := make([]byte, 1024)
  n, addr, err := client.ReadFrom(buf)

  if err != nil{
    fmt.Println(err)
    return
  }

  fmt.Println(string(buf[:n]))
  fmt.Println(addr)
}
