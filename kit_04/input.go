package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

type Input struct {

}

func(i Input) Run(){
	client := resty.New()

	s := `{"consignee":{"address":"四川省成都市龙泉驿区金科东方雅郡","addressId":"896174308634976256","mobile":"18227755589","name":"覃枭"},"coupons":[],"isFromCart":true,"isDeleteCart":false,"isUseAllPromotion":false,"isUseCoupon":false,"orderType":"NORMAL","products":[{"productCode":"1111111010","productId":"798812608298868736","amount":1,"deliveryType":"THRID_TRANSPORT","remark":null,"promotionCode":null,"couponCode":null,"grouponRecordId":null,"productType":"CONDITION","productName":"茅台 53度酱香型新飞天茅台酒瓶装500ml","logo":"commodity/product/2021-06-02/product1622619337801.jpg","finalPrice":"14990000","marketPrice":"14990000","usePoint":null,"freight":null,"promotionPrice":null,"couponPrice":null,"inToDetailChannel":"NORMAL"}],"promotionCode":null,"authorizationCode":"011Uxf000pQHOL1FMv100YHmbM1Uxf0q"}`
	result := gjson.Parse(s)
	url := "https://prod.ggszhg.com/xgt-app/applet/order/orderSettlement?os=APPLET&osVersion=1.0.0&userId=896164256616378368&userToken=4e5ddf91a42a4adfb834232f0ea91296&sign=5B26E34E0745717547FA42E1DB3C3D58"
	resp, err := client.SetProxy("http://127.0.0.1:7890").R().
		EnableTrace().
		SetHeader("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 14_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.6(0x1800062f) NetType/4G Language/zh_CN").
		SetBody(result.Value()).Post(url)

	fmt.Println(err,string(resp.String()))
}
