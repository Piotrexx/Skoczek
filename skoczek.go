package main

import (
	"fmt"
	"strconv"
)

var x int
var y int
var plansza [][]int

func main() {
	var input string

	fmt.Print("Podaj szerokość planszy: ")
	fmt.Scan(&input)

	var err error
	x, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nie podałeś liczby!")
		return
	}

	fmt.Print("Podaj wysokość planszy: ")
	fmt.Scan(&input)

	y, err = strconv.Atoi(input)
	if err != nil {
		fmt.Println("Nie podałeś liczby!")
		return
	}

	plansza = make([][]int, y)
	for i := range plansza {
		plansza[i] = make([]int, x)
	}

	fmt.Println("Plansza:")
	for _, row := range plansza {
		fmt.Println(row)
	}

	position := [2]int{0, 4}
	var previous_position [2]int
	current_moveset := check_movement(&position)

	move(current_moveset, previous_position, position, 1)

	fmt.Println("Plansza:")
	for _, row := range plansza {
		fmt.Println(row)
	}
}

func check_movement(position *[2]int) [][2]int {
	Xoffsets := [4]int{-2, -1, 1, 2}
	Yoffsets := Xoffsets
	var available_moves [][2]int

	for i := range Xoffsets {
		for j := range Yoffsets {
			if Yoffsets[i] == Xoffsets[j] || Yoffsets[i] == -Xoffsets[j] {
				continue
			}

			newX := position[1] + Xoffsets[j]
			newY := position[0] + Yoffsets[i]

			if newX < 0 || newY < 0 || newX >= x || newY >= y {
				continue
			}

			if plansza[newY][newX] == 0 {
				available_moves = append(available_moves, [2]int{newY, newX})
			}
		}
	}
	return available_moves
}

// dokończyć

func move(current_moveset [][2]int, previous_position [2]int, position [2]int, path int) {
	path = 0
	for nextMove := range current_moveset {
		previous_position = position
		position = current_moveset[nextMove+path]

		current_moveset = check_movement(&position)
		if len(current_moveset) == 0 {
			if len(check_movement(&previous_position)) == 1 {
				return
			}
			position = previous_position
			path++
			break
		}

		plansza[position[0]][position[1]] = 1
		break
	}

	move(current_moveset, previous_position, position, path)
}
