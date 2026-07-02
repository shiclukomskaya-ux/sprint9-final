package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = rand.Intn(SIZE)
	}
	return numbers
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkMaximums := make([]int, CHUNKS)
	var wg sync.WaitGroup

	chunkSize := len(data) / CHUNKS
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		begin := i * chunkSize
		end := begin + chunkSize

		go func(slice []int, idx int) {
			defer wg.Done()
			chunkMaximums[idx] = maximum(slice)
		}(data[begin:end], i)
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE) // ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")
	start1 := time.Now()
	max := maximum(data)
	elapsed := time.Since(start1).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start2 := time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start2).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
