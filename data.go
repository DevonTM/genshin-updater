package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type GenshinData struct {
	Data struct {
		Game    *GameData `json:"game"`
		PreGame *GameData `json:"pre_download_game"`
	} `json:"data"`
}

type GameData struct {
	Latest struct {
		Ver string `json:"version"`
	} `json:"latest"`
	Diffs []*Diff `json:"diffs"`
}

type Diff struct {
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
}

func getData() (*GenshinData, error) {
	URL := "https://hk4e-launcher-static.hoyoverse.com/hk4e_global/mdk/launcher/api/resource?channel_id=1&key=gcStgarh&launcher_id=10&sub_channel_id=0"
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, URL, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data *GenshinData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
