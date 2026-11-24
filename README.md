# Minimum Insertions at Front to Make Palindrome

> Striver SDE Sheet 100 · Day 16 String II — the answer is the length minus the longest palindromic prefix, and KMP's failure function finds that prefix in one pass. Go, stdlib only, tested.

## Problem

How few characters must be *prepended* to a string to make it a palindrome? `"abcd"` needs 3 (`dcb·abcd`); `"aacecaaa"` needs 2.

## Approach

Whatever gets prepended mirrors the string's tail, so the characters that *don't* need help are exactly the longest prefix that is already a palindrome — answer = n − (longest palindromic prefix). Now the reduction that makes this a string-matching problem: a prefix of s is a palindrome iff it is also a suffix of reverse(s). "Longest prefix of A that is a suffix of B" is literally what KMP's LPS value computes on `A + separator + B` — one failure-function pass over `s + '\0' + reverse(s)` and the last entry is the palindromic prefix length. The separator prevents the match from spilling across the join. A property test prepends the computed characters and asserts the result really is a palindrome.

**Complexity:** O(n) time and space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
