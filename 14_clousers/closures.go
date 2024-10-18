package main

import "fmt"

func main() {
	getCounter := counter()
	fmt.Println(getCounter())
	fmt.Println(getCounter())
	fmt.Println(getCounter())
	fmt.Println(getCounter())
}

func counter() func() int {
	count := 0
	return func() int {
		count +=1
		return count
	}
}