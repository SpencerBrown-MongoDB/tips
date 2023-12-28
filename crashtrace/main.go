package main

import "runtime/debug"

func main() {
	defer printTrace()
	crash()
}

func crash() {
	var x *int
	var y int = *x
	_ = y
}

func printTrace() {
	if r := recover(); r != nil {
		debug.PrintStack()
	}
}