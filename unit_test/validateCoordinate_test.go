package unit_test

import "testing"
import mf "github.com/haqim007/chess/myfunctions"

func TestValidateCoordinate(t *testing.T) {
	input := "2,1"

	err := mf.ValidateCoordinate(input)

	if err != nil {
		t.Errorf("Error : %s", err)
	}
}
