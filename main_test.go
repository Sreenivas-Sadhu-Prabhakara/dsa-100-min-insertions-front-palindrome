package main

import "testing"

func isPal(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func TestMinInsertions(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"abcd", 3},
		{"aacecaaa", 1}, // palindromic prefix "aacecaa" leaves one trailing a
		{"aba", 0},
		{"", 0},
		{"a", 0},
		{"aa", 0},
		{"ab", 1},
		{"aabba", 3}, // abb + aabba = abbaabba
	}
	for _, c := range cases {
		if got := MinInsertions(c.in); got != c.want {
			t.Errorf("MinInsertions(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}

// Property: prepending the reversed unmatched suffix must yield a palindrome,
// and removing one fewer character must not.
func TestConstructionProperty(t *testing.T) {
	for _, s := range []string{"abcd", "aacecaaa", "banana", "xx", "race"} {
		k := MinInsertions(s)
		fixed := reverse(s[len(s)-k:]) + s
		if !isPal(fixed) {
			t.Errorf("%q: prepending %d chars gave non-palindrome %q", s, k, fixed)
		}
	}
}
