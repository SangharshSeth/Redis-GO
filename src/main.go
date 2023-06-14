package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

		command, arguments := parser.RESPParser(string(buffer))

		if command == "PING" {
			connection.Write([]byte("+PONG\r\n"))
		} else if command == "SET" {
			storage.Store(arguments[0], arguments[1])
			log.Printf("Data is Stored in Redis-Storage Successfully.")
			connection.Write([]byte("+OK\r\n"))
		} else if command == "GET" {
			data, err := storage.Load(arguments[0])
			if !err  {
				fmt.Printf("Data is not Found in Redis-Storage.")
				connection.Write([]byte("+Data Not Found in Redis Storage\r\n"))
			}
			response := fmt.Sprintf("+%s\r\n",data)
			connection.Write([]byte(response))
		}
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
