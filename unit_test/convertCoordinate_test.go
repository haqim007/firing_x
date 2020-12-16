package unit_test

import "testing"
import "reflect"
import mf "github.com/haqim007/chess/myfunctions"

func TestConvertCoordinate(t *testing.T) {
	input := "2,1"
	expectedOutput := [2]int{2, 1}

	res, err := mf.ConvertCoordinate(input)

	if !reflect.DeepEqual(res, expectedOutput) {
		t.Errorf("Error : %s", err)
	}
}
