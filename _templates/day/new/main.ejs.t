---
to: day<%=day%>/main.go
---
package main

import (
	"flag"
)

var (
	useTestInputs = flag.Bool("test", false, "Should we use the test inputs")
)

func init() {
	flag.Parse()
}

func main() {
}