/**
* @program: utils
*
* @description:
*
* @author: lemo
*
* @create: 2023-07-29 20:15
**/

package args

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	testCases := []struct {
		input    []string
		flag     []string
		expected string
	}{
		{[]string{"--name", "John", "--age", "30"}, []string{"--name"}, "John"},
		{[]string{"--name", "John", "--age", "30"}, []string{"--age"}, "30"},
		{[]string{"--name", "John", "--age", "30"}, []string{"--gender"}, ""},
	}

	for _, tc := range testCases {
		os.Args = append([]string{"test"}, tc.input...)
		result := Get(tc.flag...)
		if result != tc.expected {
			t.Errorf("Get(%v, %v) = %s; want %s", tc.input, tc.flag, result, tc.expected)
		}
	}
}

func TestHas(t *testing.T) {
	testCases := []struct {
		input    []string
		flag     []string
		expected bool
	}{
		{[]string{"--name", "John", "--age", "30"}, []string{"--name"}, true},
		{[]string{"--name", "John", "--age", "30"}, []string{"--age"}, true},
		{[]string{"--name", "John", "--age", "30"}, []string{"--gender"}, false},
	}

	for _, tc := range testCases {
		os.Args = append([]string{"test"}, tc.input...)
		result := Has(tc.flag...)
		if result != tc.expected {
			t.Errorf("Has(%v, %v) = %t; want %t", tc.input, tc.flag, result, tc.expected)
		}
	}
}
