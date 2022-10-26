package main

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

func main(){

	url := "https://api.chanify.net/v1/sender/CIDtkIkGEiJBQzdCVFJJQllaWFNBRlozSFRKUERKVjZWTksyM0c0U0lZIgwIAhoIbnVtYmVyMDE.-iToL-I7OYvoNI2tNE-CFzmpLiwbwPsuvqmRKAVHexY"
	post, err := resty.New().R().EnableTrace().SetFormData(map[string]string{"text": "<a href=\"http://www.baidu.com\">hello</a>","sound":"1",
		"title":"小明快跑",
		"action":"debug|http://www.baidu.com",
	}).Post(url)
	if err != nil {
		return
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	sugar.Debug(post)

	post,err = resty.New().R().EnableTrace().SetBody(map[string]interface{}{"text": "<a href=\"http://www.baidu.com\">hello</a>","sound":"1",
		"title":"小明快跑",
		"action":[]interface{}{"第一个按钮|http://www.baidu.com","debug|http://www.baidu.com","第二个按钮|http://www.baidu.com"},
	}).SetHeader("Content-Type", "multipart/form-data").Post(url)


	sugar.Info(post,err)


}