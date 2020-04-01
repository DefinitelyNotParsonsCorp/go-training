package main

import (
	"log"
	"net"
)

type ClientConnection struct {
	net.Conn
	closeChan     chan string
	broadcastChan chan []byte
	sendChan      chan []byte
}

func (cc *ClientConnection) Close() error {
	cc.closeChan <- cc.RemoteAddr().String()
	return cc.Conn.Close()
}

func NewConnectionAccepterChannel(addr string) (chan net.Conn, error) {

	acceptChan := make(chan net.Conn, 1)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer l.Close()
		for {
			conn, err := l.Accept()
			if err != nil {
				return
			}
			acceptChan <- conn
		}
	}()

	return acceptChan, nil
}

func main() {

	broadcastChan := make(chan []byte, 1)
	closeChan := make(chan string, 1)
	clients := make(map[string]chan []byte)
	acceptChan, _ := NewConnectionAccepterChannel(":1234")

	for {
		select {
		// On new connections, spawn a new goroutine to service them.
		case conn := <-acceptChan:
			// Holds client information.
			client := &ClientConnection{
				Conn: conn,
				// This channel is the "send channel", which allows us to send messages
				// to this connection ONLY.
				sendChan:      make(chan []byte, 1),
				broadcastChan: broadcastChan,
				closeChan:     closeChan,
			}

			// We could/should hold the entire connection, but we only need the channel
			clients[conn.RemoteAddr().String()] = client.sendChan

			go handleConnection(client)

		// In this case, a client wants to broadcast to all the connected clients.
		case msg := <-broadcastChan:
			for _, sendChan := range clients {
				sendChan <- msg
			}
		case remoteaddr := <-closeChan:
			delete(clients, remoteaddr)
		}
	}
}

func handleConnection(c *ClientConnection) {

	log.Printf("Accepted connection from %s\n", c.RemoteAddr())
	defer log.Printf("Client %s disconnected.", c.RemoteAddr())
	defer c.Close()

	// Need to send messages to the masses.
	go func() {
		for {
			msg := <-c.sendChan
			c.Write(msg)
		}
	}()

	buf := make([]byte, 128)
	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[:n]
		c.broadcastChan <- data
	}
}
