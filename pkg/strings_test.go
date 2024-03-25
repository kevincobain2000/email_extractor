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
		{"Scheme Mismatch", "http://example.com", "https://example.com", false},
		{"Host Mismatch", "http://example.com", "http://example.org", false},
		{"Host as Subdomain", "http://example.com", "http://sub.example.com", true},
		{"Identical Hosts", "http://example.com", "http://example.com", true},
		{"Path Mismatch", "http://example.com/path", "http://example.com/otherpath", false},
		{"Path Prefix", "http://example.com/path", "http://example.com/path/sub", true},
		{"Query Parameter Mismatch", "http://example.com?param=1", "http://example.com?param=2", false},
		{"Query Parameter Equality", "http://example.com?param=1", "http://example.com?param=1", true},
		{"Additional Query Parameters in URL2", "http://example.com?param=1", "http://example.com?param=1&extra=2", true},
		{"Empty URLs", "", "", true},
		{"Invalid URL1", "http://%41:8080/", "http://example.com", false},
		{"Invalid URL2", "http://example.com", "http://%41:8080/", false},
		{"All Components Together", "http://example.com/path?param=1", "http://sub.example.com/path/subpath?param=1&extra=2", true},
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
