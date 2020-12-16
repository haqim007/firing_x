package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	mf "github.com/haqim007/chess/myfunctions"
)

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

			isValid := mf.ValidateCoordinate(arrInputStr[0])
			if isValid != nil {
				fmt.Println(isValid)
				i--
				continue
			}
			res, errRes := mf.ConvertCoordinate(arrInputStr[0])
			if errRes != nil {
				fmt.Println(errRes)
				os.Exit(0)
			}

			arrInput[res[0]-1][res[1]-1] = "x"

		}
	}
	resGetVictim, errGetVictim := mf.GetVictim(arrInput)
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
	fmt.Println(arrInput)
	mf.ChessBoard(arrInput)
}
