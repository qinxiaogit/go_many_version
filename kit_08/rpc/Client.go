package rpc

import (
	"fmt"
	"github.com/qinxiaogit/go_many_version/kit_08/pool"
	"github.com/qinxiaogit/go_many_version/kit_08/protocol"
	"runtime"
)

type Client struct {
	ServiceName string
	Class       string
	Protocol    protocol.Protocol
	//Con         Tcp
}

func (c *Client) init() {
	//c.Con = Tcp{}
	//c.Con.OpenServerConnection()
	//c.Protocol = &protocol.Text{}
}

func (c *Client) Call(method string, arguments []interface{}) {
	content := make(map[string]interface{})
	content["class"] = c.Class
	content["method"] = method
	content["arguments"] = arguments

	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Println(f.Name())

	c.Con.Send(c.initRpcData(c.Class, method, arguments))
	c.Con.Read()
}

/**
 *
 */
func (c *Client) initRpcData(class, method string, arguments []interface{}) string {
	pack := c.Protocol.GetPacket(class, method, arguments)
	return c.Protocol.Encode(pack)
}

func (c *Client) send() {

}
