package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	fmt.Println("Product costs $9.09. Tax is 9.75%.")

	f := 9.09
	t := 0.0975
	ft := f * t
	fmt.Printf("Floats: %.18f * %.18f = %.18f\n", f, t, ft)

	u := USDFromFloat64(9.09)
	ut := u.Multiply(t)
	fmt.Printf("USD: %v * %v = %v\n", u, t, ut)
}
