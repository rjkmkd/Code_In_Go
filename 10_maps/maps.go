package main

import (
	"fmt"
	"maps"
)

func main() {
	// creating maps
	//syntax m:= make(map[data type of key]data type of value)
	// m := make(map[string]string)
	// //setting an element
	// m["name"]="goLang"
	// m["email"]="ghh@hhv.com"
	// m["age"]="30"
	// fmt.Println(m["age"]) // 30
	// fmt.Println(m["phn"]) // zero value i.e. for string its "empty", int its 0 and bool its false

	// m:=map[string]int{"age":30}
	// // fmt.Println(m)  => map[age:30]

	// value,ok := m["age"]
	// if ok {
	// 	fmt.Println("ok: ",value)
	// }else{
	// 	fmt.Println(" not ok: ",value)

	// }

	m1:=map[string]int{"age":22, "id":1}
	m2:=map[string]int{"age":22, "id":1}
	fmt.Println(maps.Equal(m1, m2))
}