package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

type DownloadData struct {
	Name string
	URL  string
	Hash string
}

func getFile(data *DownloadData) error {
	clearScreen()
	fmt.Println("Downloading :", data.Name)

	cmd := exec.Command(aria2cPath, "--conf-path", aria2confPath, "--checksum", data.Hash, "--", data.URL)
	cmd.Dir = execDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cancel := make(chan os.Signal, 1)
	signal.Notify(cancel, os.Interrupt)
	go func() {
		for {
			<-cancel
			if err := cmd.Process.Signal(os.Interrupt); err == nil {
				return
			}
		}
	}()

	err := cmd.Run()
	return err
}
