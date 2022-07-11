package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var resultString string = Length[string]("Eko")
	assert.Equal(t, "Eko", resultString)
	var resultNumber int = Length[int](1000)
	assert.Equal(t, 1000, resultNumber)
}
