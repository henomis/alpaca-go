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
	promptFile := flag.String("f", "", "prompt file to start generation.")
	nPredict := flag.Int("n", 128, "number of tokens to predict (default: 128)")
	topK := flag.Int("top_k", 40, "top-k sampling (default: 40)")
	topP := flag.Float64("top_p", 0.9, "top-p sampling (default: 0.9)")
	repeatLastN := flag.Int("repeat_last_n", 64, "last n tokens to consider for penalize (default: 64)")
	repeatPenalty := flag.Float64("repeat_penalty", 1.3, "penalize repeat sequence of tokens (default: 1.3)")
	promptSize := flag.Int("c", 2048, "size of the prompt context (default: 2048)")
	temperature := flag.Float64("temp", 0.1, "temperature (default: 0.1)")
	batchSize := flag.Int("b", 8, "batch size for prompt processing (default: 8)")
	modelFile := flag.String("m", "ggml-alpaca-7b-q4.bin", "model path (default: ggml-alpaca-7b-q4.bin)")

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
		C.CString(*promptFile),
		C.int(*nPredict),
		C.int(*topK),
		C.double(*topP),
		C.int(*repeatLastN),
		C.double(*repeatPenalty),
		C.int(*promptSize),
		C.double(*temperature),
		C.int(*batchSize),
		C.CString(*modelFile),
	)

}
