package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/sangharshseth/src/parser"
)

func handleConnection(connection net.Conn, storage *sync.Map) {

	defer connection.Close()
	var buffer = make([]byte, 128)
	for {
		n, err := connection.Read(buffer)
		buffer = buffer[:n]
		if err != nil {
			if err == io.EOF {
				log.Println("Connection closed by client")
				break
			}
			log.Printf("Error while reading data: %s", err.Error())
			return
		}

		//RESP Parser
		command, arguments := parser.RESPParser(buffer)

		//Handle Commands
		if command == "SET" {
			storage.Store(arguments[0], arguments[1])
			connection.Write([]byte("+OK\r\n"))
		}
		if command == "GET" {
			val, exists := storage.Load(arguments[0])
			if !exists {
				connection.Write([]byte("+Nil\r\n"))
			} else {
				result := fmt.Sprintf("+%s\r\n", val)
				resultInterMediate := strings.TrimLeft(result, "+\n")
				result = fmt.Sprintf("+%s", resultInterMediate)

				connection.Write([]byte(result))
			}
		}
		connection.Write([]byte("+OK\r\n"))
	}

}

func ExpiryService(datastore *sync.Map) {

}

func main() {
	port := "6379"
	var storageEngine sync.Map
	var redisServerURL = fmt.Sprintf("127.0.0.1:%s", port)

	//Background Service to Handle Expiry
	go ExpiryService(&storageEngine)

	l, err := net.Listen("tcp", redisServerURL)

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			os.Exit(1)
		}
		log.Printf("Custom Redis-Server is Listening for connections on port %s\n", port)
		go handleConnection(conn, &storageEngine)
	}
}
