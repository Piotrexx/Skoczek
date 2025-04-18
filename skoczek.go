package main

import (
	"fmt"
	"strconv"
)

var px int
var py int
var plansza [][]int
var n int
var tura int
var odwiedzone_pola int

func main() {

	var input string

	fmt.Print("Podaj wymiar planszy: ")
	fmt.Scan(&input)

	var err error
	n, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nie podałeś liczby!")
		return
	}

	plansza = make([][]int, n)
	for i := range plansza {
		plansza[i] = make([]int, n)
	}

	fmt.Print("Podaj pozycję skoczka x: ")
	fmt.Scan(&input)
	px, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nie podałeś liczby!")
		return
	}

	fmt.Print("Podaj pozycję skoczka y: ")
	fmt.Scan(&input)

	py, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nie podałeś liczby!")
		return
	}
	odwiedzone_pola = 0
	tura = 1
	fmt.Println("Plansza:")
	for _, row := range solve(n, plansza, odwiedzone_pola, [2]int{px, py}, tura) {
		fmt.Println(row)
	}
}

func solve(n int, plansza [][]int, odwiedzone_pola int, position [2]int, tura int) [][]int {
	if odwiedzone_pola == n*n {
		return plansza
	}
	plansza[position[1]-1][position[0]-1] = tura
	Ymoveset := [8]int{-1, 1, 2, 2, -2, -2, -1, 1}
	Xmoveset := [8]int{-2, -2, -1, 1, -1, 1, 2, 2}
	for i := 1; i <= 8; i++ {
		NextXmove := position[0] + Xmoveset[i-1]
		NextYmove := position[1] + Ymoveset[i-1]
		if NextXmove <= n && NextXmove >= 1 && NextYmove <= n && NextYmove >= 1 && plansza[NextYmove-1][NextXmove-1] == 0 {
			tura++
			solve(n, plansza, odwiedzone_pola+1, [2]int{NextXmove, NextYmove}, tura)
		}
	}
	return plansza
}
