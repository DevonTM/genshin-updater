package main

import (
	"fmt"
	"os"
	"time"
)

func selectPatch(data *GenshinData) (dlData *DownloadData) {
	for dlData == nil {
		clearScreen()
		fmt.Println("Select patch")
		fmt.Println("1. Update")
		fmt.Println("2. Pre-Update")
		fmt.Println("0. Exit")

		choice := getChoice()

		switch choice {
		case 0:
			clearScreen()
			os.Exit(0)
		case 1:
			dlData = selectVer(data.Data.Game)
		case 2:
			dlData = selectVer(data.Data.PreGame)
		default:
			fmt.Println("Wrong choice")
			time.Sleep(1 * time.Second)
		}
	}
	return
}

func selectVer(data *GameData) (dlData *DownloadData) {
	for dlData == nil {
		clearScreen()

		if data == nil {
			fmt.Println("No update available")
			time.Sleep(1 * time.Second)
			return nil
		}

		fmt.Println("Select version")
		for i, d := range data.Diffs {
			fmt.Printf("%d. %s to %s\n", i+1, d.Ver, data.Latest.Ver)
		}
		fmt.Println("0. Back")

		choice := getChoice()

		if choice == 0 {
			return nil
		} else if choice < 1 || choice > len(data.Diffs) {
			fmt.Println("Wrong choice")
			time.Sleep(1 * time.Second)
			continue
		}

		diff := data.Diffs[choice-1]
		dlData = selectDownload(diff, data.Latest.Ver)
	}
	return
}

func selectDownload(diff *Diff, ver string) (dlData *DownloadData) {
	for dlData == nil {
		clearScreen()
		fmt.Printf("Version %s to %s\n", diff.Ver, ver)
		fmt.Printf("1. Game Data (%s)\n", convertSize(diff.Size))
		for i, v := range diff.Voice {
			fmt.Printf("%d. %s Voice (%s)\n", i+2, convertLang(v.Lang), convertSize(v.Size))
		}
		fmt.Println("0. Back")

		choice := getChoice()

		if choice == 0 {
			return nil
		} else if choice < 1 || choice > len(diff.Voice)+1 {
			fmt.Println("Wrong choice")
			time.Sleep(1 * time.Second)
			continue
		}

		var name, path, hash string
		if choice == 1 {
			name = diff.Name
			path = diff.Path
			hash = diff.Hash
		} else if choice > 1 && choice < len(diff.Voice)+2 {
			name = diff.Voice[choice-2].Name
			path = diff.Voice[choice-2].Path
			hash = diff.Voice[choice-2].Hash
		}

		dlData = &DownloadData{
			Name: name,
			URL:  path,
			Hash: "md5=" + hash,
		}
	}
	return
}
