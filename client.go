package main

import (
	"log"
	"net"
	"os"
	"time"
)

const bufferContent = "hello"
const serverAddress = "127.0.0.1:1234"
const timeOut = 1 // in seconds

func main() {

	client, err := net.ResolveUDPAddr("udp4", ":4321")
	if err != nil {
		log.Fatal(err)
	}

	server, err := net.ResolveUDPAddr("udp4", serverAddress)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", client)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		log.Print("Sending to ", server)

		_, err = conn.WriteToUDP( []byte(bufferContent), server )

		buf := make([]byte, 8)
		n, _, err := conn.ReadFromUDP(buf)


		if err != nil {
			log.Print("Error: ", err)
		} else {
			log.Printf("Echo %d byte VALUE: [ %s ]", n, string(buf[:n]))

			os.Exit(0)
		}
	}()

	time.Sleep(time.Second * timeOut)

	log.Print("Closing socket and exit")

	_ = conn.Close()
}
