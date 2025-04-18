package main

import (
    "fmt"
    "main/handler"
    "net"
)

func main(){
     tempServerAddr := "";
     client, err := net.Dial("unix", tempServerAddr);
     handler.ErrorHandler(err);
     defer client.Close();

     msg := []byte("hallo packet");
     for i := 1 ; i < 3; i++{
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