package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type visit struct {
	date   string
	doctor string
}

var usersVisitStorage map[string][]visit

func main() {
	initialisationUserStorage()
	run()
}

func initialisationUserStorage() {
	usersVisitStorage = make(map[string][]visit)
}

func run() {
	for {
		fmt.Printf("enter operation (Save, GetHistory or GetLastVisit): ")
		switch getInputStringFromTerminal(true) {
		case "Save":
			save()
		case "GetHistory":
			getHistory()
		case "GetLastVisit":
			getLastVisit()
		case "Exit":
			os.Exit(0)
		default:
			errorHandler(errors.New("operation not supported"))
		}
	}
}

func save() {
	fmt.Printf("enter name: ")
	patientName := getInputStringFromTerminal(false)
	var currentVisit visit
	fmt.Printf("enter doctor: ")
	currentVisit.doctor = getInputStringFromTerminal(false)
	fmt.Printf("enter date (format YYYY-MM-DD): ")
	currentVisit.date = getDate()
	visits := usersVisitStorage[patientName]
	usersVisitStorage[patientName] = append(visits, currentVisit)
}

func getHistory() {
	fmt.Printf("enter patient name: ")
	patientName := getInputStringFromTerminal(false)
	v, ok := usersVisitStorage[patientName]
	if ok == false {
		errorHandler(errors.New("user not found"))
		return
	}
	for _, v := range v {
		fmt.Printf("%s %s\n", v.date, v.doctor)
	}
}

func getLastVisit() {
	fmt.Printf("enter patient name: ")
	patientName := getInputStringFromTerminal(false)
	fmt.Printf("enter doctor: ")
	doctor := getInputStringFromTerminal(false)
	v, ok := usersVisitStorage[patientName]
	if ok == false {
		errorHandler(errors.New("user not found"))
		return
	}
	for i := len(v) - 1; i >= 0; i-- {
		if strings.EqualFold(v[i].doctor, doctor) {
			fmt.Printf("%s %s\n", v[i].date, v[i].doctor)
			return
		}
	}
	errorHandler(errors.New("visit not found"))
}

func getDate() string {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	for {
		date := getInputStringFromTerminal(true)
		if re.MatchString(date) {
			return date
		} else {
			errorHandler(errors.New("invalid date format"))
		}
	}
}

func getInputStringFromTerminal(nullable bool) string {
	reader := bufio.NewReader(os.Stdin)
	var input string
	for {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if !nullable && strings.EqualFold(input, "") {
			errorHandler(errors.New("input cannot be blank"))
		} else {
			break
		}
	}
	return input
}

func errorHandler(err error) {
	if err != nil {
		fmt.Printf("invalid input\ncause: %s\n", err.Error())
	}
}
