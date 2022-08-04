package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	fmt.Fprintf(conn, username+"\n")

	go ReceiveMessage(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}

func ReceiveMessage(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		s := strings.Split(message, "::")
		fmt.Printf("\rMessage from %v: %vText to send: ", s[1], s[0])
	}
}
