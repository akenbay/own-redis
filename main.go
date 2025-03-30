package main

import (
	"flag"
	"fmt"
)

var PortN int

func main() {
	flag.IntVar(&PortN, "port", 8080, "Port number")
	flag.Usage = func() {
		fmt.Print("Own Redis\n\n",
			"**Usage**\n",
			"\town-redis [-port <N>]\n\town-redis --help\n\n",
			"**Options:**\n",
			"- --help\tShow this screen.\n",
			"- --port N\tPort number\n")
	}
	flag.Parse()
}
