package parser

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)


func RESPParser(input string) (string,[]string) {
	
	var commandLength = input[1:2]
	var keyLength string 
	var arguments []string
	var ans []string

	if commandLength == "3" {
		keyLength = input[14:15]
		var valueLength int
		var key string
		var value string

		log.Printf("Command Length is %s\n", commandLength)
		log.Printf("Key Length is %s\n", keyLength)

		var trimFormat = fmt.Sprintf("*3\r\n$%s\r\nSET\r\n$%s\r\n", commandLength, keyLength)

		input = strings.TrimLeft(input, trimFormat)
		input = strings.TrimRight(input, "\r\n")


		keyLength, _ := strconv.Atoi(keyLength)
		log.Printf("No of Bytes of input %d", len(input))

		key = input[:keyLength]
		valueLength = keyLength + 6
		value = input[valueLength:]

		arguments = append(arguments, key)
		arguments = append(arguments, value)

		log.Printf("Input is %s", input)
		log.Printf("Key is %s", key)
		log.Printf("Value is %s", value)
		return "SET", arguments

	} else if commandLength == "2" {
		var length = 3
		var keyLength = input[14:15]

		var trimFormat = fmt.Sprintf("*2\r\n$%d\r\nGET\r\n$%s", length, keyLength)
		input = strings.TrimRight(input, "\r\n")
		input = strings.TrimLeft(input, trimFormat)

		fmt.Printf("Data is %s", (input))

		arguments = append(arguments, input)

		return "GET", arguments
	}
	return "PING", ans

}
