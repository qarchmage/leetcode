package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("192.12.254.1"))
	fmt.Println(validIPAddress("2001:16b8:68c6:4700:59c2:5802:d4ec:644a"))
	fmt.Println(validIPAddress("1e1.2.3.4"))
	fmt.Println(validIPAddress("01.01.01.01"))
}

func validIPAddress(IP string) string {
	if validIPv4Address(IP) {
		return "IPv4"
	}
	if validIPv6Address(IP) {
		return "IPv6"
	}
	return "Neither"
}

func validIPv4Address(IP string) bool {
	numbers := strings.Split(IP, ".")
	if len(numbers) != 4 {
		return false
	}
	for _, num := range numbers {
		if v, err := strconv.Atoi(num); err != nil || v < 0 || v > 255 {
			return false
		} else if strconv.Itoa(v) != num {
			return false
		}
	}
	return true
}

func validIPv6Address(IP string) bool {
	numbers := strings.Split(IP, ":")
	if len(numbers) != 8 {
		return false
	}
	re := regexp.MustCompile(`^([0-9]|[a-f]|[A-F])+$`)
	for _, num := range numbers {
		if len(num) == 0 || len(num) > 4 {
			return false
		}
		if !re.MatchString(num) {
			return false
		}
	}
	return true
}
