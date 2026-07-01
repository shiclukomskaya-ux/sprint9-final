package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	size := 10
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("Ожидали слайс длиной %d, получили %d", size, len(result))
	}
	zeroResult := generateRandomElements(0)
	if len(zeroResult) != 0 {
		t.Errorf("Ожидали пустой слайс, получили %d", len(zeroResult))
	}
}
func TestMaximum(t *testing.T) {
	empty := []int{}
	if maximum(empty) != 0 {
		t.Errorf("Ожидали 0, получили %d", maximum(empty))
	}

	one := []int{2}
	if maximum(one) != 2 {
		t.Errorf("Ожидали получить один элемент, получили %d", maximum(one))
	}

	many := []int{10, 200, 30, 80, 40}
	if maximum(many) != 200 {
		t.Errorf("Ожидали получить максимально 100, полу %d", maximum(many))
	}

}
