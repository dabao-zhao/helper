package xrandom

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	fmt.Println(Number(10))
	fmt.Println(Letter(10))
	fmt.Println(UpperLetter(10))
	fmt.Println(LowerLetter(10))
	fmt.Println(String(10))
}
