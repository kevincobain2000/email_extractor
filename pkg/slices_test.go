package pkg

import (
	"fmt"
	"testing"
)

func TestIsEqualSlice(t *testing.T) {
	tests := []struct {
		a        []string
		b        []string
		expected bool
	}{
		{[]string{"a", "b"}, []string{"a", "b"}, true},
		{[]string{"a", "b"}, []string{"b", "a"}, false},
		{[]string{"a"}, []string{"a", "b"}, false},
		{[]string{}, []string{}, true},
		{[]string{"a"}, []string{}, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v vs %v", test.a, test.b), func(t *testing.T) {
			if got := IsEqualSlice(test.a, test.b); got != test.expected {
				t.Errorf("IsEqualSlice(%v, %v) = %v, want %v", test.a, test.b, got, test.expected)
			}
		})
	}
}

func TestUniqueStrings(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "a", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			if got := UniqueStrings(test.input); !IsEqualSlice(got, test.expected) {
				t.Errorf("UniqueStrings(%v) = %v, want %v", test.input, got, test.expected)
			}
		})
	}
}

func TestStringInSlice(t *testing.T) {
	tests := []struct {
		str      string
		list     []string
		expected bool
	}{
		{"a", []string{"a", "b", "c"}, true},
		{"d", []string{"a", "b", "c"}, false},
		{"a", []string{}, false},
		{"", []string{"a", "b", "c"}, false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s in %v", test.str, test.list), func(t *testing.T) {
			if got := StringInSlice(test.str, test.list); got != test.expected {
				t.Errorf("StringInSlice(%s, %v) = %v, want %v", test.str, test.list, got, test.expected)
			}
		})
	}
}
