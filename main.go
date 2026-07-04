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

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)

	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1_000_000)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	results := make([]int, CHUNKS)

	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize

		// последний кусок забирает остаток
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(i, start, end int) {
			defer wg.Done()

			max := data[start]

			for _, v := range data[start:end] {
				if v > max {
					max = v
				}
			}

			results[i] = max
		}(i, start, end)
	}

	wg.Wait()

	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	data := generateRandomElements(SIZE)

	// 1 поток
	fmt.Println("Ищем максимум в 1 поток")

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимум: %d\nВремя: %d мкс\n", max, elapsed)

	// 8 потоков
	fmt.Printf("\nИщем максимум в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимум: %d\nВремя: %d мкс\n", max, elapsed)
}
