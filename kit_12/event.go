package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
)

/**
 * 文件变更事件
 */
func fileChangeCallBack(event fsnotify.Event){
	//fileName := event.Name
	//filepath.Abs()
		fmt.Println(event.Op.String())
}
