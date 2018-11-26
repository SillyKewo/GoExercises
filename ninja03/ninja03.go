package main

import (
	"fmt"
)

func main() {

	// ex01()
	// ex02()
	// ex03()
	ex04()
	ex05()
	ex06()
	ex07()
	ex08()
	ex09("Volleyball")
	ex10()
}

func ex01() {

	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

}

func ex02() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		fmt.Printf("\t%#U\n", i)
		fmt.Printf("\t%#U\n", i)
		fmt.Printf("\t%#U\n", i)

	}
}

func ex03() {

	y := 1995

	for y < 2019 {
		fmt.Println(y)
		y++
	}

}

func ex04() {
	y := 1995
	for {
		if y > 2018 {
			break
		}
		fmt.Println(y)
		y++
	}

}

func ex05() {

	for i := 10; i < 101; i++ {
		fmt.Println(i % 4)
	}

}

func ex06() {

	if true && false {
		fmt.Println("Wont happen")
	} else if true || false {
		fmt.Println("Should print this")
	}

}

func ex07() {

	if true && false {
		fmt.Println("Wont happen")
	} else if true || false {
		fmt.Println("Should print this")
	} else {
		fmt.Println("Not happening")
	}
}

func ex08() {

	switch {

	case 2 == 3:
		fmt.Println("NOPE")

	case 2 == 2:
		fmt.Println("YUP")

	}
}

func ex09(favSport string) {

	switch favSport {
	case "Football", "Volleyball", "Handball":
		fmt.Println("Nice choice")
	default:
		fmt.Println("Bad choice")
	}
}

func ex10() {

}
