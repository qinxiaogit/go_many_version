package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"time"
)

func worker(){

	//cmd := exec.Command("cd","/Users/owlet/code/kuaiyin-lz")
	//
	//_, err := cmd.Output()
	//if err != nil {
	//	fmt.Printf("Execute Shell:%s failed with error:%s", "aa", err.Error())
	//	return
	//}
	//fmt.Printf("Execute Shell:%s finished with output:\n%s", "bbb", string(output))
	//err := os.Chdir("/Users/owlet/code/kuaiyin-lz")
	//
	//fmt.Println("err00:",err)
	//cmd := exec.Command("git","branch","-all")
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//err = cmd.Run()
	//vms_host := out.String()
	//log.Print("vms_host:"+vms_host,err.Error())
	findCmd := cmd.NewCmd("git","branch","-all")
	// Start a long-running process, capture stdout and stderr
	//findCmd := cmd.NewCmd("find", "/", "--name", "needle")
	//statusChan := findCmd.Start() // non-blocking

	ticker := time.NewTicker(2 * time.Second)

	// Print last line of stdout every 2s
	go func() {
		for range ticker.C {
			status := findCmd.Status()
			n := len(status.Stdout)
			fmt.Println(status.Stdout[n-1])
		}
	}()

	// Stop command after 1 hour
	go func() {
		<-time.After(1 * time.Hour)
		findCmd.Stop()
	}()
	time.Sleep(time.Hour)
}
