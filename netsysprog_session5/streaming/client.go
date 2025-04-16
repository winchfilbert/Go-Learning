package main

import (
    "fmt"
    "main/handler"
    "net"
)

func main(){
     // Update this to match the address printed by your server
     tempServerAddr := "/var/folders/wz/d49gg8gs3_xck21v82fbfn_r0000gn/T/streaming_session71668480609/45120.sock";
     client, err := net.Dial("unix", tempServerAddr);
     handler.ErrorHandler(err);
     defer client.Close();

     msg := []byte("hallo cok");
     
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