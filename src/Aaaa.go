package main

func main() {
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
