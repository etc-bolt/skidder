package tcp

import (
	"bufio"
	"log"
	"net"
)

const tcp_port string = "0.0.0.0:5170"

func ListenAndServe(ch chan string) {
	l, err := net.Listen("tcp", tcp_port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("%v", err.Error())
		}
		go handleConnection(conn, ch)
	}
}


func handleConnection(c net.Conn, ch chan string){
	defer c.Close()
	message, err := bufio.NewReader(c).ReadString('\n')

	if err != nil {
		log.Println(err)
	}

	ch <- message;
}