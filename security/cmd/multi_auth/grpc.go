package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"sync"

	anothersecuredservice "github.com/goadesign/examples/security/gen/another_secured_service"
	another_secured_servicepb "github.com/goadesign/examples/security/gen/grpc/another_secured_service/pb"
	anothersecuredservicesvr "github.com/goadesign/examples/security/gen/grpc/another_secured_service/server"
	secured_servicepb "github.com/goadesign/examples/security/gen/grpc/secured_service/pb"
	securedservicesvr "github.com/goadesign/examples/security/gen/grpc/secured_service/server"
	securedservice "github.com/goadesign/examples/security/gen/secured_service"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcmdlwr "goa.design/goa/grpc/middleware"
	"goa.design/goa/middleware"
	"google.golang.org/grpc"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, securedServiceEndpoints *securedservice.Endpoints, anotherSecuredServiceEndpoints *anothersecuredservice.Endpoints, wg *sync.WaitGroup, errc chan error, logger *log.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		securedServiceServer        *securedservicesvr.Server
		anotherSecuredServiceServer *anothersecuredservicesvr.Server
	)
	{
		securedServiceServer = securedservicesvr.New(securedServiceEndpoints, nil)
		anotherSecuredServiceServer = anothersecuredservicesvr.New(anotherSecuredServiceEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	secured_servicepb.RegisterSecuredServiceServer(srv, securedServiceServer)
	another_secured_servicepb.RegisterAnotherSecuredServiceServer(srv, anotherSecuredServiceServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Printf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Printf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Printf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
