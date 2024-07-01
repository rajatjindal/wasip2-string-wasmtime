// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminalstderr represents the imported interface "wasi:cli/terminal-stderr@0.2.0".
//
// An interface providing an optional `terminal-output` for stderr as a
// link-time authority.
package terminalstderr

import (
	terminaloutput "github.com/rajatjindal/wasip2-string-wasmtime/internal/wasi/cli/v0.2.0/terminal-output"
	"github.com/ydnar/wasm-tools-go/cm"
)

// GetTerminalStderr represents the imported function "get-terminal-stderr".
//
// If stderr is connected to a terminal, return a `terminal-output` handle
// allowing further interaction with it.
//
//	get-terminal-stderr: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStderr() (result cm.Option[terminaloutput.TerminalOutput]) {
	wasmimport_GetTerminalStderr(&result)
	return
}

//go:wasmimport wasi:cli/terminal-stderr@0.2.0 get-terminal-stderr
//go:noescape
func wasmimport_GetTerminalStderr(result *cm.Option[terminaloutput.TerminalOutput])
