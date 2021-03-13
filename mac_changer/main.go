package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func executeCommand(command string, args_arr []string) {
	args := args_arr

	cmd_obj := exec.Command(command, args...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr
	cmd_obj.Stdin = os.Stdin

	err := cmd_obj.Run()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}

func run() {

	iface := flag.String("iface", "eth0", "Interface that you want to chnage the MAC address")
	newMac := flag.String("newMac", "", "New MAC address")
	flag.Parse()

	executeCommand("sudo", []string{"ifconfig", *iface, "down"})
	executeCommand("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMac})
	executeCommand("sudo", []string{"ifconfig", *iface, "up"})
}
