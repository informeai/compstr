package main

import (
	"fmt"
	"os"
	"strings"
)

func mergeString(a1, a2 string) string {
	arr1 := []rune(a1)
	arr2 := []rune(a2)
	result := ""
	if len(arr1) > len(arr2) {
		for i := 0; i < len(arr2); i++ {
			if arr1[i] == arr2[i] {
				result += fmt.Sprintf("\033[0;32m%s\033[0m", string(arr1[i]))
			} else {
				result += fmt.Sprintf("\033[0;31m%s\033[0m", string(arr2[i]))
			}
		}
    
    result += fmt.Sprintf("%s", string(arr1[len(arr2):]))

	} else {
		for i := 0; i < len(arr1); i++ {
			if arr1[i] == arr2[i] {
				result += fmt.Sprintf("\033[0;32m%s\033[0m", string(arr1[i]))
			} else {
				result += fmt.Sprintf("\033[0;31m%s\033[0m", string(arr2[i]))
			}
		}
    result += fmt.Sprintf("%s", string(arr2[len(arr1):]))
	}
	return result
}

func printResult(comp int, a1, a2 string) {
	switch comp {
	case 1:
		res := mergeString(a1, a2)
		fmt.Printf("false -> %s\n\n", res)
	case -1:
		res := mergeString(a1, a2)
		fmt.Printf("false -> %s\n\n", res)
	case 0:
		res := mergeString(a1, a2)
		fmt.Printf("true -> %s\n\n", res)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("DESCRIBE: Compare two strings.\nUSAGE: compstr [STRING_1] [STRING_2]\n")
		return
	}

	r := strings.Compare(os.Args[1], os.Args[2])
	printResult(r, os.Args[1], os.Args[2])
}
