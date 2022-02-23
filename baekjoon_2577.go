package main

import (
	"fmt"
	"strconv"
)

func Bj2577 (a int, b int, c int) {

	num := strconv.Itoa(a * b * c)

	var i int
	result := [10]int{}

	for i=0; i<=len(num)-1; i++ {
		j, _ := strconv.Atoi(num[i:i+1])
		result[j] += 1
	}

	for i=0; i<len(result); i++ {
		fmt.Println(result[i])
	}
}
