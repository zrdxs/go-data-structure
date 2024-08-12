package main

import (
	"math/rand"
)

func GenerateAndSum() {

	n1 := rand.Intn(10)
	n2 := rand.Intn(10)

	go sum(n1, n2)
}

func sum(a, b int) int {
	return a + b
}
