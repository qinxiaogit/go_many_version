package main

import (
	"context"
	"go.uber.org/zap"
)

type OutPut struct {
	
}

func(o OutPut)Run(){
	result, err := redisClient.HGetAll(context.Background(), "test_key").Result()
	if err != nil {
		return
	}
	recover()

	for it,v:= range result{
		NewLog().Log.Info("del hash key :",zap.String("it",it),zap.String("attt",v))
		redisClient.HDel(context.Background(),"test_key",it)
	}
}
