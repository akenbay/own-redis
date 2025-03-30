package handler

import (
	"log/slog"
	"net"
	"strconv"
	"strings"
	"time"

	. "own-redis/data"
)

func Set(msg string, addr *net.UDPAddr, conn *net.UDPConn) {
	// Split statment to get keys and values
	keyVals := strings.Split(msg, " ")

	// Error statement for invalid request
	if len(keyVals) <= 2 {
		conn.WriteToUDP([]byte("(error) ERR wrong number of arguments for 'SET' command\n"), addr)

		slog.Error("Wrong Number of Arguments for SET")

		return
	}

	// Check if statement has a 'px' flag
	if SetPXRe.MatchString(strings.ToUpper(msg)) {
		// Stores the value omitting 'px' part
		Data.Store(keyVals[1], strings.Join(keyVals[2:len(keyVals)-2], " "))
		conn.WriteToUDP([]byte("OK\n"), addr)
		slog.Info("Set Key: " + keyVals[1])

		// Gets the number of millisecond form 'px' flag
		t, err := strconv.Atoi(keyVals[len(keyVals)-1])
		if err != nil || t < 0 {
			conn.WriteToUDP([]byte("(error) ERR wrong input for px\n"), addr)

			slog.Error("Wrong Input for px")

			return
		}

		// lambda function that deletes the values after sleep time
		go func() {
			time.Sleep(time.Duration(t) * time.Millisecond)
			Data.Delete(keyVals[1])
		}()

	} else {
		// Stores the values and logs the info
		Data.Store(keyVals[1], strings.Join(keyVals[2:], " "))
		slog.Info("Set key: " + keyVals[1])
		conn.WriteToUDP([]byte("OK\n"), addr)
	}
}
