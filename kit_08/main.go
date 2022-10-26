package main

import (
	"fmt"
	"github.com/qinxiaogit/go_many_version/kit_08/protocol"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	protocolText := protocol.Text{
	}
	protocolText.Init("769af463a39f077a0340a189e9c1ec28","kuaiyin-lz","1BA09530-F9E6-478D-9965-7EB31A59537E")

	rpcClient := Client{
		Class: "Api",
		Protocol:&protocolText,
	}
	rpcClient.init()


	//"robot_result !=\":0,\"table_source\":\"music_audit_result\"},1,10]}
	rpcClient.Call("getList",[]interface{}{
		map[string]interface{}{"robot_result!=":0,"table_source":"music_audit_result"},
		1,
		10,
	})
	//buf := new(bytes.Buffer)
	//byteOrder := binary.LittleEndian
	//
	//binary.Write(buf, byteOrder, uint32(92301))
	//fmt.Printf("uint32: %x\n", buf.Bytes())
	//
	//buf.Reset()
	//binary.Write(buf, byteOrder, uint16(65535))
	//fmt.Printf("uint16: %x\n", buf.Bytes())
	//
	//buf.Reset()
	//binary.Write(buf, byteOrder, float32(0.0012))
	//fmt.Printf("float: %x\n", buf.Bytes())

	//http.HandleFunc("/_stack",getStackTraceHandler)

	runtime.SetMutexProfileFraction(1)
	runtime.SetBlockProfileRate(1)

	go Loop()

	_ = http.ListenAndServe("127.0.0.1:6060", nil)
}

func  Loop()  {
	for i:=0 ;i<10000000000;i++{
		time.Sleep(time.Microsecond*100)
		go func() {
			fmt.Println(runtime.GOROOT())
			time.Sleep(time.Microsecond*1000)
		}()
	}
}

//func getStackTraceHandler(w http.ResponseWriter,r *http.Request){
//	stack := debug.Stack()
//	w.Write(stack)
//	pprof.Lookup("goroutine").WriteTo(w,2)
//}
