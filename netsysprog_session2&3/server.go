package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
  //  unknown data types? don't worry, do the python way, which is to initialize with :=
  listener, err := net.Listen("tcp", "localhost:8080") 

  // null in go is nil
  if err != nil {
    fmt.Println(err)
    return
  }

  defer listener.Close() // closing the server, you know after all went run, we're gonna defer this one

  for {
    // the _  for ignoring the value, in there the variable that should be placed upon _ is conn
    conn,err := listener.Accept() // the conn is to know which connects to the gateway

    if err != nil {
      fmt.Println(err);
      return
    }

    go func(){ // defined anonymous & asynchronous function
      fmt.Println("connection accepted")

      // method 1 : this line below is to accept and read the sent data from the client
      // buffer := make([]byte, 1024)
      // n, err := conn.Read(buffer)
      // if err != nil{
      //   fmt.Println(err)
      //   return
      // }
      // fmt.Printf("%s\n", buffer[:n]) // really like python, you can take something from 0, to n

      //notes: it's essential to know that the slicing is to print only the sent string in client,
     // cuz if you just let the buffer to be print in the %s, it will print till 1024,
    // which is mignt not have been the expected value
     
    
    // method 2nd : dynamic buffer allocation by using the scanner
    // it's equivalent to Scanner function in java
      // scanner := bufio.NewScanner(conn) // the parameter, is to point which variable to read
      // the bufio have several method that can be used
      // bufio.scanbytes -> read per bytes
      // bufio.scanLines -> read per enter
      // bufio.scanWords -> read per word 
      // scanner.Split(bufio.ScanLines) // in here we use the scan line
  
      // this function is to read every scanned words (in ths case), till it can't, then print it
      // for scanner.Scan() {
      //     fmt.Println(scanner.Text())
      // }
  

      // if this for some reason you've put up \n, yep the only text that will show is before that \n 
      // so what we're actually doin is the function that for as long it can scan print 
      // fmt.Println(scanner.Text()) // output the scanner string


      //method 3rd: using the decode function (user-defined)
      //this one below, we're gonna accept the value and err (if exist)

      // we can set the deadline with for example in below code, with time.now
      // as for deadline definition, it's for managing the time to get the respose
      conn.SetDeadline(time.Now().Add(5*time.Second))

      // so if you for some reason let's just say
      // conn.SetDeadline(time.Now().Add(0)) -> which means it's the now with the milliseconds
      // it won't pass, why? because by logic the server supposed to send NOW
      // P.S: NOW -> (yes the current millisecond, second, hours, etc),
      // when it might have been seconds/ minutes/ hours elapsed

      if err != nil{
        fmt.Println(err)
        return
      }

      payload, err := decode(conn)

      if err != nil{
        fmt.Println(err)
        return
      }

      fmt.Println(payload)
      // for the method 3 works: go run server.go && payload.go cuz the writeto is the payload.go function


    }() // the parenthesis "()" is to "call the anonymous function, but do it in goroutine
  }
}
