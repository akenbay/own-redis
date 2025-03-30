package handler

import (
	"log/slog"
	"net"
	"strings"

	. "own-redis/data"
)

func Get(msg string, addr *net.UDPAddr, conn *net.UDPConn) {
	// Split request body to get key
	keyVals := strings.Split(msg, " ")

	// Error message for invalid request
	if len(keyVals) != 2 {
		conn.WriteToUDP([]byte("(error) ERR wrong number of arguments for 'GET' command\n"), addr)

		slog.Error("Wrong Number of Arguments for GET")

		return
	}

	// Load data from Data Map
	val, _ := Data.Load(keyVals[1])
	// Check if value exists and is string(needed to assert data)
	if str, ok := val.(string); ok {
		conn.WriteToUDP([]byte(str+"\n"), addr)
	} else {
		conn.WriteToUDP([]byte("(nil)\n"), addr)
	}
	slog.Info("Get Key: " + keyVals[1])
}
