package main

import "fmt"

func main() {

	strings := []string{"a", "b", "c"}
	fmt.Println(strings)
	fmt.Println(strings[0])
	fmt.Println(strings[1:])

	bbb()
	println("1")
	println("1")
	//println("1")
	//println("1")
	//println("1")
}
func aaa(m map[string]int, mt MAP) {
	println(m["a"])
	println(mt["b"])
}

func bbb() {
	aaa(map[string]int{"a": 100, "b": 200},
		MAP{"a": 100, "b": 200})
}

type MAP map[string]int
