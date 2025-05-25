package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	run()
}

func run() {
	firstInput := convertStringToIntArray(getInputStringFromTerminal())
	printSlice(
		findIntersections(
			firstInput,
			convertStringToIntArray(
				getInputStringFromTerminal(),
			),
		), firstInput)
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

func printSlice(slice, firstInput []int) {
	order := make(map[int]int)
	for idx, val := range firstInput {
		order[val] = idx
	}

	sort.Slice(slice, func(i, j int) bool {
		return order[slice[i]] < order[slice[j]]
	})
	fmt.Print(strings.Join(convertIntToStringSlice(slice), " "))
}

func convertIntToStringSlice(slice []int) []string {
	ret := make([]string, len(slice))
	for i, value := range slice {
		ret[i] = strconv.Itoa(value)
	}
	return ret
}

func getInputStringFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if len(input) == 0 {
		err = errors.New("empty intersection")
	}
	errorHandler(err)
	return input
}

func convertStringToIntArray(input string) []int {
	arrFields := strings.Fields(input)
	result := make([]int, 0, len(arrFields))
	for _, value := range arrFields {
		num, err := strconv.Atoi(value)
		errorHandler(err)
		result = append(result, num)
	}
	return result
}

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("invalid input\ncause: %s\n", err.Error())
		os.Exit(1)
	}
}
