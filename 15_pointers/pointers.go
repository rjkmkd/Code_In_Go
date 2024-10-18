package main

import "fmt"

func main() {
	// num:=1
	// changeNum(&num)
	// fmt.Println("after change the number, num = ",num)
	nums:=[]int{1,5,8}
	addElement(&nums)
	fmt.Println(nums)
}

func changeNum(num *int){
	*num = 5
	fmt.Println("num inside changeNum func: ",*num)
	
}

func addElement(nums *[]int){
	*nums = append(*nums, 4)
}