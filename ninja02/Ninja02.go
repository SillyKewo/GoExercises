package main

import (
	"fmt"
)

const (
	t     = 2
	u int = 32
)

const (
	now  = 2018 + iota
	now2 = now + iota
	now3 = now + iota
	now4 = now + iota
)

func main() {

	ex01()
	ex02()
	ex03()
	ex04()
	ex05()
	ex06()
}

func ex01() {
	x := 2

	fmt.Printf("%d\t%b\t%#x", x, x, x)

}

func ex02() {
	x := 10
	y := 12

	z := x == y
	a := x <= y
	b := x >= y
	c := x != y
	d := x < y
	e := x > y

	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}

func ex03() {
	fmt.Printf("%v\t%T\n", t, t)
	fmt.Printf("%v\t%T\n", u, u)

}

func ex04() {
	i := 4
	fmt.Printf("%d %b %#x\n", i, i, i)

	o := i << 1

	fmt.Printf("%d %b %#x\n", o, o, o)

}

func ex05() {

	k := `dadasda 
	
	asdas`
	fmt.Println(k)
}

func ex06() {
	fmt.Println(now, now2, now3, now4)
}
