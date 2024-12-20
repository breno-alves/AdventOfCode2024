package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func parseInput(left, right *[]int, occurrence map[int]int) error {
	file, err := os.Open("input")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, err = reader.Discard(0)
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		tokens := strings.Split(string(line), " ")
		first, err := strconv.Atoi(tokens[0])
		if err != nil {
			return err
		}
		last, err := strconv.Atoi(tokens[len(tokens)-1])
		if err != nil {
			return err
		}
		*left = append(*left, first)
		*right = append(*right, last)
		occurrence[last]++
	}
	return nil
}

func main() {
	leftInput := new([]int)
	rightInput := new([]int)
	occurrence := make(map[int]int)

	err := parseInput(leftInput, rightInput, occurrence)
	if err != nil {
		panic(err)
	}

	leftHeap := &MinHeap{}
	*leftHeap = *leftInput
	heap.Init(leftHeap)

	rightHeap := &MinHeap{}
	*rightHeap = *rightInput
	heap.Init(rightHeap)

	distance := float64(0)
	score := float64(0)
	for leftHeap.Len() > 0 {
		left := heap.Pop(leftHeap).(int)
		right := heap.Pop(rightHeap).(int)
		distance += math.Abs(float64(right - left))
		if occurrence[left] > 0 {
			score += float64(occurrence[left] * left)
		}
	}
	fmt.Printf("Distance: %d\n", int(distance))
	fmt.Printf("Similarity Score: %d\n", int(score))

}
