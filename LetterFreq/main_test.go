package main

import "testing"

func testFail(t *testing.T, expected bool, actual bool) {
	t.Errorf("Expected %v, got %v", expected, actual)
}

func TestLetterFreq0(t *testing.T) {
	word := "abdefgabcdefg"

	if !equalFrequency(word) {
		testFail(t, true, false)
	}
}

func TestLetterFreq1(t *testing.T) {
	word := "abdefgabccdefg"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}

func TestLetterFreq2(t *testing.T) {
	word := "abc"

	if !equalFrequency(word) {
		testFail(t, true, false)
	}
}

func TestLetterFreq3(t *testing.T) {
	word := "cbccca"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}

func TestLetterFreq4(t *testing.T) {
	word := "cbbccca"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}

func TestLetterFreq5(t *testing.T) {
	word := "aabbcc"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}

func TestLetterFreq6(t *testing.T) {
	word := "aabbbb"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}

func TestLetterFreq7(t *testing.T) {
	word := "aabbb"

	if !equalFrequency(word) {
		testFail(t, true, false)
	}
}

func TestLetterFreq8(t *testing.T) {
	word := "cbbbcccaaa"

	if !equalFrequency(word) {
		testFail(t, true, false)
	}
}

func TestLetterFreq9(t *testing.T) {
	word := "ddaccb"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}

func TestLetterFreq10(t *testing.T) {
	word := "zz"

	if !equalFrequency(word) {
		testFail(t, true, false)
	}
}

func TestLetterFreq11(t *testing.T) {
	word := "cccd"

	if !equalFrequency(word) {
		testFail(t, true, false)
	}
}

func TestLetterFreq12(t *testing.T) {
	word := "aaaabbbbccc"

	if equalFrequency(word) {
		testFail(t, false, true)
	}
}
