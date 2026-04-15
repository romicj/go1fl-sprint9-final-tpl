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
	slice := make([]int, 0, size)
	for range size {
		slice = append(slice, rand.Intn(1000))
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) (max int) {
	for _, val := range data {
		if val > max {
			max = val
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	wg := sync.WaitGroup{}
	chunkSize := len(data) / CHUNKS
	maxesArr := make([]int, 8)
	for i := range CHUNKS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			chunkedArr := data[i*chunkSize : (i+1)*chunkSize]

			max := maximum(chunkedArr)
			maxesArr[i] = max
		}()
	}
	wg.Wait()
	return maximum(maxesArr)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	start := time.Now()
	slice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	max, elapsed := maximum(slice), time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	start = time.Now()

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	max, elapsed = maxChunks(slice), time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
