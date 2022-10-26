package dl

import (
	"sync"
	"transcode/m3u8/parse"
)

const (
	tsExt            = ".ts"
	tsFolderName     = "ts"
	mergeTSFilename  = "main.ts"
	tsTempFileSuffix = "_tmp"
	progressWidth    = 40
)


type Downloader struct {
	lock     sync.Mutex
	queue    []int
	folder   string
	tsFolder string
	finish   int32
	segLen   int

	result *parse.Result
}

