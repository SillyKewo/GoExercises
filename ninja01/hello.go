package main

import "fmt"

type test int

var x = 42
var y = "James Bond"
var z = true

var t test

func main() {
	ex01()
	ex02()
	ex03()
	ex04()
}

func ex01() {
	x := 42
	y := "James bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("%T\t%T\t%T\n", x, y, z)
}

func ex02() {
	fmt.Println(x, y, z)
	fmt.Printf("%T\t%T\t%T\n", x, y, z)
}

func ex03() {
	s := fmt.Sprint(x, y, z)

	fmt.Println(s)
}

func ex04() {
	fmt.Printf("%T\n", t)
	fmt.Printf("%v\n", t)
	t = 43
	fmt.Println(t)
}
