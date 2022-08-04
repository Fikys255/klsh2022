package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type ClientMap map[string]net.Conn

func main() {
	var ln net.Listener
	var err error

	for {
		fmt.Println("Try to Launching server...")

		ln, err = net.Listen("tcp", ":8081")

		if err != nil {
			log.Println(err)
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	var Clients ClientMap = make(ClientMap)
	var Messages chan string = make(chan string)

	go TransmitMessages(Clients, Messages)

	for {
		if conn, err := ln.Accept(); err == nil {
			go ProcessMessages(conn, Clients, Messages)
			continue
		}
		log.Println(err)
	}
}

func TransmitMessages(cl ClientMap, msg <-chan string) {
	for {
		select {
		case m := <-msg:
			fmt.Print("Message Received: ", m)
			s := strings.Split(m, ":> ")

			if len(s) < 2 {
				continue
			}

			username := s[0]
			message := s[1]

			// fmt.Printf("\n%v\t%v", username, message)
			mm := strings.ToUpper(message) + "\n"

			if recipient, ok := cl[username]; ok {
				recipient.Write([]byte(mm))
			}
		}
	}
}

//
func ProcessMessages(conn net.Conn, cl ClientMap, msg chan<- string) {
	reader := bufio.NewReader(conn)
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	cl[username] = conn
	// fmt.Printf("%#v\n", cl)

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		message = strings.TrimSpace(message)

		if err != nil {
			if err == io.EOF {
				return
			} else {
				log.Println(err)
				continue
			}
		}

		msg <- message + "::" + username
	}
}
