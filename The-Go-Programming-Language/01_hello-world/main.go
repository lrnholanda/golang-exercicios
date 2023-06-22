package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// var s, sep string
// for i := 1; i < len(os.Args); i++ {
// 	s += sep + os.Args[i]
// 	sep = " "
// }
// fmt.Println(s)
// s, sep := "", ""
// i := 0
// for index, arg := range os.Args[1:] {
// 	s += sep + arg
// 	sep = " "
// 	i += index
// }
// fmt.Println(s, i)

// A simpler and more efficient solution wou ld be to use the Join func tion fro m the strings package:

// fmt.Println(strings.Join(os.Args[1:], " "))
