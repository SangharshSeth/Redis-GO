package internal

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
	"net"
)

func SendRedisResponse(msg any, connection net.Conn) {
	log.Info(msg)
	response := fmt.Sprintf("+%s\r\n", msg)
	_, err := connection.Write([]byte(response))
	if err != nil {
		panic(err.Error())
	}
}
