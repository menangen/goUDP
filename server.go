package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const bufferSize = 8
const address = ":1234"

func main() {
	var gracefulStop = make(chan os.Signal)

	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	// listen to incoming udp packets
	serverSocket, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	udpConnection, err := net.ListenUDP("udp4", serverSocket)
	if err != nil {
		log.Fatal(err)
	}
	defer udpConnection.Close()

	log.Print("UDP server started on ", address)

	go func() {
		_ = <-gracefulStop
		os.Exit(0)
	}()

	for {
		buf := make([]byte, bufferSize)
		n, addr, udpConnectionErr := udpConnection.ReadFromUDP(buf)


		if udpConnectionErr != nil {
			log.Fatal("UDP buffer error ", udpConnectionErr)
		} else {
			go serve(udpConnection, addr, buf[:n])
		}
	}

}

func serve(udpConnection net.PacketConn, addr net.Addr, buf []byte) {
	log.Print("UDP from ", addr, buf)
	log.Printf("%d byte VALUE: [ %s ]", len(buf), string(buf))

	if _, err := udpConnection.WriteTo(buf, addr); err != nil {
		log.Fatal(err)
	}
}