package utils

import "runtime"

const (
	// UTypeFile is file
	UTypeFile string = "file"
	// UTypeFolder is folder
	UTypeFolder string = "folder"
)

var maxWorkerCount = runtime.NumCPU()
