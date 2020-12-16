package unit_test

import "testing"
import mf "github.com/haqim007/chess/myfunctions"

func TestGetVictim(t *testing.T) {
	input := [8][8]string{
		{"", "", "", "", "", "x", "x", ""},
		{"x", "", "", "", "", "", "x", "x"},
		{"", "x", "", "", "", "", "", ""},
		{"", "", "", "x", "", "", "", ""},
		{"", "", "x", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
	}
	expectedOutput := [3]string{"(2,1)", "(5,3)", "(1,5)"}

	res, err := mf.GetVictim(input)
	check := false
	for i := range res {
		t.Logf("Expected output one of: %s %s %s", expectedOutput[0], expectedOutput[1], expectedOutput[2])
		t.Logf("output: %s", res[i])
		if res[i] == expectedOutput[0] || res[i] == expectedOutput[1] || res[i] == expectedOutput[2] {
			check = true
		}
		if i == len(res)-1 && !check {
			t.Errorf("Error : %s", err)
		}
	}
}
