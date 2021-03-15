package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	run()
}

func ScanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	IP := "scanme.nmap.org"
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)

	if err != nil {
		return
	}
	fmt.Printf("[+] Connection established...PORT %v - IP: %v\n", port, connection.RemoteAddr().String())
	connection.Close()
}
func run() {

	// waitgroups help us keep the main function from
	//  closing before our go routines have finihsed
	var wg sync.WaitGroup
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go ScanPort(i, &wg)
	}

	// waits until wg counter is at zero
	wg.Wait()
}
