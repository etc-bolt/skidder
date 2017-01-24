package parser

import (
	"regexp"
	"log"
)

func LoadPlugins(c chan string, parsed chan string) {
	r, err := regexp.Compile(`([A-Za-z]+\s+\d+\s+\d+\:\d+\:\d+) ([^ ]+) ([^\:]+)`)

	if err != nil {
		log.Println(err.Error())
	}

	for {
		select {
		case msg := <-c:
			m := r.FindAllString(msg, -1)
			log.Print(m)
			parsed <- msg
		}
	}

}
