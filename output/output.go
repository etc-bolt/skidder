package output

import "log"

func LoadPlugins(c chan string) {
	for {
		select {
		case msg := <-c:
			log.Println(msg)
		}
	}
}
