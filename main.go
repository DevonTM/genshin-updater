package main

import (
	"fmt"
	"time"
)

const VERSION = "v1.3.1"

func main() {
	fmt.Println("Genshin Impact Patch Downloader " + VERSION)
	fmt.Print("Fetching data . . . ")
	data, err := getData()
	if err != nil {
		logError(err)
	}
	time.Sleep(2 * time.Second)

	dlData := selectPatch(data)
	err = getFile(dlData)
	if err != nil {
		logError(err)
	}

	fmt.Print("\nPress enter to exit . . . ")
	fmt.Scanln()
}
