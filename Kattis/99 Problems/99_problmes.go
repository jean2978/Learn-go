package main

import (
	"fmt"
)

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
		if x == 0{
			break
		}
		x -= 1
		y += 1
	}

	// Check the one which is closest to the number ending with 99
	// TODO
	

}