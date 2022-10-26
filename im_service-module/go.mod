module github.com/GoBelieveIO/im_service

go 1.12

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/gomodule/redigo v1.8.1
	github.com/gorilla/websocket v1.4.2
	github.com/importcjj/sensitive v0.0.0-20190611120559-289e87ec4108
	github.com/kr/pretty v0.1.0 // indirect
	github.com/richmonkey/cfg v0.0.0-20130815005846-4b1e3c1869d4
	github.com/valyala/gorpc v0.0.0-20160519171614-908281bef774
)

replace github.com/GoBelieveIO/im_service/lru => ./lru
