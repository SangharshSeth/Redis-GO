package parser

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)


func RESPParser(input []byte) (string,[]string) {
	InputString := string(input)
	BytesBuffer := bytes.NewBuffer(input)
	log.Printf("InputString Command: %q", BytesBuffer.Bytes())

	var commandLength = InputString[1:2]
	var keyLength string 
	var arguments []string
	var ans []string

	if commandLength == "3" {
		keyLength = InputString[14:15]
		var valueLength int
		var key string
		var value string

		log.Printf("Command Length is %s\n", commandLength)
		log.Printf("Key Length is %s\n", keyLength)

		var trimFormat = fmt.Sprintf("*3\r\n$%s\r\nSET\r\n$%s\r\n", commandLength, keyLength)

		InputString = strings.TrimLeft(InputString, trimFormat)
		InputString = strings.TrimRight(InputString, "\r\n")


		keyLength, _ := strconv.Atoi(keyLength)
		log.Printf("No of Bytes of InputString %d", len(InputString))

		key = InputString[:keyLength]
		valueLength = keyLength + 6
		value = InputString[valueLength:]

		arguments = append(arguments, key)
		arguments = append(arguments, value)

		log.Printf("InputString is %s", InputString)
		log.Printf("Key is %s", key)
		log.Printf("Value is %s", value)
		return "SET", arguments

	} else if commandLength == "2" {
		keyLength = InputString[14:15]
//		var valueLength int
//		var key string
//		var value string

		log.Printf("Command Length is %s\n", commandLength)
		log.Printf("Key Length is %s\n", keyLength)

		var trimFormat = fmt.Sprintf("*2\r\n$3\r\nGET\r\n$%s\r\n", keyLength)


		InputString = strings.TrimLeft(InputString, trimFormat)
		InputString = strings.TrimRight(InputString, "\r\n")

		var ans []string
		ans = append(ans, InputString)
		return "GET", ans
	}
	return "PING", ans

}
