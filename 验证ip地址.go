package main

import (
	"strconv"
	"strings"
)

func main() {
	validIPAddress("172.16.254.")
}

func validIPAddress(queryIP string) string {
	if pts := strings.Split(queryIP, "."); len(pts) == 4 {
		for _, n := range pts {
			if len(n) > 3 || len(n) == 0 {
				return "Neither"
			}
			for i := 0; i < len(n); i++ {
				if n[i] < '0' || n[i] > '9' || (n[0] == '0' && len(n) > 1) {
					return "Neither"
				}
			}
			if num, _ := strconv.Atoi(n); num > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	if pts := strings.Split(queryIP, ":"); len(pts) == 8 {
		for _, n := range pts {
			if len(n) > 4 || len(n) == 0 {
				return "Neither"
			}
			for i := 0; i < len(n); i++ {
				if (n[i] > 'f' || n[i] < 'a') && (n[i] > 'F' || n[i] < 'A') && (n[i] > '9' || n[i] < '0') {
					return "Neither"
				}
			}
		}
		return "IPv6"
	}
	return "Neither"
}
