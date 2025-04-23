package main

import (
    "fmt"
    "main/handler"
    "net"
)

func main(){
   // in the video, you have to run server.go, then manually
   // add the shown address at server.go, then before running
   // go run client.go
   // you have to edit and paste the address here
   // then go run client.go to get the feedback

     tempServerAddr := "";
     client, err := net.Dial("unixpacket", tempServerAddr);
     handler.ErrorHandler(err);
     defer client.Close();

     msg := []byte("hallo packet");
     for i := 0 ; i < 3; i++{
        _, err = client.Write(msg);
        handler.ErrorHandler(err);
     }

     buf := make([]byte, 1024);
     for i := 0 ; i < 3 ; i++{
        n, err := client.Read(buf);
        handler.ErrorHandler(err);

        fmt.Println(string(buf[:n]));
     }
}