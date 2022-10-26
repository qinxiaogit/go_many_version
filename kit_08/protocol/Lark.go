package protocol

type Lark struct {
}

//func (L *Lark)Encode(tag string, value map[string]interface{}):


/**
 * 编码数据流.
 *
 * @param string $tag   请求标识.
 * @param array  $value 数据
 *
 * @return string
 */
//public static function encode($tag, $values = array()) {
//$buffer = '';
//foreach($values as $val) {
//$buffer .= pack('Na*', strlen($val), $val);
//}
//
//// 最后要计算buffer的总长度，便于接收方一次性将buffer全部读取出来，以空间换效率
//return pack('nNa*', $tag, strlen($buffer), $buffer);
//}