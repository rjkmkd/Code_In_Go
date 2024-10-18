package main

import "fmt"

func main() {
	//range used for iterating dataStructures
	//nums := []int{1, 2, 3, 4, 5, 6, 7}

	//itrating with normal for loop
	// sum:=0
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// 	sum +=nums[i]
	// }
	// fmt.Println(sum)
	// sum:=0
	// for index, element:= range nums{
	// 	fmt.Print(index)
	// 	fmt.Print(" ==> ")
	// 	fmt.Println(element)
	// 	sum +=element
	// }
	// fmt.Println("sum = ",sum)

	//iterate over map

	// m:= map[string]string{"name":"Raj","email":"ghgh@jhj.com"}

	// for key,value := range m{
	// 	fmt.Print(key + " : ")
	// 	fmt.Println(value)
	// }

	//iterate over string
	//iunicode code point rune
	for starting_byte, unicode_code := range "🏠golang"{
		fmt.Println(starting_byte, string(unicode_code))
	}

}