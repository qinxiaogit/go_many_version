package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/otiai10/gosseract/v2"
	"io/ioutil"
)

var path string
var data = []string{"a", "string", "list"}

func main() {

	client := gosseract.NewClient()
	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	err := client.SetImage("/Users/owlet/Downloads/配置平台.png")
	if err != nil {
		return 
	}
	text, _ := client.Text()
	fmt.Println(text)

	fmt.Println("-----------------------------------------")

	root := "/Users/owlet/Downloads/"
	Init(root, fileChangeCallBack)
	/*

		path := "/Users/owlet/Downloads/录音"

		flag.String("","","配置文件地址")

		recursionDir(path,transcoding)
	*/

	//err := ffmpeg.Input("./sample_data/in1.mp4").
	//	Output("./sample_data/out1.mp4", ffmpeg.KwArgs{"c:v": "libx265"}).
	//	OverWriteOutput().ErrorToStdOut().Run()
	//fmt.Println(err)
	time.Sleep(1000*time.Hour)
}

func transcoding(filename string) {
	filenameArr := strings.Split(filename, ".")
	ctFA := len(filenameArr)
	if ctFA < 2 {
		return
	}
	//ext := filenameArr[ctFA-1]

	////ffmpeg -i "/Users/owlet/Downloads/录音/$filename/$file" -f mp3 "/Users/owlet/Downloads/录音/$filename/${arr[0]}.mp3"
	//o := exec.Command("ffmpeg" ,"-i",filename,"-f","mp3", "new.mp3")
	//err := o.Run()
	//fmt.Println(err.Error())
}

func recursionDir(path string, callback func(filename string)) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, f := range files {
		if f.IsDir() {
			recursionDir(path+"/"+f.Name(), callback)
			continue
		}
		//fmt.Println(path+"/"+f.Name())
		callback(path + "/" + f.Name())
	}
}
