package main

import (
	"github.com/johntdyer/anybar-go/Godeps/_workspace/src/github.com/wsxiaoys/terminal"
	"github.com/johntdyer/anybar-go/Godeps/_workspace/src/github.com/wsxiaoys/terminal/color"
)

func main() {
	terminal.Stdout.Color("y").
		Print("Hello world").Nl().
		Reset().
		Colorf("@{kW}Hello world\n")

	color.Print("@rHello world")
}
