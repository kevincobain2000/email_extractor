package pkg

import (
	"reflect"
	"testing"
)

func TestByteSliceToString(t *testing.T) {
	tests := []struct {
		bytes    []byte
		expected string
	}{
		{[]byte("hello"), "hello"},
		{[]byte(""), ""},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			if got := ByteSliceToString(test.bytes); got != test.expected {
				t.Errorf("ByteSliceToString(%v) = %s, want %s", test.bytes, got, test.expected)
			}
		})
	}
}

func TestStringToByteSlice(t *testing.T) {
	tests := []struct {
		str      string
		expected []byte
	}{
		{"hello", []byte("hello")},
		{"", []byte("")},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			if got := StringToByteSlice(test.str); !reflect.DeepEqual(got, test.expected) {
				t.Errorf("StringToByteSlice(%s) = %v, want %v", test.str, got, test.expected)
			}
		})
	}
}

// TestIsSameDomain - Testing various scenarios for IsSameDomain function.
func TestIsSameDomain(t *testing.T) {
	tests := []struct {
		name string
		url1 string
		url2 string
		want bool
	}{
		{"Scheme Mismatch", "http://example.com", "https://example.com", true},
		{"Host Mismatch", "http://example.com", "http://example.org", false},
		{"Host as Subdomain", "http://example.com", "http://sub.example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSameDomain(tt.url1, tt.url2)
			if got != tt.want {
				t.Errorf("IsSameDomain(%q, %q) = %v, want %v", tt.url1, tt.url2, got, tt.want)
			}
		})
	}
}
