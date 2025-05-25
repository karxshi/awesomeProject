package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func main() {
	run()
}

func run() {
	input := getWordsFromTerminal()
	k := getK()
	printResultSlice(
		getResultFromKV(
			limitKV(
				sortSlice(
					parseStringToKV(input)), k),
		),
	)
}

func printResultSlice(sliceResult []string) {
	println(strings.Join(sliceResult, " "))
}

func getResultFromKV(inputKV []kv) []string {
	sliceResult := make([]string, 0, len(inputKV))
	for _, v := range inputKV {
		sliceResult = append(sliceResult, v.Key)
	}
	return sliceResult
}

func limitKV(inputKV []kv, k int) []kv {
	var kVLimit []kv
	if len(inputKV) >= k {
		kVLimit = inputKV[:k]
	} else {
		kVLimit = inputKV
	}
	return kVLimit
}

func parseStringToKV(input string) []kv {
	arrFields := strings.Fields(input)
	m := make(map[string]int)
	for i := 0; i < len(arrFields); i++ {
		m[arrFields[i]]++
	}
	inputKV := make([]kv, 0, len(m))
	for k, v := range m {
		inputKV = append(inputKV, kv{k, v})
	}
	return inputKV
}

func sortSlice(inputKV []kv) []kv {
	sort.Slice(inputKV, func(i, j int) bool {
		if inputKV[i].Value == inputKV[j].Value {
			return inputKV[i].Key < inputKV[j].Key
		}
		return inputKV[i].Value > inputKV[j].Value
	})
	return inputKV
}

func getWordsFromTerminal() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	errorHandler(err)
	return input
}

func getK() int {
	var k int
	_, err := fmt.Scanf("%d", &k)
	errorHandler(err)
	return k
}

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("invalid input\ncause: %s\n", err.Error())
		run()
	}
}
