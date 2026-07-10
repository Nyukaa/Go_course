package stringutil

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        input    string
        expected string
    }{
        {"hello", "olleh"},
        {"Go", "oG"},
        {"", ""},
    }

    for _, c := range cases {
        result := Reverse(c.input)
        if result != c.expected {
            t.Errorf("Reverse(%q) = %q; want %q", c.input, result, c.expected)
        }
    }
}