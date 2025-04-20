package main

import (
	"fmt"
	"io/ioutil"
	"main/handler"
	"net"
	"os"
	"path/filepath"
)


func main(){
	tempDir := ``;
	// we're making own socket for client
	clientSocket := filepath.Join(tempDir, fmt.Sprintf("client: %d.sock", os.Getpid()));
	serverAddr,err :=  net.ResolveUnixAddr("unixgram", tempDir);
	handler.ErrorHandler(err)
	listener, err := net.ListenPacket("unixgram", clientSocket);
	handler.ErrorHandler(err);
	defer listener.Close();


	// we're gonna do the same thing in client, that is to manage the permission using chmod
	err = os.Chmod(clientSocket, os.ModeSocket|0622)
	handler.ErrorHandler(err)


	msg := []byte("hallo datagram")	
	for i := 0 ; i < 3 ; i++{
		// bcs the address cannot be string, we use, net.ResolveUnixAddr, pls refer to the serveraddr
		_, err = listener.WriteTo(msg, serverAddr);
		handler.ErrorHandler(err)
	}

	buf := make([]byte, 1024);
	for i := 0 ; i < 3 ; i++{
		n, _, err := listener.ReadFrom(buf);
		handler.ErrorHandler(err);

		fmt.Println(string(buf[:n]))
	}
}