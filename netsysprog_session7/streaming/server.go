package main

import (
	"io/ioutil"
	"main/handler"
	"os"
	"path/filepath"
	"fmt"
	"net"
)

func main() {
	// making temporary directory, arg1 = null -> using default path
	// arg2 = the name of the directory that will be used
	dir, err := ioutil.TempDir("", "streaming_session7")

	// in order to import a function/ variable,
	// the exported function/variable must start with uppercase
	handler.ErrorHandler(err);

	// after the process done, we want the temporary path, and
	// it's content gone
	defer func(){
		err = os.RemoveAll(dir)
		handler.ErrorHandler(err);
	}()

	// gain processid
	socket := filepath.Join(dir, fmt.Sprintf("%d.sock", os.Getpid()));
	listener, err := net.Listen("unix",socket);
	handler.ErrorHandler(err)
	fmt.Printf("listening at %s\n", listener.Addr());
	defer listener.Close()

	// loop for taking the connection from client
	for{
		conn, err := listener.Accept()
		handler.ErrorHandler(err)
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	//handling with the 1 method, conn.read, static allocation, reading
	defer conn.Close();
	buf :=  make([]byte,1024);
	n, err := conn.Read(buf);
	handler.ErrorHandler(err);

	fmt.Printf("Received from client: %s\n", string(buf[:n]))

    // handling write with as long as the buf
    _, err = conn.Write(buf[:n]);
    handler.ErrorHandler(err);
    
    // Log that we sent a response
    fmt.Printf("Sent response to client: %s\n", string(buf[:n]))
}