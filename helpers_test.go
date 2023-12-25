package main

import "testing"

func TestMaxPrecision(t *testing.T) {

    got := maxPrecision([]float64{1.1, 2.22, 3.333, 4.4444})
    want := 4

    if got != want {
        t.Errorf("YourFunction(\"input\") = %d, want %d", got, want)
    }
}