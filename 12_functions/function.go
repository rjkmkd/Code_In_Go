package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("welcome to golang")
	// greeter()
	// result:=proAdder(1,2,34,5,67,7)
	// fmt.Println(result)
	// lang1, _, isLang := getLang()
	// fmt.Println(lang1,isLang)
	//function as argument to other function
	// executefunc("gilang", toUppercase)
	retFnc:=getMultiply(2)
	fmt.Println(retFnc(2))

}

// func proAdder(values ...int)int {
// 	 fmt.Println(values)
// 	total:=0
// 	for _,val :=range values{
// 		total+=val
// 	}
// 	return total
// }
// func greeter(){
// 	fmt.Println("namaste from golang")
// }

func getLang()(string, string, bool){
	return "golang" ,"javascript", true
}

func executefunc(s string, f func(s string)){
	f(s)
}
func toUppercase(s string){
	fmt.Println("uppercase string: ",strings.ToUpper(s))
}

func getMultiply(factor int) func(int)int{
	return func(i int) int {
		return factor*i
	}
}

