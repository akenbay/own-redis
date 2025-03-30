package main

import (
	"flag"
	"fmt"
	"net"
)

var PortNum int

func main() {
	flag.IntVar(&PortNum, "port", 8080, "Port number")
	flag.Usage = func() {
		fmt.Print("Own Redis\n\n",
			"**Usage**\n",
			"\town-redis [-port <N>]\n\town-redis --help\n\n",
			"**Options:**\n",
			"- --help\tShow this screen.\n",
			"- --port N\tPort number\n")
	}
	flag.Parse()

	port := ":" + fmt.Sprint(PortNum)

	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error creating UDP server:", err)
		return
	}
	defer conn.Close()
}
