package handler

import (
	"log/slog"
	"net"
	"strings"

	. "own-redis/data"
)

func RequestHandler(msg string, addr *net.UDPAddr, conn *net.UDPConn) {
	// Split and trim statements used to trim data to request only and get rid of garbage data
	msg = strings.Split(msg, "\n")[0]
	msg = strings.TrimFunc(msg, func(c rune) bool {
		return c == '\n' || c == ' '
	})

	// Switch statement to check which request it is
	switch {
	case SetRe.MatchString(strings.ToUpper(msg)):
		go Set(msg, addr, conn) // 'go' keyword is used to avoid data racing
		return
	case GetRe.MatchString(strings.ToUpper(msg)):
		go Get(msg, addr, conn) // 'go' keyword is used to avoid data racing
		return
	case strings.ToUpper(msg) == "PING":
		conn.WriteToUDP([]byte("PONG\n"), addr)
		slog.Info("PONG")
		return
	default:
		// Error statement for invalid request
		conn.WriteToUDP([]byte("(error) ERR invalid request\n"), addr)
		slog.Error("Invalid Request")

		return
	}
}
