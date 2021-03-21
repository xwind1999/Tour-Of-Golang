// Tour of go basic number 5
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main()  {
	fmt.Println(add(42, 13))
}