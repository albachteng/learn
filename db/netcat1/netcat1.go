package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var msg []byte
	for {
		_, err = os.Stdin.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(conn, string(msg))
		if err != nil {
			log.Fatal(err)
		}
		go mustCopy(os.Stdout, conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
