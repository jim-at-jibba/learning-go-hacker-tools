package main

import (
	"fmt"
	"net"
)

func main() {
	run()
}

func run() {
	IP := "scanme.nmap.org"

	for i := 1; i < 100; i++ {
		address := fmt.Sprintf(IP+":%d", i)
		connection, err := net.Dial("tcp", address)

		if err == nil {
			fmt.Printf("[+] Connection established...PORT %v - IP: %v\n", i, connection.RemoteAddr().String())
		}
	}

}
