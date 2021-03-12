#!/usr/bin/env python

import subprocess
import optparse
import re


def get_arguments():
    parser = optparse.OptionParser()
    parser.add_option("-i", "--interface", dest="interface",
                      help="Interface to change its MAC address")
    parser.add_option("-m", "--new_nac", dest="new_mac",
                      help="New MAC address")
    (options, arguments) = parser.parse_args()
    if not options.interface:
        parser.error("[-] Please specify an interface")
    elif not options.new_mac:
        parser.error("[-] Please specify an interface")
    return options


def change_mac(interface, new_mac):
    print(f"[+] Changing MAC address for {interface} to {new_mac}")
    subprocess.call(["ifconfig", interface, "down"])
    subprocess.call(["ifconfig", interface, "hw", "ether", new_mac])
    subprocess.call(["ifconfig", interface, "up"])


def get_current_mac(interface):

    ifconfig_result = subprocess.check_output(['ifconfig', interface])
    mac_address_search_result = re.search(r"([0-9a-fA-F]{2}[:]){5}([0-9a-fA-F]{2})", str(ifconfig_result))

    if mac_address_search_result:
        return mac_address_search_result.group(0)
    else:
        print("[-] Could not read MAC address")


options = get_arguments()

current_mac = get_current_mac(options.interface)
print(f"Current Mac {str(current_mac)}")

# change_mac(options.interface, options.new_mac)
current_mac = get_current_mac(options.interface)
if current_mac == options.new_mac:
    print(f"[+] MAC address successfully changed to {options.new_mac}")
else:
    print("[-] MAC address not changed")
