package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
)

func NewDoveClient(address string) *doveClient {
	return &doveClient{
		address: address,
	}
}

type doveClient struct {
	address string
}

func (dc *doveClient)Call(methodName string,args map[string]interface{}) (string,[]byte,error) {
	var clientData = make(map[string]interface{})

	clientData["args"] = args
	clientData["cmd"]  = methodName

	conn ,err := net.Dial("unix",dc.address)
	fmt.Println("2222",err)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	clientDataJs ,err := json.Marshal(clientData)
	fmt.Println("2222",err,string(clientDataJs),len(clientDataJs))


	lenStr := fmt.Sprintf("%08d",len(clientDataJs))

	fmt.Println(lenStr)
	conn.Write([]byte(lenStr))

	conn.Write(clientDataJs)

	b, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("333",string(b))
	return "ok",b,err

}

func (dc *doveClient)Close()  {

}

func main()  {
	dc := NewDoveClient("/var/lib/doveclient/doveclient.sock")
	status, result, err := dc.Call("GetEtcdAddr", map[string]interface{}{})
	fmt.Println(status,string(result),err)
}