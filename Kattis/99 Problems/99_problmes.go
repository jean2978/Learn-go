package main

import (
	"fmt"
)

// https://open.kattis.com/problems/99problems
func main() {

	var n int
	var i, y, x int
	fmt.Scan(&n)
	x = n
	
	for n % 100 != 99 {
		n += 1
		i += 1
	}
	for x % 100 != 99{
		if x == 0 {
			break
		}
		x -= 1
		y += 1
	}
	
	if i == y {
		fmt.Println(n)
	}else if i < y{
		fmt.Println(n)
	}else{
		fmt.Println(x)
	}
}