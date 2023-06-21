package parser

import (
	"bytes"
	"strconv"
	"strings"
	"fmt"


	"github.com/pieterclaerhout/go-log"
)

func RESPParser(input []byte) (string, []string) {
	InputString := string(input)
	BytesBuffer := bytes.NewBuffer(input)
	log.Infof("InputString Command: %q", BytesBuffer.Bytes())

	var commandLength = InputString[1:2]
	var keyLength string
	var arguments []string
	var ans []string

	if commandLength == "3" {
		keyLength = InputString[14:15]
		var valueLength int
		var key string
		var value string

		log.Infof("Command Length is %s\n", commandLength)
		log.Infof("Key Length is %s\n", keyLength)

		var trimFormat = fmt.Sprintf("*3\r\n$%s\r\nSET\r\n$%s\r\n", commandLength, keyLength)

		InputString = strings.TrimLeft(InputString, trimFormat)
		InputString = strings.TrimRight(InputString, "\r\n")

		keyLength, _ := strconv.Atoi(keyLength)
		log.Infof("No of Bytes of InputString %d", len(InputString))

		key = InputString[:keyLength]
		valueLength = keyLength + 6
		value = InputString[valueLength:]

		arguments = append(arguments, key)
		arguments = append(arguments, value)

		log.Infof("InputString is %s", InputString)
		log.Infof("Key is %s", key)
		log.Infof("Value is %s", value)
		return "SET", arguments

	} else if commandLength == "2" {
		
		keyLength = InputString[14:15]
		log.Infof("Command Length is %s\n", commandLength)
		log.Infof("Key Length is %s\n", keyLength)

		var trimFormat = fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%s\r\n", keyLength)

		InputString = strings.TrimLeft(InputString, trimFormat)
		InputString = strings.TrimRight(InputString, "\r\n")

		var ans []string
		ans = append(ans, InputString)
		return "GET", ans
	}
	return "PING", ans

}
