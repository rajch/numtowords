package numtowords_test

import (
	"testing"

	"github.com/rajch/numtowords"
)

func TestInvalidNumber(t *testing.T) {
	_, err := numtowords.Convert(numtowords.MaxNum + 1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}

	_, err = numtowords.Convert(-1)
	if err == nil {
		t.Log("expected error")
		t.Fail()
	}
}

func TestZero(t *testing.T) {
	result, err := numtowords.Convert(0)
	if err != nil {
		t.Log("Convert with 0 returned an error")
		t.FailNow()
	}

	if result != "zero" {
		t.Logf("expected 'zero', received '%v'", result)
		t.FailNow()
	}
}

func TestUnits(t *testing.T) {
	units := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	for index, value := range units {
		t.Logf("Now testing %v...", index)

		result, err := numtowords.Convert(index)
		if err != nil {
			t.Logf("error while converting %v: %v", index, err)
			t.Fail()
		}

		if result != units[index] {
			t.Logf("while converting %v, expected %v, got %v", index, value, result)
			t.Fail()
		}
	}
}

func TestTens(t *testing.T) {
	testcases := map[int]string{
		34: "thirty four",
		48: "forty eight",
		99: "ninety nine",
		60: "sixty",
	}

	for k, v := range testcases {
		t.Logf("Now testing %v...", k)

		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("error while converting %v: %v", k, err)
			t.Fail()
		}

		if result != v {
			t.Logf("while converting %v, expected %v, got %v", k, v, result)
			t.Fail()
		}
	}
}

func TestHundreds(t *testing.T) {
	testcases := map[int]string{
		109: "one hundred and nine",
		333: "three hundred and thirty three",
		700: "seven hundred",
		340: "three hundred and forty",
	}

	for k, v := range testcases {
		t.Logf("Now testing %v...", k)

		result, err := numtowords.Convert(k)
		if err != nil {
			t.Logf("error while converting %v: %v", k, err)
			t.Fail()
		}

		if result != v {
			t.Logf("while converting %v, expected %v, got %v", k, v, result)
			t.Fail()
		}
	}

}
