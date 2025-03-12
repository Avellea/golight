package main

import (
	"fmt"
	"regexp"
	"os"
	"strconv"
//	"github.com/codegoalie/golibnotify"
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

/*
func SendNotification(title string, body string, icon string) {

	notifier := golibnotify.NewSimpleNotifier("golight")
	err := notifier.Update(title, body, icon)

	if err != nil {
		err = fmt.Errorf("Failed to send notification: %w", err)
		fmt.Println(err)
	}
}

 
func SetBrightness(value, int) {
	
}
*/

func IncBrightness(value int) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	s := GetBrightness()
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	data := []byte(strconv.Itoa(i+value))

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	
	file.Close()
}


func DecBrightness(value int) {

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	
	s := GetBrightness()
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	
	data := []byte(strconv.Itoa(i-value))

	_, err = file.Write(data)
	if err != nil {
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
		IncBrightness(5)
	} else if args[0] == "dec" {
		DecBrightness(5)
	} else {
		fmt.Println("Usage: golight <inc/dec>")
		os.Exit(1)
	}

	os.Exit(0)
}
