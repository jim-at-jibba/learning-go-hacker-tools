package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//1. Get arguments
//2. change_mac
//3. Get current Mac

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup err: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("[+] Running the mac changer")

	hardwareInterface := flag.String("hardwareInterface", "", "Interface to change its MAC address")
	newMac := flag.String("newMac", "", "Interface to change its MAC address")
	flag.Parse()
	if len(strings.TrimSpace(*hardwareInterface)) == 0 {
		return errors.New("[-] Please specify an interface")
	}

	if len(strings.TrimSpace(*newMac)) == 0 {
		return errors.New("[-] Please specify a new MAC")
	}

	fmt.Printf("[+] Changing %v to new MAC %v\n", *hardwareInterface, *newMac)

	currentMac := getCurrentMac(*hardwareInterface)
	fmt.Printf("[+] Current MAC %v: ", currentMac)
	return nil
}

func getCurrentMac(hardwareInterface string) string {
	fmt.Printf("[+] Getting current MAC address %v\n", hardwareInterface)
	ifconfigOutput, err := exec.Command("ifconfig", hardwareInterface).Output()

	if err != nil {
		log.Fatalln("[-] Could not read MAC address")
	}

	r, err := regexp.Compile("([0-9a-fA-F]{2}[:]){5}([0-9a-fA-F]{2})")

	if err != nil {
		log.Fatal(err)
	}

	match := r.FindString(string(ifconfigOutput))

	if len(match) > 0 {
		return match
	}

	return ""
}
