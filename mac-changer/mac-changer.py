#!/usr/bin/env python

import subprocess

interface = input("What is the interface?")
new_mac = input("What is the new MAC")

print(f"[+] Changing MAC address for {interface} to {new_mac}")

# subprocess.call(f"ifconfig {interface} down", shell=True)
# subprocess.call(f"ifconfig {interface} hw ether {new_mac}", shell=True)
# subprocess.call(f"ifconfig {interface} up", shell=True)

# More secure version of the above.
# Stop highjacking commands
subprocess.call(["ifconfig", interface, "down"])
subprocess.call(["ifconfig", interface, "hw", "ether", new_mac])
subprocess.call(["ifconfig", "up"])
