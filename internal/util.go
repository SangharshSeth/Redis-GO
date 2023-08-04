package internal

import (
	"fmt"
	"net"
)

func SendRedisResponse(msg any, connection net.Conn) {
	response := fmt.Sprintf("+%s\r\n", msg)
	_, err := connection.Write([]byte(response))
	if err != nil {
		panic(err.Error())
	}
}
