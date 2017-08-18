package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type thing1mocker struct {
	mock.Mock
}

func (t *thing1mocker) add(x, y int) int {
	args := t.Called(x, y)
	return args.Int(0)
}

func TestCallAdd(t *testing.T) {
	testObj := new(thing1mocker)
	testObj.On("add", 1, 2).Return(3)

	callAdd(testObj)
}
