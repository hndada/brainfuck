package main

import "fmt"

var memory = make(map[uint8]byte)

func Run(s string) {
	var ptr uint8
	i := 0
	for i < len(s) {
		c := s[i]
		// fmt.Println(i, ptr, memory[ptr])
		switch c {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			fmt.Printf("%c", memory[ptr])
			// fmt.Println()
		case ',':
			// Todo: need a test
			var b byte
			fmt.Scanf("%c", &b)
			memory[ptr] = b
		case '[':
			if memory[ptr] != 0 {
				break
			}
			for i < len(s) {
				if s[i] == ']' {
					break
				}
				i++
			}
		case ']':
			if memory[ptr] == 0 {
				break
			}
			for i >= 0 {
				if s[i] == '[' {
					break
				}
				i--
			}
		}
		i++
	}
}

func main() {
	const code = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`
	Run(code)
}
