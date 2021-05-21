package main

import (
	"context"
	"flag"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/examples/addsvc/pkg/addtransport"
	"github.com/go-kit/kit/sd"
	"github.com/qinxiaogit/go_many_version/kit_03/addService"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"io"
	"os"
	"time"
	consulsd "github.com/go-kit/kit/sd/consul"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
)

func main(){
	var (
		httpAddr = flag.String("http.addr",":8080","address for http (json) server")
		consulAddr   = flag.String("consul.addr", "", "Consul agent address")
		retryMax     = flag.Int("retry.max", 3, "per-request retries to different instances")
		retryTimeout = flag.Duration("retry.timeout", 500*time.Millisecond, "per-request timeout, including retries")
	)

	flag.Parse()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger,"ts",log.DefaultTimestampUTC)
		logger = log.With(logger,"caller",log.DefaultCaller)
	}

	var client consulsd.Client
	{
		consulConfig := api.DefaultConfig()
		if len(*consulAddr)>0{
			consulConfig.Address = *consulAddr
		}
		consulClient ,err := api.NewClient(consulConfig)
		if err != nil{
			logger.Log("err",err)
			os.Exit(1)
		}
		client = consulsd.NewClient(consulClient)
	}

	//transport domain
	tracer := stdopentracing.GlobalTracer()
	zipkinTracer, _ := stdzipkin.NewTracer(nil, stdzipkin.WithNoopTracer(true))
	ctx := context.Background()
	r := mux.NewRouter()

	{
		var (
			tags = []string{}
			passingOnly = true
			endpoints = addendpoint.Set{}
			instancer = consulsd.NewInstancer(client,logger,"addsvc",tags,passingOnly)
		)

		{
			factory :=  addsvcfaFactory(addendpoint.MakeConcatEndpoint)
			endpointer := sd.NewEndpointer(instancer, factory, logger)

		}
	}
}

func addsvcfaFactory(makeEndpoint func(service addService.Service)endpoint.Endpoint,tracer stdopentracing.Tracer,zipkinTracer *stdzipkin.Tracer,logger log.Logger)sd.Factory{
	return func(instance string) (endpoint.Endpoint,io.Closer,error){
		conn,err := grpc.Dial(instance,grpc.WithInsecure())
		if err != nil{
			return nil,nil,err
		}
		service := addtransport.NewGRPCClient()
	}
}
