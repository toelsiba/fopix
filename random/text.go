package random

import (
	"math/rand"
	"strings"
)

var (
	lowerLetters = []rune("abcdefghijklmnopqrstuvwxyz")
	upperLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func RandWord(r *rand.Rand) string {
	n := IntForInterval(r, 3, 12)
	rs := randRunes(r, n, lowerLetters)
	return string(rs)
}

func randRunes(r *rand.Rand, n int, corpus []rune) []rune {
	rs := make([]rune, n)
	for i := range rs {
		rs[i] = corpus[r.Intn(len(corpus))]
	}
	return rs
}

func RandLine(r *rand.Rand, maxLen int) string {
	var b strings.Builder
	for {
		word := RandWord(r)
		b.WriteString(word)
		b.WriteByte(' ')
	}
	return b.String()
}
