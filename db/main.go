package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			log.Println("out of loop")
			log.Print(err)
			return
		}
		go mustCopy(os.Stdout, c)
		time.Sleep(1 * time.Second)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.CopyN(dst, src, 4); err != nil {
		log.Println("exiting through mustCopy")
		log.Fatal(err)
	}
}
