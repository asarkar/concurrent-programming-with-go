package main

import (
	listing10_9 "concurrent-programming-with-go/ch10"
	"fmt"
	"net"
)

/*
2. Change listing 10.12 so that the work queue channel between the main() goroutine and the worker
pool has a buffer of 10 messages. Doing so will give you a capacity buffer so that when all the
goroutines are busy, some of the requests are queued before they can be picked up.

ANSWER:
Simply add a buffer to the connections channel.

Usage:
1. Start the server: go run ch10/ex10.2/httpserver.go &
2. Request a file: curl localhost:8080/file1.txt
*/
func main() {
	incomingConnections := make(chan net.Conn, 10)
	listing10_9.StartHttpWorkers(3, incomingConnections)
	server, _ := net.Listen("tcp", "localhost:8080")
	defer server.Close() //nolint:errcheck
	for {
		conn, _ := server.Accept()
		select {
		case incomingConnections <- conn:
		default:
			fmt.Println("Server is busy")
			fmt.Fprintln(conn, "HTTP/1.1 429 Too Many Requests\r\n\r\n<html>Busy</html>") //nolint:errcheck
			conn.Close()                                                                  //nolint:errcheck
		}
	}
}
