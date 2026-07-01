package main

import (
	"fmt"
	"math/rand"

	sqlite3 "modernc.org/sqlite/lib"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size == 0 {
		return []int{}
	}
	numbers := make([]int, size) // ваш код здесь
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(SIZE)
	}
	return numbers

}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	maximum := data[0] // ваш код здесь
	for i := 1; i < len(data); i++ {
		if data[i] > maximum {
			maximum = data[i]
		}
	}
	return maximum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	res := []int{}
	length := len(data)

	part := length / CHUNKS
	remainder := length % CHUNKS

	start := 0
	for i := 0; i < CHUNKS; i++ {
		currentSize := part
		if remainder > 0 {
			currentSize++
			remainder--
		}
		end := start + currentSize
		max := maximum(data[start:end])
		res = append(res, max)
		start = end
	}
	return maximum(res)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
