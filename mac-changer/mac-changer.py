#!/usr/bin/env python

import subprocess
import optparse


def change_mac(interface, new_mac):

    print(f"[+] Changing MAC address for {interface} to {new_mac}")
    subprocess.call(["ifconfig", interface, "down"])
    subprocess.call(["ifconfig", interface, "hw", "ether", new_mac])
    subprocess.call(["ifconfig", "up"])


parser = optparse.OptionParser()
parser.add_option("-i", "--interface", dest="interface",
                  help="Interface to change its MAC address")
parser.add_option("-m", "--new_nac", dest="new_mac",
                  help="New MAC address")

(options, arguments) = parser.parse_args()

change_mac(options.interface, options.new_mac)
