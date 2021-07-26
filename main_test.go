package main

import (
	"testing"
)

func TestIndexOfMaxAndMaxElement(t *testing.T) {
	tests := []struct {
		name  string
		number []int
		IndexWant  int
		MaxElementWant int
	}{
		{ "All of numbers is positive",  []int{1, 4, 25, 25,  100},  4,  100},
		{ "All of numbers is negative", []int{-1, -4, -25, -25, -100}, 0, -1},
		{ "Numbers has pos and neg",  []int{-1, 5, -20, 2},  1,  5},
	}
	for _, test := range tests {
		index, element := IndexOfMaxAndMaxElement(test.number)
		if index != test.IndexWant || element != test.MaxElementWant{
			t.Errorf("IndexOfMaxAndElement test %s gotIndex %d Indexwant %d gotElement %d ElementWant %d",
			test.name, index, test.IndexWant, element, test.MaxElementWant)
		}

	}
}
//написать функц, которая принимает слайс и возвращает сумму элементов слайса
//найти среднее значение элемента слайса
//к фнкциям написать unit tests к обоим цункциям