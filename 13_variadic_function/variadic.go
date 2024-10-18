package main

// import "fmt"

func main() {
	// total := sum(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
	total:=[]int{10,20,30}
	fmt.Println(sum(total...))
}

func sum(values ...int) int {
	total := 0
	for _,val := range values {
		total += val
		// fmt.Println(val)
	}
	return total
}