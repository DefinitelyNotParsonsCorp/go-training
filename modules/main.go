package main

import (
	"net"

	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func RunServer(c *cli.Context) error {

	l, err := net.Listen("tcp", c.String("listen-addr"))
	if err != nil {
		log.Fatal("Couldn't listen on socket",
			zap.String("listen-addr", c.String("listen-addr")),
		)
	}
	defer l.Close()

	log.Info("Now accepting client connections",
		zap.String("listen-addr", c.String("listen-addr")))

	for {
		conn, err := l.Accept()
		if err != nil {
			break
		}

		go handleConnection(conn)
	}

	return err
}

func handleConnection(c net.Conn) {
	log.Info("Client connect", zap.String("remote-addr", c.RemoteAddr().String()))
	defer c.Close()
	defer log.Info("Client disconnect", zap.String("remote-addr", c.RemoteAddr().String()))

	buf := make([]byte, 128)
	for {
		n, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[:n]
		c.Write(data)
	}
}
