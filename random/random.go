package random

import (
	"math/rand"
	"time"
)

var (
	number      = "1234567890"
	upperLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerLetter = "abcdefghijklmnopqrstuvwxyz"
	letter      = upperLetter + lowerLetter
	all         = number + letter
)

func Number(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune(number)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func Letter(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune(letter)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func UpperLetter(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune(upperLetter)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func LowerLetter(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune(lowerLetter)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

func String(n int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	letters := []rune(all)
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
