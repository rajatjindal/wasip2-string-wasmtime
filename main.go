package main

import (
	"fmt"

	"github.com/rajatjindal/wasip2-string-wasmtime/internal/wasi/cli/v0.2.0/environment"
)

func main() {
	fmt.Println("hello world")
	result := environment.GetArguments()
	fmt.Println("slice ", result.Slice())

	fmt.Println("slice ", environment.InitialCWD())
	fmt.Println("slice ", environment.Hello())
}
