package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := [10]int{5, 1, 2, 5, 4, 2, 5, 1, 1, 2}
	first := arr[:4]
	second := arr[4:]

	run(first, second)
}

func run(first, second []int) {
	printSlice(getUniqueValues(first))
	fmt.Print(", ")
	printSlice(getUniqueValues(second))
	fmt.Println("")
	printSlice(findIntersections(first, second))
	fmt.Println("")
	printSlice(unionSlicesWithUniqueValues(first, second))
}

func unionSlicesWithUniqueValues(first, second []int) []int {
	res := append(first, second...)
	return getUniqueValues(res)
}

func getUniqueValues(slice []int) []int {
	m := make(map[int]struct{}, len(slice))
	for _, value := range slice {
		m[value] = struct{}{}
	}
	return parseMapToSlice(m)
}

func findIntersections(first, second []int) []int {
	m := make(map[int]int)

	for _, value := range getUniqueValues(first) {
		m[value]++
	}
	for _, value := range getUniqueValues(second) {
		m[value]++
	}

	var intersections []int
	for key, value := range m {
		if value > 1 {
			intersections = append(intersections, key)
		}
	}
	return intersections
}

func parseMapToSlice(m map[int]struct{}) []int {
	res := make([]int, 0, len(m))
	for key := range m {
		res = append(res, key)
	}
	return res
}

func printSlice(slice []int) {
	fmt.Print("[")
	fmt.Print(strings.Join(convertIntToStringSlice(slice), " "))
	fmt.Print("]")
}

func convertIntToStringSlice(slice []int) []string {
	ret := make([]string, len(slice))
	for i, value := range slice {
		ret[i] = strconv.Itoa(value)
	}
	return ret
}
