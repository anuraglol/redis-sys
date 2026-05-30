package server

import (
	"fmt"
	"io"
	"log"
	"net"
	core "rediss/core"
	"strings"
)

func readCommand(conn io.ReadWriter) (*core.RedisCmd, error) {
	var buf []byte = make([]byte, 512)
	n, err := conn.Read(buf[:])
	if err != nil {
		return nil, err
	}

	tokens, err := core.DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}

	return &core.RedisCmd{
		Cmd:  strings.ToUpper(tokens[0]),
		Args: tokens[1:],
	}, nil
}

func writeResponse(msg string, conn net.Conn) error {
	response := fmt.Sprintf(msg + "\n")
	_, err := conn.Write([]byte(response))
	return err
}

func respondError(err error, c io.ReadWriter) {
	c.Write([]byte(fmt.Sprintf("-%s\r\n", err)))
}

func respond(cmd *core.RedisCmd, c io.ReadWriter) {
	err := core.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(err, c)
	}
}

func RunSyncTCPServer() {
	log.Println("starting tcp server on localhost 8000")
	conn_clients := 0

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("error captured while starting the server: %v", err.Error())
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		conn_clients++
		log.Println("concurrent clients:", conn_clients)
		log.Println("client connected with address", conn.RemoteAddr())

		for {
			cmd, err := readCommand(conn)
			if err != nil {
				conn.Close()
				conn_clients--
				log.Println("client disconnected")
				if err == io.EOF {
					break
				}
				log.Println("error: ", err)
			}
			respond(cmd, conn)
		}
	}
}
