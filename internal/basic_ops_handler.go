package internal

import (
	"github.com/pieterclaerhout/go-log"
	"net"
	"sync"
)

func HandleBasicOps(arr []string, storage *sync.Map, connection net.Conn) {
	switch arr[0] {
	case "SET":
		log.Info("SET")
		result, _ := SetKey(arr[1], arr[2], storage)
		SendRedisResponse(result, connection)
	case "GET":
		log.Info("GET")
		result, _ := GetKey(arr[1], storage)
		SendRedisResponse(result, connection)
	case "EXPIRY":
		log.Info("Expiry")
	case "DEL":
		log.Info("Del")
		result, _ := DelKey(arr[1], storage)
		SendRedisResponse(result, connection)
	}
}
