package audit

import (
	"github.com/qinxiaogit/go_many_version/kit_08/rpc"
)

type ApiServiceClient struct {
	*rpc.Client
}

func (api *ApiServiceClient) GetList(params ...interface{}) {
	api.Call("getList", params)
}
