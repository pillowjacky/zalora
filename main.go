package main

import (
	"os"
	"fmt"
	// "net"
	"strings"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("should be 2 args: tcp host:port, command")
		return
	}

	command := args[1]
	if !strings.Contains(command, "GET") {
		fmt.Println("command format invalid")
		return
	}

	splitedCommand := strings.Split(command, "GET")
	
	for _, d := range splitedCommand {
		var commandLenght int64
		commandLenght, err := strconv.ParseInt(d, 10, 32)
		if err != nil {
			fmt.Println("command format invalid")
		}
		fmt.Println(commandLenght)
	}
	
	// if _, err := strconv.Atoi(v); err == nil {
	// }

	host := args[0]
	fmt.Println(host)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println(err)
	}
}