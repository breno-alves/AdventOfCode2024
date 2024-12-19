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

func parseInput(left, right *[]int) error {
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
		*right = append(*right, first)
		*left = append(*left, last)
	}
	return nil
}

func main() {
	leftInput := new([]int)
	rightInput := new([]int)
	err := parseInput(leftInput, rightInput)
	if err != nil {
		panic(err)
	}

	leftHeap := &MinHeap{}
	*leftHeap = *leftInput
	heap.Init(leftHeap)

	rightHeap := &MinHeap{}
	*rightHeap = *rightInput
	heap.Init(rightHeap)

	ans := float64(0)
	for leftHeap.Len() > 0 {
		left := heap.Pop(leftHeap).(int)
		right := heap.Pop(rightHeap).(int)
		ans += math.Abs(float64(right - left))
	}
	fmt.Println(ans)
}
