package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func getChoice() int {
	choice := -1
	fmt.Print("\nChoice : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Sscan(input, &choice)
	return choice
}

func convertSize(s string) string {
	size, _ := strconv.ParseInt(s, 10, 64)
	switch {
	case size > GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size > MB:
		return fmt.Sprintf("%d MB", size/MB)
	case size > KB:
		return fmt.Sprintf("%d KB", size/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func convertLang(lang string) string {
	switch lang {
	case "zh-cn":
		return "Chinese"
	case "en-us":
		return "English"
	case "ja-jp":
		return "Japanese"
	case "ko-kr":
		return "Korean"
	default:
		return lang
	}
}

func logError(err error) {
	fmt.Println("Error:", err)
	fmt.Print("\nPress enter to exit . . . ")
	fmt.Scanln()
	os.Exit(1)
}

func clearScreen() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
