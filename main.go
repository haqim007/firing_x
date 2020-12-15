package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	// ErrInvalidInputFormat string
	ErrInvalidInputFormat = errors.New("format koordinat yang dimasukkan salah. contoh: 2,3")
	// ErrInvalidInputLength string
	ErrInvalidInputLength = errors.New("batas minimal 1,1 dan maksimal koordinat yaitu 8,8. contoh: 2,3")
)

func chessBoard(input [8][8]string) {
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
				fmt.Printf("  %s  |", input[j-1][7-i])
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

func validateCoordinate(inp string) error {
	inpArr := strings.Split(inp, ",")
	for i := range inpArr {
		inpArr[i] = strings.TrimSpace(inpArr[i])
		inpInt, err := strconv.Atoi(inpArr[i])
		if err != nil || len(inpArr) != 2 {
			return ErrInvalidInputFormat
		}
		if inpInt > 8 || inpInt < 1 {
			return ErrInvalidInputLength
		}
	}
	return nil
}

func convertCoordinate(inp string) ([2]int, error) {
	inpArr := strings.Split(inp, ",")
	var res [2]int
	for i := range inpArr {
		inpArr[i] = strings.TrimSpace(inpArr[i])
		inpInt, err := strconv.Atoi(inpArr[i])
		if err != nil {
			res[i] = 0
			return res, err
		}
		res[i] = inpInt

	}
	return res, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("> Start ")
	var arrInput [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			arrInput[i][j] = " "
		}
	}

	for i := 0; i < 8; i++ {

		fmt.Println("> Masukkan Bidak", i+1, " (x,y):")
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			inputStr := strings.TrimSuffix(cmdString, "\n")
			arrInputStr := strings.Fields(inputStr)

			if arrInputStr[0] == "exit" {
				os.Exit(0)
			}
			// else{

			// }

			// inputArrStr := strings.Split(arrInputStr[0], ",")
			isValid := validateCoordinate(arrInputStr[0])
			if isValid != nil {
				fmt.Println(isValid)
				i--
				continue
			}
			res, errRes := convertCoordinate(arrInputStr[0])
			if errRes != nil {
				fmt.Println(errRes)
				os.Exit(0)
			}
			// text := strings.Split(arrInputStr[0], ",")
			// for i := range text {
			// 	text[i] = strings.TrimSpace(text[i])
			// }
			// fmt.Println(res[0], res[1])
			arrInput[res[0]-1][res[1]-1] = "x"
			// fmt.Println(arrInput)
		}
	}
	chessBoard(arrInput)
}
