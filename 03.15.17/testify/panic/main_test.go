package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCat(t *testing.T) {
	res := cat("hi", "hi")
	assert.Equal(t, "hihi", res, "should be the same")

	// defer a recovery
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:\t", r)
	// 		if r != "Oh snap!" {
	// 			t.Error("not the right panic...")
	// 		}
	// 	}
	// }()
	// res = cat("hi", "panic")

	// assert.Panics(t, func() { cat("hi", "panic") }, "should panic with arg2 == panic")
	// assert.Panics(t, func() { cat("hi", "yarr") }, "should panic with arg2 == panic")
}
