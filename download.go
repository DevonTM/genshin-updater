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

	done := make(chan struct{}, 1)
	cancel := make(chan os.Signal, 1)
	signal.Notify(cancel, os.Interrupt)
	go func() {
		defer signal.Stop(cancel)
		select {
		case <-done:
		case <-cancel:
			cmd.Process.Signal(os.Interrupt)
		}
	}()

	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 7 {
				return nil
			} else {
				fmt.Println()
				return fmt.Errorf("aria2c exited with code %d", exitError.ExitCode())
			}
		}
		return err
	}

	done <- struct{}{}
	return nil
}
