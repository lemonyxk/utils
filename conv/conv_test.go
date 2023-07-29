/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-29 20:17
**/

package conv

import (
	"bytes"
	"testing"
)

func TestItoa(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{123, "123"},
		{-456, "-456"},
		{0, "0"},
	}

	for _, tc := range testCases {
		result := Itoa(tc.input)
		if result != tc.expected {
			t.Errorf("Itoa(%d) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestAtoi(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"123", 123},
		{"-456", -456},
		{"0", 0},
	}

	for _, tc := range testCases {
		result := Atoi(tc.input)
		if result != tc.expected {
			t.Errorf("Atoi(%s) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

func TestStringToBytes(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{"hello", []byte{104, 101, 108, 108, 111}},
		{"world", []byte{119, 111, 114, 108, 100}},
		{"", []byte{}},
	}

	for _, tc := range testCases {
		result := StringToBytes(tc.input)
		if !bytes.Equal(result, tc.expected) {
			t.Errorf("StringToBytes(%s) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

func TestBytesToString(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected string
	}{
		{[]byte{104, 101, 108, 108, 111}, "hello"},
		{[]byte{119, 111, 114, 108, 100}, "world"},
		{[]byte{}, ""},
	}

	for _, tc := range testCases {
		result := BytesToString(tc.input)
		if result != tc.expected {
			t.Errorf("BytesToString(%v) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestFloat64ToString(t *testing.T) {
	testCases := []struct {
		input    float64
		expected string
	}{
		{123.456, "123.456"},
		{-456.789, "-456.789"},
		{0, "0"},
	}

	for _, tc := range testCases {
		result := Float64ToString(tc.input)
		if result != tc.expected {
			t.Errorf("Float64ToString(%f) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestFloat32ToString(t *testing.T) {
	testCases := []struct {
		input    float64
		expected string
	}{
		{123.456, "123.456"},
		{-456.789, "-456.789"},
		{0, "0"},
	}

	for _, tc := range testCases {
		result := Float32ToString(tc.input)
		if result != tc.expected {
			t.Errorf("Float32ToString(%f) = %s; want %s", tc.input, result, tc.expected)
		}
	}
}

func TestStringToFloat64(t *testing.T) {
	testCases := []struct {
		input    string
		expected float64
	}{
		{"123.456", 123.456},
		{"-456.789", -456.789},
		{"0", 0},
	}

	for _, tc := range testCases {
		result := StringToFloat64(tc.input)
		if result != tc.expected {
			t.Errorf("StringToFloat64(%s) = %f; want %f", tc.input, result, tc.expected)
		}
	}
}

func TestStringToFloat32(t *testing.T) {
	testCases := []struct {
		input    string
		expected float64
	}{
		{"123.456", 123.456},
		{"-456.789", -456.789},
		{"0", 0},
	}

	for _, tc := range testCases {
		result := StringToFloat32(tc.input)
		if int(result*1000) != int(tc.expected*1000) {
			t.Errorf("StringToFloat32(%s) = %f; want %f", tc.input, result, tc.expected)
		}
	}
}
