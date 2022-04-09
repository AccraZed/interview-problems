package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const CLASS_CHANGE_URL = "https://qwerasd205.github.io/DiscordClassChanges/differences.csv"
const TARGET_URL = "/mnt/f/Documents/BetterDiscord/YoRHA-UI-BetterDiscord/src"

func main() {
	downloadFile("replace.txt", CLASS_CHANGE_URL)

	f, err := os.Open("replace.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	legend := make(map[string]string)
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		names := strings.Split(s.Text(), ",")

		for i, name := range names {
			if i == len(names)-1 {
				continue
			}
			legend[name] = names[len(names)-1]
		}
	}

	target := TARGET_URL
	filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		go FindAndReplace(path, legend)
		return nil
	})

	fmt.Println("Sleeping...")
	time.Sleep(30 * time.Minute)
}

func FindAndReplace(path string, legend map[string]string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
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
	return err
}
