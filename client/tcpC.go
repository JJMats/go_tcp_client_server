package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		// Send text through the connection to the TCP server
		fmt.Fprintf(c, text+"\n")

		// Read the TCP server's response
		message, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println("Server returned error: " + err)
		}
		fmt.Print("->: " + message)

		// Listen for exit message
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
