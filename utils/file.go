package utils

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

var searchChanNum = 0
var searchList = make(map[int]string)
var searchChanDone = make(chan bool)
var searchChanRequest = make(chan map[string]string)
var searchChanRespond = make(chan string)

var searchChanCount = 0

// Search file
func Search(searchType string, searchDir string, search string) {
	startTime := time.Now()
	go runSearch(searchType, searchDir, search, true)
	waitWorker()
	fmt.Println(searchList)
	fmt.Println(time.Since(startTime))
}

func waitWorker() {
	for {
		select {
		case searchArgs := <-searchChanRequest:
			searchChanNum++
			go runSearch(searchArgs["type"], searchArgs["dir"], searchArgs["search"], true)
		case result := <-searchChanRespond:
			searchChanCount++
			if searchList == nil {
				searchList[0] = result
			} else {
				searchList[len(searchList)] = result
			}
		case <-searchChanDone:
			searchChanNum--
			if searchChanNum <= 0 {
				return
			}
		}

	}
}

func runSearch(searchType string, searchDir string, search string, master bool) {
	files, errDir := ioutil.ReadDir(searchDir)
	if errDir == nil {
		switch searchType {
		case UTypeFile:
			// TODO: search file
		case UTypeFolder:
			checkFolder(files, searchDir, search, master)
		}
	}

}

func checkFolder(files []fs.FileInfo, searchDir string, search string, master bool) {
	for _, file := range files {
		if file.IsDir() {
			if file.Name() == search {
				searchChanRespond <- searchDir + file.Name()
			}
			newDir := searchDir + file.Name() + string(os.PathSeparator)
			if searchChanNum < maxWorkerCount {
				tmpMap := map[string]string{
					"type":   UTypeFolder,
					"dir":    newDir,
					"search": search,
				}
				searchChanRequest <- tmpMap
			} else {
				runSearch(UTypeFolder, newDir, search, false)
			}
		}
	}
	if master {
		searchChanDone <- true
	}
}
