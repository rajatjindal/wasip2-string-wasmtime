reproducing issue with wasm-tools-go

- use tinygo version with support for wasip2
- use wasmtime from fork: 

^^ the above fork only adds following two functions to environment.wit

```
  justhello: func() -> string;
  hello: func() -> result<string, error>;

  variant error {
    /// The provided variable name is invalid.
    invalid-name(string),
    /// The provided variable is undefined.
    undefined(string),
    /// A variables provider specific error has occurred.
    provider(string),
    /// Some implementation-specific error has occurred.
    other(string),
  }
```

- build using `./build.sh`
- run using `wasmtime run main.wasm`

Notice following output:

```
$ wasmtime run main.wasm
hello world
slice  [main.wasm]
slice  {{true whatever}}

## This is the output from the function that is trying to replicate what is reported in original issue
hello value is "\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"

## this is from a function that just returns a string
just hello value is  justhello from wasmtime

```