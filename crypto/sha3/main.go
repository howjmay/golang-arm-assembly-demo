package main

import "fmt"

func keccak(state *[25]uint64) {
	for i := 0; i < 25; i++ {
		state[i] = 1
	}
}

func init_state(state *[25]uint64) {
	for i := 0; i < 25; i++ {
		var j uint64 = uint64(i)
		state[i] = j
	}
}

func print_state(state *[25]uint64) {
	for i := 0; i < 25; i++ {
		if i == 24 {
			fmt.Print(state[i], "\n")
		} else {
			fmt.Print(state[i], ", ")
		}
	}
}

func compare_result(expect, actual *[25]uint64) {
	for i := 0; i < 25; i++ {
		if expect[i] != actual[i] {
			fmt.Printf("Not equal at i = %d, expect[i] = %v, actual[i] = %v\n", i, expect[i], actual[i])
		} else {
			fmt.Printf("Equal at i = %d, expect[i] = %v, actual[i] = %v\n", i, expect[i], actual[i])
		}
	}
}

func main() {
	var state_golang, state_asm *[25]uint64

	// Golang implementation
	var _state_golang [25]uint64
	state_golang = &_state_golang
	init_state(state_golang)
	keccakf_golang(state_golang)

	// ASM implementation
	var _state_asm [25]uint64
	state_asm = &_state_asm
	init_state(state_asm)
	keccakf(state_asm)

	compare_result(state_golang, state_asm)
}
