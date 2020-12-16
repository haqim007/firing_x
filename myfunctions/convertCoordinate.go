package myfunctions

import (
	"strconv"
	"strings"
)

// ConvertCoordinate to convert coordinate to integer
func ConvertCoordinate(inp string) ([2]int, error) {
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
