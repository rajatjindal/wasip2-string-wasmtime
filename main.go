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

	// following block does not print anything for value
	// hello: func() -> result<string, error>;
	{
		hello := environment.Hello()
		if hello.IsErr() {
			fmt.Println("hello returned error ", hello.Err())
			return
		}

		value := hello.OK()
		fmt.Printf("hello value is %#v\n", *value)
	}

	// justhello: func() -> string;
	{
		justhello := environment.Justhello()
		fmt.Println("just hello value is ", justhello)
	}
}
