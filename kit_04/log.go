package main

import "go.uber.org/zap"


type log struct {
	Log *zap.Logger
}

func NewLog()*log{
	tmp,_ := zap.NewProduction()
	defer tmp.Sync()
	return &log{
		Log: tmp,
	}
}
