package protocol

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

type Text struct {
	rpcSecret string
	user      string
	secret    string
}

const TextVersion = "2.0"

func (text *Text) Init(rpcSecret, user, secret string) {
	text.rpcSecret = rpcSecret
	text.user = user
	text.secret = secret
}

func (text *Text) GetPacket(class, method string, params []interface{}) map[string]interface{} {

	//fmt.Println(text.user,text.rpcSecret)
	data := make(map[string]interface{})
	data["version"] = TextVersion
	data["user"] = text.user
	md5Str := text.user + ":" + text.secret
	data["password"] = fmt.Sprintf("%x", md5.Sum([]byte(md5Str)))
	data["class"] = "RpcClient_" + class
	data["method"] = method
	data["params"] = params

	//timestamp := time.Now().UnixNano()/100000
	//fmt.Println(timestamp)
	data["timestamp"] = 1652350865.135583 //strconv.Itoa(int(timestamp/1000)) + "."+strconv.Itoa(int(timestamp%10000))

	packet := make(map[string]interface{})
	jsData, err := json.Marshal(data)
	//fmt.Println("err:",string(jsData),err)

	jsStr := fmt.Sprintf("%s", jsData)
	packet["data"] = jsStr
	if err != nil {
		fmt.Println("无法序列化")
	}
	packet["signature"] = text.encrypt(jsStr)

	return packet
}

func (text *Text) encrypt(data string) string {
	md5Str := data + "&" + text.rpcSecret
	return text.md5(md5Str)
}

func (text *Text)md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// Encode /**
func (text *Text) Encode(param map[string]interface{}) string {
	command := "RPC"
	jsParam, err := json.Marshal(param)
	if err != nil {
		fmt.Println("无法序列化")
	}
	return fmt.Sprintf("%d\n%s\n%d\n%s\n", len(command), command, len(string(jsParam)), string(jsParam))
}

// Decode 解码/**
func (text Text) Decode(data string) string {
	pos := strings.Index(data, "\n")
	return data[pos+1:]
}

/*
public function decode($data) {
return substr($data, strpos($data, "\n") + 1, -1);
}
*
*/
