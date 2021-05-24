package addtransport

import (
	"context"
	"errors"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/tracing/opentracing"
	"github.com/go-kit/kit/tracing/zipkin"
	"github.com/go-kit/kit/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"github.com/qinxiaogit/go_many_version/kit_03/addService"
	"github.com/qinxiaogit/go_many_version/kit_03/addendpoint"
	"github.com/qinxiaogit/go_many_version/kit_03/pb"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"time"
)

type grpcService struct {
	sum grpctransport.Handler
	concat grpctransport.Handler
}

func (g *grpcService) Sum(ctx context.Context, request *pb.SumRequest) (*pb.SumReply, error) {
	_,rep,err := g.sum.ServeGRPC(ctx,request)
	if err != nil{
		return nil,err
	}
	return rep.(*pb.SumReply),nil
}

func (g *grpcService) Concat(ctx context.Context, request *pb.ConcatRequest) (*pb.ConcatReply, error) {
	_,rep,err := g.concat.ServeGRPC(ctx,request)
	if err !=nil{
		return nil,err
	}
	return rep.(*pb.ConcatReply),nil
}

func NewGrpcServer(endpoints addendpoint.Set,otTracer stdopentracing.Tracer,zipkinTracer *stdzipkin.Tracer,logger log.Logger)pb.AddServer{
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}
	if zipkinTracer != nil{
		options = append(options,zipkin.GRPCServerTrace(zipkinTracer))
	}
	return &grpcService{
		sum: grpctransport.NewServer(
			endpoints.SumEndpoint,
			decodeGRPCSumRequest,
			encodeGRPCSumResponse,
			append(options,grpctransport.ServerBefore(opentracing.GRPCToContext(otTracer,"sum",logger)))...,
			),
	}
}

func NewGrpcClient(conn *grpc.ClientConn,otTracer stdopentracing.Tracer,zipinTracer *stdzipkin.Tracer,logger log.Logger)addService.Service{
	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second),100))

	var options []grpctransport.ClientOption
	if zipinTracer != nil{
		options = append(options,zipkin.GRPCClientTrace(zipinTracer))
	}

	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = grpctransport.NewClient(conn,"pb.Add",
			"sum",
			encodeGRPCSumRequest,
			decodeGRPCSumResponse,
			pb.SumReply{},
			append(options,grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer,logger)))...,
			).Endpoint()

		sumEndpoint = opentracing.TraceClient(otTracer,"sum")(sumEndpoint)
		sumEndpoint = limiter(sumEndpoint)
		sumEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name: "Sum",
			Timeout: 30*time.Second,
		}))(sumEndpoint)
	}

	var concatEndpoint endpoint.Endpoint
	{
		concatEndpoint = grpctransport.NewClient(
			conn,
			"pb.Add",
			"Concat",
			encodeGRPCConcatRequest,
			decodeGRPCConcatResponse,
			pb.ConcatReply{},
			append(options, grpctransport.ClientBefore(opentracing.ContextToGRPC(otTracer, logger)))...,
		).Endpoint()
		concatEndpoint = opentracing.TraceClient(otTracer, "Concat")(concatEndpoint)
		concatEndpoint = limiter(concatEndpoint)
		concatEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:    "Concat",
			Timeout: 10 * time.Second,
		}))(concatEndpoint)
	}
	return addendpoint.Set{
		SumEndpoint:        sumEndpoint,
			ConcatEndpoint: concatEndpoint,
	}
}

// decodeGRPCSumRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC sum request to a user-domain sum request. Primarily useful in a server.
func decodeGRPCSumRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.SumRequest)
	return addendpoint.SumRequest{A: int(req.A), B: int(req.B)}, nil
}

// decodeGRPCConcatRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC concat request to a user-domain concat request. Primarily useful in a
// server.
func decodeGRPCConcatRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.ConcatRequest)
	return addendpoint.ConcatRequest{A: req.A, B: req.B}, nil
}

// decodeGRPCSumResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC sum reply to a user-domain sum response. Primarily useful in a client.
func decodeGRPCSumResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.SumReply)
	return addendpoint.SumResponse{V: int(reply.V), Err: str2err(reply.Err)}, nil
}

// decodeGRPCConcatResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response. Primarily useful in a
// client.
func decodeGRPCConcatResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.ConcatReply)
	return addendpoint.ConcatResponse{V: reply.V, Err: str2err(reply.Err)}, nil
}

// encodeGRPCSumResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain sum response to a gRPC sum reply. Primarily useful in a server.
func encodeGRPCSumResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.SumResponse)
	return &pb.SumReply{V: int64(resp.V), Err: err2str(resp.Err)}, nil
}

// encodeGRPCConcatResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain concat response to a gRPC concat reply. Primarily useful in a
// server.
func encodeGRPCConcatResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(addendpoint.ConcatResponse)
	return &pb.ConcatReply{V: resp.V, Err: err2str(resp.Err)}, nil
}

// encodeGRPCSumRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain sum request to a gRPC sum request. Primarily useful in a client.
func encodeGRPCSumRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(addendpoint.SumRequest)
	return &pb.SumRequest{A: int64(req.A), B: int64(req.B)}, nil
}

// encodeGRPCConcatRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain concat request to a gRPC concat request. Primarily useful in a
// client.
func encodeGRPCConcatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(addendpoint.ConcatRequest)
	return &pb.ConcatRequest{A: req.A, B: req.B}, nil
}


func str2err(s string)error{
	if s == ""{
		return nil
	}
	return errors.New(s)
}

func err2str(err error)string{
	if err == nil{
		return ""
	}
	return err.Error()
}




