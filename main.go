package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"
	"strings"

	. "own-redis/handler"
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
		slog.Error(err.Error())
		os.Exit(1)

	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)

	}
	defer conn.Close()

	fmt.Printf("Server Started\nPort: %s\n", addr)

	buffer := make([]byte, 1024)

	for {
		// Read from UDP socket
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			slog.Error(err.Error())
			conn.WriteToUDP([]byte("(error) ERR "+err.Error()+"\n"), addr)
			os.Exit(1)
		}

		// Convert message to string and trim spaces
		message := strings.TrimSpace(string(buffer[:n]))

		RequestHandler(message, clientAddr, conn)
	}
}
