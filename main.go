package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var pathToScan = "./src"

func main() {
	if len(os.Args) == 2 {
		pathToScan = os.Args[1]
	}
	fmt.Printf("Scanning dir [%s]\n", pathToScan)
	entries := readDir(pathToScan)
	for _, registry := range entries {
		registryPath := filepath.Join(pathToScan, registry.Name())
		users := readDir(registryPath)
		for _, module := range users {
			modulePath := filepath.Join(registryPath, module.Name())
			modules := readDir(modulePath)
			printModules(modulePath, modules)
		}
	}
}

func readDir(dirPath string) []fs.DirEntry {
	myEntries := make([]fs.DirEntry, 0)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, repo := range entries {
		if !repo.IsDir() {
			continue
		}
		if strings.HasPrefix(repo.Name(), ".") {
			continue
		}
		myEntries = append(myEntries, repo)
	}
	return myEntries
}

func printModules(registryPath string, modules []fs.DirEntry) {
	for _, e := range modules {

		modulePath := filepath.Join(registryPath, e.Name())
		modulePath = strings.TrimPrefix(modulePath, "./")
		modulePath = strings.TrimPrefix(modulePath, pathToScan)
		modulePath = strings.TrimPrefix(modulePath, "/")
		fmt.Println(modulePath)
	}
}
