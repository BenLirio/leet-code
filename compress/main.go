package main

import (
	"fmt"
	"errors"
	"strconv"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("Enter Number of binary numbers to print")
		os.Exit(1)
	}
	var err error
	var N int
	
 	N, err = strconv.Atoi(os.Args[1])
	check(err)


	var binaryArray[] bool
	var numBits int
	var tmpN int = N
	for tmpN > 0 {
		tmpN = tmpN / 2
		numBits = numBits + 1
	}
	for i := 0; i < N; i++ {
		binaryArray, err = decimalToBinaryArray(i, numBits)
		check(err)
		for _, bit := range binaryArray {
			if bit == true {
				fmt.Printf("%c", '1')
			} else {
				fmt.Printf("%c", '0')
			}
		}
		fmt.Printf("\n")
	}
	os.Exit(0)
}

func power(base int, exponent int) (int) {
	var out int = 1
	for i := 0; i < exponent; i++ {
		out = out * base
	}
	return out
}

func decimalToBinaryArray(decimal int, numBits int) ([]bool, error) {
	var out[] bool = make([]bool, numBits)
	var significance = power(2, (numBits - 1))
	if decimal > ((significance * 2) - 1) {
		return out, errors.New("Num bits too small (Overflow)")
	}
	var digit = numBits - 1
	for digit >= 0 {
		if significance <= decimal {
			decimal = decimal - significance
			out[digit] = true
		} else {
			out[digit] = false
		}
		significance = significance / 2
		digit = digit - 1
	}
	return out, nil
}

