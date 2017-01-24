package input

import "github.com/etc-bolt/skidder/input/tcp"

func LoadPlugins(c chan string) {
	go tcp.ListenAndServe(c)
}
