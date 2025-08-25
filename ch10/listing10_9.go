package ch10

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

var r, _ = regexp.Compile("GET (.+) HTTP/1.1\r\n")

func handleHttpRequest(conn net.Conn) {
	defer conn.Close() //nolint:errcheck
	buff := make([]byte, 1024)
	size, _ := conn.Read(buff)
	if r.Match(buff[:size]) {
		file, err := os.ReadFile(
			fmt.Sprintf("files/%s", r.FindSubmatch(buff[:size])[1])) //nolint:errcheck
		if err == nil {
			fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n", len(file)) //nolint:errcheck
			conn.Write(file)                                                              //nolint:errcheck
		} else {
			fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n\r\n<html>Not Found</html>") //nolint:errcheck
		}
	} else {
		fmt.Fprint(conn, "HTTP/1.1 500 Internal Server Error\r\n\r\n") //nolint:errcheck
	}
}

func StartHttpWorkers(n int, incomingConnections <-chan net.Conn) {
	for range n {
		go func() {
			for c := range incomingConnections {
				handleHttpRequest(c)
			}
		}()
	}
}
