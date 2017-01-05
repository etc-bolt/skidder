package main

import (
	"github.com/etc-bolt/skidder/input/tcp"
	"log"
	"net"
	"regexp"
)

func main() {

	c := make(chan string)

	r, err := regexp.Compile(`([A-Za-z]+\s+\d+\s+\d+\:\d+\:\d+) ([^ ]+) ([^\:]+)`)

	if err != nil {
		log.Println(err.Error())
	}

	log.Print("bootstraping skidder log mangement system.")

	go tcp.ListenAndServe(c)

	for {
		select {
		case msg := <-c:
			matches := r.FindAllString(msg, -1)
			log.Println(matches)
		}
	}
}
