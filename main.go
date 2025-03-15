package main

import (
	"fmt"
	"regexp"
	"os"
	"strconv"
)

var filename string = "/sys/class/backlight/nv_backlight/brightness"

func GetBrightness() string {

	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	brightnessString := string(data)

	re := regexp.MustCompile("[^0-9]+")
	result := re.ReplaceAllString(brightnessString, "")

	return result
}
 
func SetBrightness(direction string, value int) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	s := GetBrightness()
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	var data []byte

	switch direction {
		case "inc":
			data = []byte(strconv.Itoa(i+value))
		case "dec":
			data = []byte(strconv.Itoa(i-value))
	}

	_, err = file.Write(data)
	if err != nil {
		file.Close()
		panic(err)
	}

	file.Close()
}

func main() {

	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: golight <inc/dec>")
		os.Exit(1)
	}

	if args[0] == "inc" {
		SetBrightness("inc", 5)
	} else if args[0] == "dec" {
		SetBrightness("dec", 5)
	} else {
		fmt.Println("Usage: golight <inc/dec>")
		os.Exit(1)
	}

	os.Exit(0)
}
