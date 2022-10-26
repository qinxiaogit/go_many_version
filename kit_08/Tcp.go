package main

import (
	"bufio"
	"fmt"
	"github.com/qinxiaogit/go_many_version/kit_08/pool"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var RPCCentPool pool.Pool

func init() {
	fmt.Println("init rpcClient")
	var err error
	RPCCentPool, err = pool.NewGenericPool(10, 100, time.Minute, func() (net.Conn, error) {
		conn, err := net.Dial("tcp", "127.0.0.1:2222")
		return conn, err
	})
	if err != nil {
		panic(err)
	}
}

type Tcp struct {
	conn net.Conn
}

func (tcp *Tcp) OpenServerConnection() {
	var err error
	tcp.conn,err = RPCCentPool.Acquire()
	if err!=nil{
		panic(err)
	}

}

func (tcp *Tcp) Send(data string) {
	write, err := tcp.conn.Write([]byte(data))
	fmt.Println(write, err)
	if err != nil {
		return
	}
}

func (tcp *Tcp) Read() interface{} {

	br := bufio.NewReader(tcp.conn)
	l, err := br.ReadBytes('\n')
	if err != nil {
		log.Println("读取数据长失败")
	}
	lInt, err := strconv.Atoi(strings.TrimRight(string(l), "\n"))
	//fmt.Println(lInt, string(l), err)
	by := make([]byte, lInt)
	_, err = br.Read(by)
	//fmt.Println(read,err)
	if err != nil {
		return nil
	}
	fmt.Println(string(by))
	return string(by)
}
