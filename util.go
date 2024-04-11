package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/inancgumus/screen"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

var (
	execDir       string
	aria2cPath    string
	aria2confPath string
)

func init() {
	exec, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execDir = filepath.Dir(exec)
}

func checkAria2() error {
	aria2c := "aria2c"
	aria2conf := "aria2.conf"

	if runtime.GOOS == "windows" {
		aria2c += ".exe"
	}

	aria2cPath = filepath.Join(execDir, aria2c)
	aria2confPath = filepath.Join(execDir, aria2conf)

	_, err := os.Stat(aria2cPath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		aria2cPath = ""
	}

	if aria2cPath == "" {
		aria2cPath, err = exec.LookPath("aria2c")
		if err != nil {
			return fmt.Errorf("%s not found", aria2c)
		}
	}

	_, err = os.Stat(aria2confPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("%s not found", aria2conf)
		}
		return err
	}

	return nil
}

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
	screen.Clear()
	screen.MoveTopLeft()
}
