package main

// #cgo CFLAGS:   -I. -O3 -DNDEBUG -std=c17 -fPIC -pthread -mavx -mavx2 -mfma -mf16c -msse3
// #cgo CXXFLAGS: -O3 -DNDEBUG -std=c++17 -fPIC -pthread -I.
// #include "chat.h"
import "C"
import (
	"flag"
)

func main() {

	help := flag.Bool("h", false, "show this help message and exit")
	interactive := flag.Bool("i", false, "run in interactive mode")
	interactiveStart := flag.Bool("interactive-start", false, "run in interactive mode and poll user input at startup")
	reversePrompt := flag.String("r", "", "in interactive mode, poll user input upon seeing PROMPT")
	color := flag.Bool("color", false, "colorise output to distinguish prompt and user input from generations")
	seed := flag.Int("s", -1, "RNG seed (default: -1)")
	threads := flag.Int("t", 4, "number of threads to use during computation (default: 4)")
	prompt := flag.String("p", "", "prompt to start generation with (default: random)")

	flag.Parse()

	var helpAsInt C.int = 0
	if *help {
		helpAsInt = 1
	}

	var interactiveAsInt C.int = 0
	if *interactive {
		interactiveAsInt = 1
	}

	var interactiveStartAsInt C.int = 0
	if *interactiveStart {
		interactiveStartAsInt = 1
	}

	var colorAsInt C.int = 0
	if *color {
		colorAsInt = 1
	}

	C.alpaca(
		helpAsInt,
		interactiveAsInt,
		interactiveStartAsInt,
		C.CString(*reversePrompt),
		colorAsInt,
		C.int(*seed),
		C.int(*threads),
		C.CString(*prompt),
	)

}
