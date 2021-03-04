package utils

import (
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var searchChanNum = 0
var searchList = make(map[int]string)
var searchChanDone = make(chan bool)
var searchChanRequest = make(chan map[string]string)
var searchChanRespond = make(chan string)

var searchChanCount = 0

// SearchResult is search file funciton result struct
type SearchResult struct {
	Data     map[int]string
	Num      int
	TakeTime time.Duration
}

// Search file
func Search(searchDir string, search string) (resultData SearchResult) {
	startTime := time.Now()
	go runSearch(searchDir, search, true)
	waitWorker()
	resultData.Data = searchList
	resultData.Num = len(resultData.Data)
	resultData.TakeTime = time.Since(startTime)
	return resultData
}

func waitWorker() {
	for {
		select {
		case searchArgs := <-searchChanRequest:
			searchChanNum++
			go runSearch(searchArgs["dir"], searchArgs["search"], true)
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

func runSearch(searchDir string, search string, master bool) {
	files, errDir := ioutil.ReadDir(searchDir)
	if errDir == nil {
		checkFolder(files, searchDir, search, master)
	}
}

func checkFolder(files []fs.FileInfo, searchDir string, search string, master bool) {
	for _, file := range files {
		if strings.Contains(file.Name(), search) {
			searchChanRespond <- searchDir + file.Name()
		}
		if file.IsDir() {
			newDir := searchDir + file.Name() + string(os.PathSeparator)
			if searchChanNum < maxWorkerCount {
				tmpMap := map[string]string{
					"dir":    newDir,
					"search": search,
				}
				searchChanRequest <- tmpMap
			} else {
				runSearch(newDir, search, false)
			}
		}
	}
	if master {
		searchChanDone <- true
	}
}
