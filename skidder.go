package main

import (
	"log"
	"github.com/etc-bolt/skidder/input"
	"github.com/etc-bolt/skidder/parser"
	"github.com/etc-bolt/skidder/output"
)

func main() {

	c := make(chan string)

	log.Print("bootstraping skidder log mangement system.")

	go input.LoadPlugins(c)
	go parser.LoadPlugins(c)

	output.LoadPlugins(c)

}
