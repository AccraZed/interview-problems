package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	ClassChangeURL = "https://qwerasd205.github.io/DiscordClassChanges/differences.csv"
	TargetURL      = "/mnt/f/Documents/BetterDiscord/YoRHA-UI-BetterDiscord/src"
	TmpFilename    = "replace.txt"
)

func main() {
	if err := downloadFile(TmpFilename, ClassChangeURL); err != nil {
		os.Remove(TmpFilename)
		panic(err)
	}

	f, err := os.Open(TmpFilename)
	if err != nil {
		os.Remove(TmpFilename)
		panic(err)
	}
	defer f.Close()

	legend := make(map[string]string)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		names := strings.Split(s.Text(), ",")

		// Set all names in comma list to last name (newest)
		for i, name := range names {
			if i == len(names)-1 {
				continue
			}
			legend[name] = names[len(names)-1]
		}
	}

	errChan := make(chan error)
	fileCount := 0
	filepath.Walk(TargetURL, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fileCount++
		go FindAndReplace(path, legend, errChan)
		return nil
	})

	for i := 0; i < fileCount; i++ {
		err := <-errChan
		if err != nil {
			os.Remove(TmpFilename)
			panic(err)
		}
	}

	os.Remove(TmpFilename)
	fmt.Println("Successfully modified all files! Closing...")
}

func FindAndReplace(path string, legend map[string]string, failChan chan error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		failChan <- err
		return
	}

	names := string(dat)
	for k, v := range legend {
		names = strings.ReplaceAll(names, k, v)
	}

	f, err := os.Create(path)
	if err != nil {
		failChan <- err
		return
	}
	defer f.Close()
	f.WriteString(names)

	fmt.Printf("finished %s\n", path)
	failChan <- nil
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
