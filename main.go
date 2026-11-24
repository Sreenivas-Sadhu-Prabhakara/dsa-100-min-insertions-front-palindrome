package main

import "fmt"

// MinInsertions returns how many characters must be prepended to s to make
// it a palindrome: len(s) minus the longest palindromic prefix, found with
// one KMP failure-function pass over s + sep + reverse(s).
func MinInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}
	rev := reverse(s)
	t := s + "\x00" + rev
	lps := make([]int, len(t))
	k := 0
	for i := 1; i < len(t); i++ {
		for k > 0 && t[i] != t[k] {
			k = lps[k-1]
		}
		if t[i] == t[k] {
			k++
		}
		lps[i] = k
	}
	// lps of the full string = longest prefix of s that is a suffix of
	// reverse(s), i.e. the longest palindromic prefix of s.
	return len(s) - lps[len(t)-1]
}

func reverse(s string) string {
	b := []byte(s)
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}

func main() {
	fmt.Println(MinInsertions("abcd"))     // 3: prepend dcb
	fmt.Println(MinInsertions("aacecaaa")) // 2: prepend aa
}
