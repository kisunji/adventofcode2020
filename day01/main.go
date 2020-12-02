package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
	r := strings.NewReader(input)
	ints := extractIntArr(r)
	for i := 0; i < len(ints); i++ {
		a := ints[i]
		for j := i + 1; j < len(ints); j++ {
			b := ints[j]
			for k := j + 1; k < len(ints); k++ {
				c := ints[k]
				if a+b+c == 2020 {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}

// func main() {
// 	r := strings.NewReader(input)
// 	ints := extractIntArr(r)
// 	for i := 0; i < len(ints); i++ {
// 		a := ints[i]
// 		for j := i + 1; j < len(ints); j++ {
// 			b := ints[j]
// 			if a + b == 2020 {
// 				fmt.Println(a*b)
// 				return
// 			}
// 		}
// 	}
// }
func extractIntArr(reader io.Reader) []int {
	var ints []int
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
