package protocol

type Protocol interface {

	GetPacket(class, method string, params []interface{}) map[string]interface{}
	Encode(param map[string]interface{}) string
	Decode(data string) string
}
