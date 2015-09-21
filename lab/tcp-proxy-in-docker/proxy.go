package main

import (
	"log"
	"net"
)
import "github.com/docker/docker/pkg/proxy"

func main() {
	newAddr := &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080}
	originAddr := &net.TCPAddr{IP: net.IPv4(125, 221, 232, 253), Port: 80}
	proxy, err := proxy.NewProxy(newAddr, originAddr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("run", originAddr, newAddr)
	proxy.Run()
}
