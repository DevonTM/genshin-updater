package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

type GenshinData struct {
	Data struct {
		Game    *gameData `json:"game"`
		PreGame *gameData `json:"pre_download_game"`
	} `json:"data"`
}

type gameData struct {
	Latest struct {
		Ver string `json:"version"`
	} `json:"latest"`
	Diff []struct {
		Name  string `json:"name"`
		Ver   string `json:"version"`
		Path  string `json:"path"`
		Size  string `json:"package_size"`
		Hash  string `json:"md5"`
		Voice []struct {
			Lang string `json:"language"`
			Name string `json:"name"`
			Path string `json:"path"`
			Size string `json:"package_size"`
			Hash string `json:"md5"`
		} `json:"voice_packs"`
	} `json:"diffs"`
}

var URL = "https://hk4e-launcher-static.hoyoverse.com/hk4e_global/mdk/launcher/api/resource?channel_id=1&key=gcStgarh&launcher_id=10&sub_channel_id=0"

func main() {
	fmt.Printf("Genshin Impact Patch Downloader\n")

	client := &http.Client{Timeout: 10e9}
	req, err := http.NewRequest("GET", URL, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var data *GenshinData
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	g := data.Data.Game
	p := data.Data.PreGame
	gvL := len(g.Diff[0].Voice)
	pvL := 0

	var preDL bool
	if p != nil {
		pvL = len(p.Diff[0].Voice)
		preDL = true
	}

	fmt.Printf("\nUpdate Version : %s to %s\n", g.Diff[0].Ver, g.Latest.Ver)
	fmt.Printf("1. game update (%s)\n", convertSize(g.Diff[0].Size))
	for i, v := range g.Diff[0].Voice {
		fmt.Printf("%d. %s voice (%s)\n", i+2, v.Lang, convertSize(v.Size))
	}

	if preDL {
		fmt.Printf("\nPre Update Version : %s to %s\n", p.Diff[0].Ver, p.Latest.Ver)
		fmt.Printf("%d. game update (%s)\n", gvL+2, convertSize(p.Diff[0].Size))
		for i, v := range p.Diff[0].Voice {
			fmt.Printf("%d. %s voice (%s)\n", gvL+i+3, v.Lang, convertSize(v.Size))
		}
	}

	var name, path, hash string
	for {
		var choice int
		fmt.Print("\nSelect Update : ")
		fmt.Scanln(&choice)
		if choice < 1 || choice > gvL+pvL+2 {
			fmt.Println("wrong choice")
			continue
		} else if choice == 1 {
			name = g.Diff[0].Name
			path = g.Diff[0].Path
			hash = g.Diff[0].Hash
		} else if choice < gvL+2 {
			name = g.Diff[0].Voice[choice-2].Name
			path = g.Diff[0].Voice[choice-2].Path
			hash = g.Diff[0].Voice[choice-2].Hash
		} else if preDL && choice == gvL+2 {
			name = p.Diff[0].Name
			path = p.Diff[0].Path
			hash = p.Diff[0].Hash
		} else if preDL && choice < gvL+pvL+3 {
			name = p.Diff[0].Voice[choice-gvL-3].Name
			path = p.Diff[0].Voice[choice-gvL-3].Path
			hash = p.Diff[0].Voice[choice-gvL-3].Hash
		}
		break
	}

	fmt.Println("\nDownloading :", name)
	command := fmt.Sprintf("aria2c --conf-path=aria2.conf --checksum=md5=%s -- %s", hash, path)
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

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

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
