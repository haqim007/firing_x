package myfunctions

import (
	"fmt"
)

// ChessBoard to draw chess board
func ChessBoard(input [8][8]string) {
	for i := 0; i < len(input[i]); i++ {
		if i == 0 {

			for j := 0; j < len(input[i])+1; j++ {
				if j == 0 {
					fmt.Print("     ")
				} else if j == 1 {
					fmt.Printf("    %d   ", j)
				} else {
					fmt.Printf("  %d   ", j)
				}
			}
			fmt.Print("\n")
			for j := 0; j < len(input[i])+1; j++ {
				if j == 0 {
					fmt.Print("       ")
				} else {
					fmt.Print("______")
				}
			}
		} else {

			for j := 0; j < len(input[i])+1; j++ {
				if j == 0 {
					fmt.Print("       ")
				} else {
					fmt.Print("______")
				}
			}
		}

		fmt.Print("\n")
		fmt.Print("\n")
		for j := 0; j < len(input[i])+1; j++ {
			if j == 0 {
				fmt.Printf("   %d  ", 8-i)
			} else if j == 1 {
				fmt.Printf("|  %s  |", input[j-1][7-i])
			} else {
				fmt.Printf("  %s  |", input[j-1][7-i]) //7-i
			}
		}

		if i == 7 {
			fmt.Print("\n")
			for j := 0; j < len(input[i])+1; j++ {
				if j == 0 {
					fmt.Print("       ")
				} else {
					fmt.Print("______")
				}
			}
		} else {
			fmt.Print("\n")
		}
	}
}
