#!/usr/bin/env python

import subprocess
import optparse

parser = optparse.OptionParser()
parser.add_option("-i", "--interface", dest="interface",
                  help="Interface to change its MAC address")
parser.add_option("-m", "--new_nac", dest="new_mac",
                  help="New MAC address")

(options, arguments) = parser.parse_args()

interface = options.interface
new_mac = options.new_mac
print(f"[+] Changing MAC address for {interface} to {new_mac}")

# subprocess.call(f"ifconfig {interface} down", shell=True)
# subprocess.call(f"ifconfig {interface} hw ether {new_mac}", shell=True)
# subprocess.call(f"ifconfig {interface} up", shell=True)

# More secure version of the above.
# Stop highjacking commands
subprocess.call(["ifconfig", interface, "down"])
subprocess.call(["ifconfig", interface, "hw", "ether", new_mac])
subprocess.call(["ifconfig", "up"])
