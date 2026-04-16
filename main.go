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
		return nil
	}
	slice := make([]int, 0, size)
	for range size {
		slice = append(slice, rand.Int())
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
	if len(data) == 0 {
		return 0
	}
	chunkSize := (len(data) / CHUNKS) + 1
	maxesArr := make([]int, 8)
	for i := range CHUNKS {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := i * chunkSize
			end := (i + 1) * chunkSize
			if start > len(data) {
				return
			}
			if end > len(data) {
				end = len(data)
			}
			chunkedArr := data[start:end]

			maxesArr[i] = maximum(chunkedArr)
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
