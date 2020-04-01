package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Printf("Printing lines to stdout and echoing responses\n")
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

type LogWriter struct {
	Prefix string
	Writer io.Writer
}

func (w *LogWriter) Write(b []byte) (int, error) {
	log.Printf("%s: %s", w.Prefix, string(b))
	return w.Writer.Write(b)
}

func handleConnection(c net.Conn) {
	log.Printf("Accepted connection from %s\n", c.RemoteAddr())
	defer c.Close()
	defer log.Printf("Client %s disconnected.", c.RemoteAddr())

	writer := LogWriter{
		Prefix: c.RemoteAddr().String(),
		Writer: c,
	}

	buf := make([]byte, 128)
	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[:n]
		writer.Write(data)
	}
}
