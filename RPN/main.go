package main

import (
	"fmt"
	"os"
	"strconv"
)

func Rpn(s string) {
	if len(s) < 3 {
		fmt.Println("Error")
		return
	}
	val := []int{}
	for _, i := range s {
		if i == ' ' {
			continue
		}
		if i >= '0' && i <= '9' {
			a, _ := strconv.Atoi(string(i))
			val = append(val, a)
		} else {
			if len(val) != 2 {
				fmt.Println("Error")
				return
			}
			Val1 := val[0]
			Val2 := val[1]
			val = []int{}
			Op := i
			switch Op {
			case '+':
				val = append(val, Val1+Val2)
			case '-':
				val = append(val, Val1-Val2)
			case '*':
				val = append(val, Val1*Val2)
			case '/':
				if Val2 == 0 {
					fmt.Println("Error")
					return
				}
				val = append(val, Val1/Val2)
			case '%':
				if Val2 == 0 {
					fmt.Println("Error")
					return
				}
				val = append(val, Val1%Val2)

			}
			if len(val) == 0 {
				fmt.Println("Error")
				return
			}
		}
	}
	if len(val) != 1 {
		fmt.Println("Error")
		return
	}
	fmt.Println(val[0])
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}
	Rpn(os.Args[1])
}
