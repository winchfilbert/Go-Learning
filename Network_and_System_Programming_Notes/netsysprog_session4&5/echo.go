package main

import (
	"context"
	"fmt"
	"net"
)

func echoServerUDP(ctx context.Context, addr string)(net.Addr, error){
  listener, err := net.ListenPacket("udp", addr)
  
  if err != nil{
    return nil, err
  }

  go func(){
    
    // this goroutine is to check if that the context is done or not
    go func(){
      <-ctx.Done()
      listener.Close()
    }()
    // if there's data listened buf is used to contain the data
    buf := make([]byte, 1024)
    for{
      // no checking in UDP so we can just directly use readfrom
      // because it's UDP, the sender might have been not the client, so we handle using n and clientaddr 
      n, clientAddr, err := listener.ReadFrom(buf)

      if err != nil{
        fmt.Println(err)
        return 
    }

      _, err = listener.WriteTo(buf[:n], clientAddr)
    }
  }()

  // this below is for returning the sender address, and null if error
  return listener.LocalAddr(), nil
}
