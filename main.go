package main

import (
	"bytes"
	"fmt"
	"github.com/sangharshseth/internal"
	"io"
	"net"
	"os"
	"sync"

	"github.com/pieterclaerhout/go-log"
)

var BasicOps = map[string]bool{
	"SET":    true,
	"GET":    true,
	"DEL":    true,
	"EXPIRY": true,
	"INCR":   true,
	"DECR":   true,
	"TTL":    true,
}

func HandleConnection(connection net.Conn, storage *sync.Map) {

	defer connection.Close()
	var buffer = make([]byte, 128)
	for {
		n, err := connection.Read(buffer)
		buffer = buffer[:n]
		if err != nil {
			if err == io.EOF {
				log.Error("Connection closed by client")
				break
			}
			log.Errorf("Error while reading data: %s", err.Error())
			return
		}

		buffer = bytes.TrimSpace(buffer)
		command := internal.RESPParser(buffer)

		if BasicOps[command[0]] {
			log.Info("Basic Operation")
			internal.HandleBasicOps(command, storage, connection)
		}

		//connection.Write([]byte("+OK\r\n"))
	}

}

func ExpiryService(datastore *sync.Map) {
	//TODO: Implement it

}

func main() {

	//Initialize Loggers
	log.DebugMode = true
	log.PrintTimestamp = true
	log.PrintColors = true

	port := 6379
	var storageEngine sync.Map
	var redisServerURL = fmt.Sprintf("127.0.0.1:%d", port)

	//Background Service to Handle Expiry
	go ExpiryService(&storageEngine)

	l, err := net.Listen("tcp", redisServerURL)

	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	log.Infof("Redis Server is Listening on PORT %d", port)

	//Loop to handle Concurrent Clients
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Infof("Error accepting connection: %s", err.Error())
			os.Exit(1)
		}
		log.Infof("Custom Redis-Server is Listening for connections on port %d\n", port)
		go HandleConnection(conn, &storageEngine)
	}
}
