package arp

import (
	"bufio"
	"os"
	"strings"
)

const (
	iPAddr int = iota
	hWType
	flags
	hWAddr
	mask
	device
)

// Value item in ARP table cache
type Value struct {
	Iface   string
	MacAddr string
	IPAddr  string
}

func newArpItem(iface string, mac string, ip string) *Value {
	return &Value{Iface: iface, MacAddr: mac, IPAddr: ip}
}

// List list of rows from ARP cache
func List() (table []Value) {
	f, err := os.Open("/proc/net/arp")

	if err != nil {
		return nil
	}

	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan() // skip the field descriptions

	for s.Scan() {
		line := s.Text()
		fields := strings.Fields(line)
		table = append(table, *newArpItem(fields[device], fields[hWAddr], fields[iPAddr]))
	}

	return table
}

// IPLookup find IP address by MAC addr.
func IPLookup(mac string) string {
	for _, item := range List() {
		if item.MacAddr == mac {
			return item.IPAddr
		}
	}
	return ""
}
