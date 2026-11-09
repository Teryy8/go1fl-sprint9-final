package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		fmt.Println("slice size error. Size must be bigger then 0")
		return []int{}
	}

	now := time.Now()
	seed := now.Unix()
	source := rand.NewSource(seed)
	random := rand.New(source)

	res := make([]int, size)
	for i := range size {
		res[i] = random.Int() + 1
	}
	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if (len(data) == 0) {
		fmt.Println("slice size error. Size must be bigger then 0")
		return 0
	} else if (len(data) == 1) {
		return data[0]
	}

	maxi := data[0]
	for _, v := range data {
		if v > maxi {
    		maxi = v
		}
	}
	return maxi
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup
	maximus := make([]int, 8)

	cutSize := len(data) / CHUNKS
	for i := range CHUNKS {
		wg.Add(1)
		start := i * cutSize
		end := start + cutSize	

		if i == CHUNKS-1 {
			end = len(data)
		}
		go func(arr []int) {
			defer wg.Done()
			maximus[i] = maximum(arr)
		}(data[start:end])
	}
	wg.Wait()
	return maximum(maximus)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	before := time.Now()
	max := maximum(data)
	elapsed := time.Since(before).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)

	before = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(before).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
