package main

import (
	"log"
	"github.com/etc-bolt/skidder/input"
	"github.com/etc-bolt/skidder/parser"
	"github.com/etc-bolt/skidder/output"
)

func main() {

	raw := make(chan string)
	parsed := make(chan string)

	log.Print("bootstraping skidder log mangement system.")

	go input.LoadPlugins(raw)
	go parser.LoadPlugins(raw, parsed)

	output.LoadPlugins(parsed)

}
