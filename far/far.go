package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
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

	var wg sync.WaitGroup
	filepath.Walk(TargetURL, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := FindAndReplace(path, legend); err != nil {
				os.Remove(TmpFilename)
				panic(err)
			}
		}()

		return nil
	})

	wg.Wait()

	os.Remove(TmpFilename)
	fmt.Println("Successfully modified all files! Closing...")
}

func FindAndReplace(path string, legend map[string]string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	names := string(dat)
	for k, v := range legend {
		names = strings.ReplaceAll(names, k, v)
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(names)

	fmt.Printf("finished %s\n", path)
	return nil
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
