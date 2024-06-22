package helpers

import "github.com/pynezz/pynezzentials/ansi"

var Warning = func(warning string) {
	r, g, b, _ := ansi.HexToRGB("#e64553") // catppuccin latte maroon
	ansi.PrintBold(ansi.HexColor256(r, g, b, warning+"\n"))
}
