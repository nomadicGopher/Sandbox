package main

import (
	"fmt"
	"syscall/js"
)

func helloDom(this js.Value, args []js.Value) interface{} {
	fmt.Println("go via wasm: Hello DOM!")
	return "Hello World!"
}

func main() {
	js.Global().Set("helloDom", js.FuncOf(helloDom))
	select {}
}
