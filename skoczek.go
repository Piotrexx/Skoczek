package main

import (
	"fmt"
	"math"
	"strconv"
)

var px int
var py int
var plansza [][]int
var n int
var NextXmove, NextYMove int

const alfabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

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
	tura := 1
	fmt.Println("Plansza:")
	answer := Warnsdorff(n, plansza, [2]int{px, py}, tura)
	// formatowanie planszy

	for i := range n {
		fmt.Printf("   %s", string(rune(alfabet[i])))
	}
	fmt.Println()
	for range n {
		fmt.Print("------")
	}
	fmt.Println()

	for iter, row := range answer {
		fmt.Printf("%d| ", iter+1)
		for _, step := range row {
			if step < 10 {
				fmt.Printf("%d   ", step)
				continue
			}
			fmt.Printf("%d  ", step)
		}
		fmt.Println()

	}
}

func count_moves(plansza [][]int, position [2]int, n int) int {
	// tablice dostępnych ruchów
	Ymoveset := [8]int{-1, 1, 2, 2, -2, -2, -1, 1}
	Xmoveset := [8]int{-2, -2, -1, 1, -1, 1, 2, 2}
	counter := 0
	for i := 1; i <= 8; i++ {
		NextXmove := position[0] + Xmoveset[i-1]
		NextYmove := position[1] + Ymoveset[i-1]
		// Jeżeli można wykonać ruch, podlicz go
		if NextXmove <= n && NextXmove >= 1 && NextYmove <= n && NextYmove >= 1 && plansza[NextYmove-1][NextXmove-1] == 0 {
			counter++
		}
	}
	return counter
}

func Warnsdorff(n int, plansza [][]int, position [2]int, tura int) [][]int {
	// Jeżeli liczba wykonanych ruchów równa się liczbie wszystkich pól planszy zwróć rozwiązaną planszę
	if tura == n*n {
		plansza[position[1]-1][position[0]-1] = tura
		return plansza
	}
	// Największa liczba 64 bitowa
	min_counter := math.MaxInt
	// tablice dostępnych ruchów
	Ymoveset := [8]int{-1, 1, 2, 2, -2, -2, -1, 1}
	Xmoveset := [8]int{-2, -2, -1, 1, -1, 1, 2, 2}
	// Zaznaczenie w komórce na której skoczek się znajduję, który to jest ruch
	plansza[position[1]-1][position[0]-1] = tura
	for i := 1; i <= 8; i++ {
		// tymczasowe wartości x i y
		temp_x := position[0] + Xmoveset[i-1]
		temp_y := position[1] + Ymoveset[i-1]
		// liczba możliwych ruchów dla tymczasowych pozycji
		count := count_moves(plansza, [2]int{temp_x, temp_y}, n)
		// Sprawdzenie czy ruch można wykonać
		if temp_x <= n && temp_x >= 1 && temp_y <= n && temp_y >= 1 && plansza[temp_y-1][temp_x-1] == 0 {
			// Jeżeli liczba możliwych ruchów jest większa od najnowszej minimalnej liczby możliwych ruchów nadpisz pozycję i NOWE tymczasowe pozycję
			if count < min_counter {
				min_counter = count
				// tymczasowe NOWE pozycje skoczka
				NextXmove = temp_x
				NextYMove = temp_y
			}
		} else {
			continue
		}

	}
	// Zmiana pozycji skoczka
	position[0] = NextXmove
	position[1] = NextYMove
	// Inkrementacja tury
	tura++
	Warnsdorff(n, plansza, position, tura)
	return plansza
}
