package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type GameData struct {
	Data struct {
		Game struct {
			Latest struct {
				Version string `json:"version"`
			} `json:"latest"`
			Diffs []struct {
				Version     string `json:"version"`
				Path        string `json:"path"`
				MD5         string `json:"md5"`
				VoicePacks  []VoicePack `json:"voice_packs"`
			} `json:"diffs"`
		} `json:"game"`
	} `json:"data"`
}

type VoicePack struct {
	Path string `json:"path"`
	MD5  string `json:"md5"`
}

func main() {
	url := "https://sdk-os-static.mihoyo.com/hk4e_global/mdk/launcher/api/resource?channel_id=1&key=gcStgarh&launcher_id=10&sub_channel_id=0"

	// Fetch the JSON data
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the JSON data
	var data GameData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// Extract the required values
	curVer := data.Data.Game.Diffs[0].Version
	newVer := data.Data.Game.Latest.Version

	fmt.Printf("Genshin Impact Patch Downloader\nupdate version : %s to %s\n\n", curVer, newVer)
	fmt.Println("1. Game Update")
	fmt.Println("2. Chinese Voice")
	fmt.Println("3. English Voice")
	fmt.Println("4. Japanese Voice")
	fmt.Println("5. Korean Voice")

	// Prompt user to select the option
	var option int
	validOption := false
	for !validOption {
		fmt.Print("select an option (1-5): ")
		_, err = fmt.Scanln(&option)
		if err != nil || option < 1 || option > 5 {
			fmt.Println("invalid option selected")
			continue
		}
		
		validOption = true
	}

	// Download the selected file
	path := ""
	md5 := ""
	if option == 1 {
		path = data.Data.Game.Diffs[0].Path
		md5 = data.Data.Game.Diffs[0].MD5
	} else if option >= 2 && option <= 5 {
		index := option - 2
		path = data.Data.Game.Diffs[0].VoicePacks[index].Path
		md5 = data.Data.Game.Diffs[0].VoicePacks[index].MD5
	}

	fmt.Println("\ndownloading...")
	downloadCommand := fmt.Sprintf("aria2c.exe --conf-path=config.txt --checksum=md5=%s %s", md5, path)
	cmd := exec.Command("cmd", "/C", downloadCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\npress enter to exit...")
	fmt.Scanln()
}
