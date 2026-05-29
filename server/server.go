package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func readCommand(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return message, nil
}

func writeResponse(msg string, conn net.Conn) error {
	response := fmt.Sprintf(msg + "\n")
	_, err := conn.Write([]byte(response))
	return err
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

			log.Println("you said", cmd)
			if err = writeResponse(cmd, conn); err != nil {
				log.Println("error writing: ", err)
			}
		}
	}
}
