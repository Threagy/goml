package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/goml/gomlet/server/driver"
	"github.com/google/wire"
	"golang.org/x/net/websocket"
)

// Set is a Wire provider set that produces a *Server given the fields of
// Options.
var Set = wire.NewSet(
	NewServer,
	wire.Struct(new(Options), "Driver", "GomlapiServer", "GomletServer", "UserServer", "AuthServer"),
	NewDefaultDriver,
	NewGomletServer,
	NewGomlapiServer,
	NewUserServer,
	NewAuthServer,
	wire.Bind(new(driver.Server), new(*DefaultDriver)),
)

// Server is a preconfigured HTTP server with diagnostic hooks.
// The zero value is a server with the default options.
type Server struct {
	handler       http.Handler
	once          sync.Once
	driver        driver.Server
	gomletServer  *GomletServer
	gomlapiServer *GomlapiServer
	userServer    *UserServer
	authServer    *AuthServer
}

// Options is the set of optional parameters.
type Options struct {
	// Driver serves HTTP requests.
	Driver driver.Server

	// GomlapiServer serves GRPC requests.
	GomlapiServer *GomlapiServer

	// GolmetServer serves GRPC requests.
	GomletServer *GomletServer

	// UserServer serves GRPC requests.
	UserServer *UserServer

	// AuthServer serves GRPC requests.
	AuthServer *AuthServer
}

// NewServer creates a new server. New(nil, nil) is the same as new(Server).
func NewServer(h http.Handler, opts *Options) *Server {
	srv := &Server{handler: h}
	if opts != nil {
		srv.driver = opts.Driver
		srv.gomletServer = opts.GomletServer
		srv.gomlapiServer = opts.GomlapiServer
		srv.userServer = opts.UserServer
		srv.authServer = opts.AuthServer
	}
	return srv
}

func (srv *Server) init() {
	srv.once.Do(func() {
		if srv.handler == nil {
			srv.handler = http.DefaultServeMux
		}

		if srv.driver == nil {
			srv.driver = NewDefaultDriver()
		}
	})
}

// ListenAndServe is a wrapper to use wherever http.ListenAndServe is used.
func (srv *Server) ListenAndServe(addr string) error {
	srv.init()

	mux := http.NewServeMux()

	grpcMux := newGrpcMux(addr)
	mux.Handle("/v1/", grpcMux)
	mux.Handle("/v1/exec/", websocket.Handler(execContainer))
	mux.HandleFunc("/v1/exec/resizeTty", execResizeTty)

	handler := grpcHandlerFunc(
		srv.authServer.GrpcServer,
		srv.userServer.GrpcServer,
		srv.gomlapiServer.GrpcServer,
		srv.gomletServer.GrpcServer, mux)

	return srv.driver.ListenAndServe(addr, handler)
}

// Shutdown gracefully shuts down the server without interrupting any active connections.
func (srv *Server) Shutdown(ctx context.Context) error {
	if srv.driver == nil {
		return nil
	}
	return srv.driver.Shutdown(ctx)
}
