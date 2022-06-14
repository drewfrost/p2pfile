package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

const (
	TCP         = "tcp"
	UDP         = "udp"
	GOOGLE_DNS  = "8.8.8.8:80"
	CMD_CONNECT = "connect"
	CMD_LIST    = "list"
	BUFF_SIZE   = 256
)

func main() {
	go client()
	server()
}

func client() {
	for {
		args := getArgs()
		if strings.Split(args, " ")[0] == CMD_CONNECT {
			connectionIP := strings.Split(args, " ")[1]
			conn, err := net.Dial(TCP, connectionIP)
			if err != nil {
				log.Print("Error connection to server from client: ", err)
			}
			go sendMessage(conn)
		}
	}
}

func sendMessage(conn net.Conn, message []byte) {
	_, err := conn.Write(message)
	if err != nil {
		log.Print("Error send message to server: ", err)
	}

}

func getArgs() string {
	args, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Print("Error reading client arguments: ", err)
	}
	return args
}

func server() {
	localAddrs := getLocalAddress()
	log.Print("Local ip address: ", localAddrs)
	listener, err := net.Listen(TCP, localAddrs)
	if err != nil {
		log.Print("Error listeting to local address:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Error accepting the connection: ", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	message := make([]byte, BUFF_SIZE)
	n, err := conn.Read(message)
	if err != nil {
		log.Print("Error reading message from the connection: ", err)
	}
	log.Print(len(message[:n]))
}

func getLocalAddress() string {
	log.Print("Connected to DNS")
	conn, err := net.Dial(UDP, GOOGLE_DNS)
	if err != nil {
		log.Print("Error in DNS connection: ", err)
	}
	defer conn.Close()
	localAddress := conn.LocalAddr().String()
	return localAddress
}
