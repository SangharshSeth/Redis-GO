package internal

import (
	"bytes"
	"github.com/pieterclaerhout/go-log"
	"strconv"
)

func RESPParser(input []byte) []string {
	inputStringBufferBytes := bytes.NewBuffer(input)

	var (
		bulkArraySize  int
		parsedCommands []string
		clientLen      int
		tempString     string
		noDigits       int
	)

	var commandPrintable string
	for _, item := range inputStringBufferBytes.Bytes() {
		if item == 13 {
			continue
		} else if item == 10 {
			commandPrintable += "*"
			continue
		} else {
			commandPrintable += string(item)
		}
	}

	log.Infof("Printable Command is %s", commandPrintable)

	//var commandLength = inputString[1:2]
	bulkArraySize, _ = strconv.Atoi(commandPrintable[1:2])

	log.Info(bulkArraySize) //this is for assertion

	for i := 0; i < len(commandPrintable); i++ {
		if commandPrintable[i] == '$' {
			var clientLength string

			for j := i + 1; j < len(commandPrintable); j++ {
				if string(commandPrintable[j]) == "*" {
					break
				}
				clientLength += string(commandPrintable[j])
			}

			clientLen, _ = strconv.Atoi(clientLength)
			noDigits = len(strconv.Itoa(clientLen))

			tempString = commandPrintable[i+noDigits+2 : i+noDigits+2+clientLen]
			parsedCommands = append(parsedCommands, tempString)

		}
	}

	//log.Info("Parsed Commands Are Written Below.")
	log.Info(parsedCommands)

	return parsedCommands
}
