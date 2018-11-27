package main

import (
	"fmt"
)

func main() {

	// ex01()
	// ex02()
	// ex03()
	// ex04()
	// ex05()
	// ex06()
	// ex07()
	// ex08()
	// ex09()
	ex10()
}

func ex01() {
	x := make([]int, 5, 5)

	x[0] = 1
	x[1] = 1
	x[2] = 1
	x[3] = 1
	x[4] = 1

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}

func ex02() {
	x := make([]int, 10)

	for i := 0; i < len(x); i++ {
		x[i] = i

	}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func ex03() {
	x := make([]int, 10)

	for i := 0; i < len(x); i++ {
		x[i] = i

	}

	x1 := x[:5]
	x2 := x[5:]
	x3 := x[2:7]
	x4 := x[1:6]

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
	fmt.Println(x4)

}

func ex04() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)

	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)

	fmt.Println(x)

}

func ex05() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:]...)

	fmt.Println(y)

}

func ex06() {

	states := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`,
		` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`,
		` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`,
		` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`,
		` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`,
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`,
		` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`,
		` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}

	fmt.Printf("Length: %v, cap: %v\n", len(states), cap(states))
	for _, v := range states {
		fmt.Println(v)
	}

}

func ex07() {
	xx := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i, xs := range xx {
		fmt.Printf("record: %v, values: %v\n", i, xs)
	}

}

func ex08() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Printf("Key: %v\n", k)
		fmt.Printf("\t Values: ")
		for i, x := range v {
			fmt.Printf("%v %v\t", i, x)

		}
		fmt.Println()
	}

}

func ex09() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["jensen_kevin"] = []string{"Something", "something else"}

	for k, v := range m {
		fmt.Println("Keys: ", k, "\t", "Values", v)
	}

}

func ex10() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["jensen_kevin"] = []string{"Something", "something else"}

	delete(m, "bond_james")

	for k, v := range m {
		fmt.Println("Keys: ", k, "\t", "Values", v)
	}
}
