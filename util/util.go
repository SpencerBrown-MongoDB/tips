package util

import "strings"

// NilString takes a pointer to a string and returns the string or "" if the pointer is nil
func NilString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// AddrOfString returns the address of a string.
// Useful to get the address of a string constant.
func AddrOfString(s string) *string {
	return &s
}

// AddrOfInt returns the address of an int.
// Useful to get the address of an int constant.
func AddrOfInt(i int) *int {
	return &i
}

// AddrOfBool returns the address of an bool.
// Useful to get the address of an bool constant.
func AddrOfBool(b bool) *bool {
	return &b
}

// DerefBool is a helper function for templates that dereferences a pointer to bool
func DerefBool(x *bool) bool { return *x }

// ConvertSizeToMB converts a size in bytes to megabytes.
// It is used as a function in templates.
func ConvertSizeToMB(size int64) int {
	return int(size / 1024 / 1024)
}

// ConvertSliceToStringList converts a slice of strings to a string listing the strings.
// It is used as a function in templates.
func ConvertSliceToStringList(theList []string) string {
	return strings.Join(theList, ", ")
}
