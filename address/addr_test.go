/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-29 20:07
**/

package address

import (
	"net"
	"testing"
)

func TestIP2long(t *testing.T) {
	testCases := []struct {
		input    string
		expected uint32
	}{
		{"127.0.0.1", 2130706433},
		{"192.168.0.1", 3232235521},
		{"10.0.0.1", 167772161},
	}

	for _, tc := range testCases {
		result := IP2long(tc.input)
		if result != tc.expected {
			t.Errorf("IP2long(%s) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

func TestIsLocalIP(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"192.168.0.1", true},
		{"8.8.8.8", false},
	}

	for _, tc := range testCases {
		result := IsLocalIP(tc.input)
		if result != tc.expected {
			t.Errorf("IsLocalIP(%s) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}

func TestIsLocalNet(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"127.0.0.1", true},
		{"192.168.0.1", true},
		{"8.8.8.8", false},
	}

	for _, tc := range testCases {
		ip := net.ParseIP(tc.input)
		result := IsLocalNet(ip)
		if result != tc.expected {
			t.Errorf("IsLocalNet(%s) = %t; want %t", tc.input, result, tc.expected)
		}
	}
}

func TestParse(t *testing.T) {
	testCases := []struct {
		input        string
		expectedIP   string
		expectedPort int
		expectedErr  error
	}{
		{"localhost", "localhost", 80, nil},
		{"127.0.0.1:8080", "127.0.0.1", 8080, nil},
		{"example.com:443", "example.com", 443, nil},
		{"invalid", "invalid", 80, nil},
	}

	for _, tc := range testCases {
		ip, port, err := Parse(tc.input)
		if ip != tc.expectedIP || port != tc.expectedPort || err != tc.expectedErr {
			t.Errorf("Parse(%s) = (%s, %d, %v); want (%s, %d, %v)", tc.input, ip, port, err, tc.expectedIP, tc.expectedPort, tc.expectedErr)
		}
	}
}
