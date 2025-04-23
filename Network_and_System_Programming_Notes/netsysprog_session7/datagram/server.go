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
	dir, err := ioutil.TempDir("", "datagram_session7");
	handler.ErrorHandler(err);
	fmt.Printf("Temp dir: %s\n", dir);


// if the directory done used, 
// we're gonna wipe the directory and it's content using os 
	defer func(){
		err := os.RemoveAll(dir)
		handler.ErrorHandler(err)
	}()

	// making server socket file
	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()))
	listener, err := net.ListenPacket("unixgram",socket)
	handler.ErrorHandler(err)
	fmt.Printf("Listening at: %s\n", listener.LocalAddr())
	defer listener.Close();

	// we need to manage the access server socket file
	// so that the client can read the directory

	// the os.Modesocket is to ensure or tell the system that the things we want to change is the socket
	// we're gonna handle the permission for the socket with 6 -> user , 2 -> group , 2 -> users
	// plz refer to the image.png in the datagram folder for permisson
	// add 0 in front of 622 in order to tell the system that it's the permission number
	err = os.Chmod(socket, os.ModeSocket | 0622)
	handler.ErrorHandler(err)

	buf := make([]byte, 1024);

	// read the request from the client 
	for {
		n, client, err := listener.ReadFrom(buf);
		handler.ErrorHandler(err);

		_, err = listener.WriteTo(buf[:n], client)
		handler.ErrorHandler(err);
	}

}

//N3xtG3nC0mput1ng