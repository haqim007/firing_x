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
	ErrInvalidInputLength = errors.New("batas minimal 1,1 dan maksimal koordinat yaitu 8,8")
	// ErrVictimNotFound string
	ErrVictimNotFound = errors.New("tidak ditemukan")
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

func getVictim(inp [8][8]string) ([7]string, error) {
	var res [7]string
	for x := 0; x < len(inp); x++ {
		// fmt.Println(inp)
		for y := 0; y < len(inp); y++ {
			if inp[x][y] != "" {
				for i := 1; i < len(inp)-1; i++ {
					if x-i > 0 && y+i < 8 {
						if inp[x][y] == "x" && inp[x-i][y+i] == "x" {
							fmt.Println(x+1, ",", y+1, ",", inp[x-i][y+i])
							res[y] = "(" + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1) + ")"
							inp[x-i][y+i] = ""
						}
					}

					if y+i < 8 {
						if inp[x][y] == "x" && inp[x][y+i] == "x" {
							fmt.Println(x+1, ",", y+1, ",", inp[x][y+i])
							res[y] = "(" + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1) + ")"
							inp[x][y+i] = ""
						}
					}

					// if x+i < 8 && y+i < 8 {
					// 	if inp[x+i][y+i] == "x" {
					// 		fmt.Println(x+1, ",", y+1, ",", inp[x+i][y+i])
					// 		res[y] = "(" + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1) + ")"
					// 		inp[x+i][y+i] = ""
					// 	}
					// }
				}
			}
		}
	}
	if len(res) == 0 {
		return [7]string{}, ErrVictimNotFound
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

			arrInput[res[0]-1][res[1]-1] = "x"

		}
	}
	resGetVictim, errGetVictim := getVictim(arrInput)
	if errGetVictim != nil {
		fmt.Println(errGetVictim)
	} else {
		fmt.Print("> Hasil: ")
		for i := range resGetVictim {
			if resGetVictim[i] != "" {
				fmt.Print(resGetVictim[i], " ")
			}
		}
		fmt.Print("\n")
	}
	chessBoard(arrInput)
}
