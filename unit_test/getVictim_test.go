package unit_test

import "testing"
import "errors"
import mf "github.com/haqim007/chess/myfunctions"

func TestGetVictim(t *testing.T) {
	t.Log("\n GetVictim 1")
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
	t.Log("\n GetVictim 2")
	input2 := [8][8]string{
		{"", "", "x", "", "", "", "", ""},
		{"x", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "x", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "", "", "", "", ""},
		{"", "", "", "x", "", "", "", ""},
	}
	expectedError := errors.New("tidak ditemukan")

	_, err2 := mf.GetVictim(input2)
	if err2.Error() != expectedError.Error() {
		t.Errorf("Expected Error : %s . What we got : %s", expectedError, err2)
	} else {
		t.Logf("Expected Error : %s . What we got : %s", expectedError, err2)
	}
}
