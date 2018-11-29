package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	ice   []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println()
	ex01_02()
}

func ex01_02() {
	p1 := person{
		"Kevin",
		"Jensen",
		[]string{
			"Choco",
		},
	}

	p2 := person{
		"Ramos",
		"Ramos",
		[]string{
			"Mint",
		},
	}

	fmt.Println(p1, p2)

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
	}
}

func ex03() {
	t := truck{
		vehicle: vehicle{
			doors: 6,
			color: "blue",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 7,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t, s)

	fmt.Println(t.color, s.color)
}

func ex04() {
	s := struct {
		age  int
		name string
	}{
		4,
		"Kevin",
	}
	fmt.Println(s)
}
