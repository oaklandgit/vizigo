package main

import (
	"reflect"
	"testing"
)

func TestPadStringToCenter(t *testing.T) {

    got := padStringToCenter("hello", 10)
    want := "  hello   "

    if got != want {
        t.Errorf("YourFunction(\"input\") = %s, want %s", got, want)
    }

    got = padStringToCenter("hello2", 10)
    want = "  hello2  "

    if got != want {
        t.Errorf("YourFunction(\"input\") = %s, want %s", got, want)
    }
}

func TestSplitStringAt(t *testing.T) {
    
    got1, got2, got3, _ := splitStringAt("hello", 2)
    want1, want2, want3 := "he", "l", "lo"

    if got1 != want1 || got2 != want2 || got3 != want3 {
        t.Errorf("YourFunction(\"input\") = %s, %s, %s, want %s, %s, %s",
        got1, got2, got3, want1, want2, want3)
    }
}

func TestUnderlineChar(t *testing.T) {

    got := underlineChar("hello", 2)
    want := "he\033[4ml\033[0mlo"

    if got != want {
        t.Errorf("YourFunction(\"input\") = %s, want %s", got, want)
    }
}

func TestColumnToLetters(t *testing.T) {
    
    got := columnToLetters(1)
    want := "A"

    if got != want {
        t.Errorf("YourFunction(\"input\") = %s, want %s", got, want)
    }

    got = columnToLetters(27)
    want = "AA"

    if got != want {
        t.Errorf("YourFunction(\"input\") = %s, want %s", got, want)
    }
}

func TestLettersToCol(t *testing.T) {
    
    got := lettersToColumn("A")
    want := 1

    if got != want {
        t.Errorf("YourFunction(\"input\") = %d, want %d", got, want)
    }

    got = lettersToColumn("AA")
    want = 27

    if got != want {
        t.Errorf("YourFunction(\"input\") = %d, want %d", got, want)
    }
}

func TestSplitAlphaNumeric(t *testing.T) {
        
    got1, got2 := splitAlphaNumeric("A1")
    want1, want2 := "A", "1"

    if got1 != want1 || got2 != want2 {
        t.Errorf("YourFunction(\"input\") = %s, %s, want %s, %s", got1, got2, want1, want2)
    }

    got1, got2 = splitAlphaNumeric("AA11")
    want1, want2 = "AA", "11"

    if got1 != want1 || got2 != want2 {
        t.Errorf("YourFunction(\"input\") = %s, %s, want %s, %s", got1, got2, want1, want2)
    }
}

func TestAlphaNumericToPosition(t *testing.T) {

    got := alphaNumericToPosition("A1")
    want := Vector{row: 1, col: 1}

    if got != want {
        t.Errorf("YourFunction(\"input\") = %v, want %v", got, want)
    }

    got = alphaNumericToPosition("AA11")
    want = Vector{row: 11, col: 27}

    if got != want {
        t.Errorf("YourFunction(\"input\") = %v, want %v", got, want)
    }
}

func TestMaxPrecision(t *testing.T) {

    got := maxPrecision([]float64{1.1, 2.22, 3.333, 4.4444})
    want := 4

    if got != want {
        t.Errorf("YourFunction(\"input\") = %d, want %d", got, want)
    }
}

func TestExtractReferences(t *testing.T) {
    input := "=SUM(AA10:AB12, B3, HH1)"
    got := extractReferences(input)
    want := []string{"AA10:AB12", "B3", "HH1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseReferences(%q) = %v, want %v", input, got, want)
	}
}

func TestPositionsFromReferences(t *testing.T) {
    input := []string{"A1:B2", "B3"}
    got := positionsFromReferences(input)
    want := []Vector{
        {col: 1, row: 1}, // A1
        {col: 2, row: 1}, // B1
        {col: 1, row: 2}, // A2
        {col: 2, row: 2}, // B2
        {col: 2, row: 3}} // B3

    if !reflect.DeepEqual(got, want) {
        t.Errorf("parseReferences(%q) = %v, want %v", input, got, want)
    }
}
