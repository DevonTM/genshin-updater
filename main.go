package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB

	URL = "https://sdk-os-static.mihoyo.com/hk4e_global/mdk/launcher/api/resource?channel_id=1&key=gcStgarh&launcher_id=10&sub_channel_id=0"
)

type GenshinData struct {
	Data struct {
		Game struct {
			Latest struct {
				Version string `json:"version"`
			}
			Diffs []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
				Path    string `json:"path"`
				Size    string `json:"package_size"`
				Hash    string `json:"md5"`
				Voice   []struct {
					Language string `json:"language"`
					Name     string `json:"name"`
					Path     string `json:"path"`
					Size     string `json:"package_size"`
					Hash     string `json:"md5"`
				} `json:"voice_packs"`
			} `json:"diffs"`
		} `json:"game"`
	} `json:"data"`
}

func main() {
	// create http client with 10s timeout
	client := &http.Client{
		Timeout: 10e9,
	}

	// create http request
	req, err := http.NewRequest("GET", URL, http.NoBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// send http request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// parse http response
	var data GenshinData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// extract the required values
	oldVer := data.Data.Game.Diffs[0].Version
	newVer := data.Data.Game.Latest.Version
	gameData := data.Data.Game.Diffs[0]

	// print available updates
	fmt.Printf("Genshin Impact Patch Downloader\nUpdate Version : %s to %s\n\n", oldVer, newVer)
	fmt.Printf("1. game update (%s)\n", convertSize(gameData.Size))
	for i, voiceData := range gameData.Voice {
		fmt.Printf("%d. %s voice (%s)\n", i+2, voiceData.Language, convertSize(voiceData.Size))
	}

	var name, path, hash string

	// ask user to select update
	for {
		var choice int
		fmt.Print("\nSelect Update : ")
		fmt.Scanln(&choice)
		if choice == 1 {
			name = gameData.Name
			path = gameData.Path
			hash = gameData.Hash
		} else if choice > 1 && choice < len(gameData.Voice)+2 {
			name = gameData.Voice[choice-2].Name
			path = gameData.Voice[choice-2].Path
			hash = gameData.Voice[choice-2].Hash
		} else {
			fmt.Println("Wrong Choice")
			continue
		}
		break
	}

	// download update file
	fmt.Println("Downloading:", name)
	command := fmt.Sprintf("aria2c.exe --conf-path=aria2.conf --checksum=md5=%s %s", hash, path)
	cmd := exec.Command("cmd", "/C", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("\npress enter to exit...")
	fmt.Scanln()
}

// convert size from string bytes to human readable format
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
