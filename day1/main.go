package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(left, right *[]float64) error {
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
		*right = append(*right, float64(first))
		*left = append(*left, float64(last))
	}
	return nil
}

func main() {
	left := new([]float64)
	right := new([]float64)
	err := parseInput(left, right)
	if err != nil {
		panic(err)
	}

	sort.Float64s(*right)
	sort.Float64s(*left)
	ans := float64(0)
	for i := 0; i < len(*right); i++ {
		diff := math.Abs((*right)[i] - (*left)[i])
		ans += diff
	}
	fmt.Println(ans)
}
