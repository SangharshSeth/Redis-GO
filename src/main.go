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
				// Reached end of file (connection closed), exit the loop
				log.Println("Connection closed by client")
				break
			}
			log.Printf("Error while reading data: %s", err.Error())
			return
		}

		log.Printf("No of Bytes Read: %d\n", n)

		command, arguments := parser.RESPParser(buffer)

		if command == "SET" {
			storage.Store(arguments[0], arguments[1])
			connection.Write([]byte("+OK\r\n"))
		}
		if command == "GET" {
			val, exists := storage.Load(arguments[0])
			if !exists {
				log.Printf("Here %s", val)
				log.Printf("Value not Found for Key %s", arguments[0])
				connection.Write([]byte("+No Key\r\n"))
			} else {
				log.Printf("Value for Key %s is %s", arguments[0], val)
				result := fmt.Sprintf("+%s\r\n", val)
				resultInterMediate := strings.TrimLeft(result, "+\n")
				result = fmt.Sprintf("+%s", resultInterMediate)
				log.Print(len(result))
				log.Print([]byte(result))
				connection.Write([]byte(result))
			}
		}
		fmt.Printf("Command %s", command)
		fmt.Print(arguments[0])
		connection.Write([]byte("+OK\r\n"))
	}

}

func main() {

	port := "6379"
	var storageEngine sync.Map

	redisServerURL := fmt.Sprintf("127.0.0.1:%s", port)
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
