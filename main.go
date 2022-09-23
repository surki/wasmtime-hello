package main

import (
	"fmt"
	"os"

	"github.com/bytecodealliance/wasmtime-go"
)

func main() {
	engine := wasmtime.NewEngine()

	wasm, err := os.ReadFile("./rust-wasm-hello-world/target/wasm32-wasi/debug/hello-world.wasm")
	check(err)

	module, err := wasmtime.NewModule(engine, wasm)
	check(err)

	linker := wasmtime.NewLinker(engine)
	err = linker.DefineWasi()
	check(err)

	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritStdin()
	wasiConfig.InheritStdout()
	wasiConfig.InheritStderr()

	store := wasmtime.NewStore(engine)
	store.SetWasi(wasiConfig)

	err = linker.DefineFunc(store, "env", "go_hello", func() {
		fmt.Println("Hello from Go!")
	})
	check(err)

	instance, err := linker.Instantiate(store, module)
	check(err)

	// Call Rust function
	rh := instance.GetFunc(store, "rust_hello")
	_, err = rh.Call(store)
	check(err)

	// Call Rust main function
	main := instance.GetFunc(store, "_start")
	_, err = main.Call(store)
	check(err)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
