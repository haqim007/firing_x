package myfunctions

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrInvalidInputFormat string
	ErrInvalidInputFormat = errors.New("format koordinat yang dimasukkan salah. contoh: 2,3")
	// ErrInvalidInputLength string
	ErrInvalidInputLength = errors.New("batas minimal 1,1 dan maksimal koordinat yaitu 8,8")
)

// ValidateCoordinate to validate input coordinate
func ValidateCoordinate(inp string) error {
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
