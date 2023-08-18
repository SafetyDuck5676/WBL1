package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(size int) string {
	var builder strings.Builder
	builder.WriteString(strings.Repeat("a", size))
	return builder.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	var builder strings.Builder
	builder.WriteString(v[:100])
	justString = builder.String()
}

func main() {
	someFunc()
	fmt.Println(justString)
}
