package myfunctions

import (
	"errors"
	"strconv"
)

var (
	// ErrVictimNotFound string
	ErrVictimNotFound = errors.New("tidak ditemukan")
)

// GetVictim to find victim
func GetVictim(inp [8][8]string) ([7]string, error) {
	var res [7]string
	for x := 0; x < len(inp); x++ {

		for y := 0; y < len(inp); y++ {
			if inp[x][y] != "" {
				for i := 1; i < len(inp)-1; i++ {

					if x-i > 0 && y+i < 8 {
						if inp[x][y] == "x" && inp[x-i][y+i] == "x" {
							res[y] = "(" + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1) + ")"
							inp[x-i][y+i] = ""
						}
					}

					if y+i < 8 {
						if inp[x][y] == "x" && inp[x][y+i] == "x" {
							res[y] = "(" + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1) + ")"
							inp[x][y+i] = ""
						}
					}

					if x+i < 8 && y+i < 8 {
						if inp[x][y] == "x" && inp[x+i][y+i] == "x" {
							res[y] = "(" + strconv.Itoa(x+1) + "," + strconv.Itoa(y+1) + ")"
							inp[x+i][y+i] = ""
						}
					}

				}
			}
		}
	}
	check := false
	for i := range res {
		if res[i] != "" {
			check = true
		}

		if i == len(res)-1 && !check {
			return [7]string{}, ErrVictimNotFound
		}
	}
	return res, nil
}
