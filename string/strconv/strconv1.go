package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IPAddr IP地址
type IPAddr [4]byte

func (p IPAddr) String() string {
	var ipParts []string
	for _, item := range p {
		ipParts = append(ipParts, strconv.Itoa(int(item)))
	}

	return strings.Join(ipParts, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}