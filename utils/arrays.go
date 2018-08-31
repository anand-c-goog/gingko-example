// Package utils contains common utility functions.
package utils

// ContainsString returns true if arr contains the value s.
func ContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}
