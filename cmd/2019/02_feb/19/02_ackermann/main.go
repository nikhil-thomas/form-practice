package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

//ackermann function
//
// A(m,n) = n + 1,             if m = 0
//        = A(m-1,1),          if n = 0
//        = A(m-1, A(m, n-1)), if m > 0 and n > 0
func ackermann(m, n int64) (int64, error) {
	if m == 0 {
		return n + 1, nil
	}
	if n == 0 {
		return ackermann(m-1, 1)
	}
	if m > 0 && n > 0 {
		n2, err := ackermann(m, n-1)
		if err != nil {
			return 0, err
		}
		return ackermann(m-1, n2)
	}
	return 0, errors.New("invalid input")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("please provide 2 integer arguments")
		os.Exit(1)
	}
	m, err := strconv.ParseInt(os.Args[1], 10, 10)
	if err != nil {
		log.Fatalln(err)
	}
	n, err := strconv.ParseInt(os.Args[2], 10, 10)
	if err != nil {
		log.Fatalln(err)
	}
	ackVal, err := ackermann(m, n)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("ackerman(%d, %d) = %d\n", m, n, ackVal)
}
