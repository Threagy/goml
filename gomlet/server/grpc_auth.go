package server

import (
	"context"
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/goml/gomlet/db"
	pbAuth "github.com/goml/gomlet/proto/v1/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	privKeyPath = "rsa-key/jwtRS256.key"
	pubKeyPath  = "rsa-key/jwtRS256.key.pub"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

// AuthServer ...
type AuthServer struct {
	GrpcServer *grpc.Server
}

// NewAuthServer ...
func NewAuthServer() *AuthServer {
	// create server
	grpcServer := grpc.NewServer()
	server := &AuthServer{GrpcServer: grpcServer}
	pbAuth.RegisterAuthServer(grpcServer, server)

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}

	return server
}

// NewUser ...
func (srv *AuthServer) NewUser(ctx context.Context, req *pbAuth.NewUserRequest) (*pbAuth.NewUserResponse, error) {
	user := db.UserModel{
		UserID:   req.UserId,
		Password: req.Password,
		Name:     req.Name,
		Role:     db.Role(req.Role),
	}

	err := db.NewUser(user)
	if err != nil {
		return nil, err
	}

	response := &pbAuth.NewUserResponse{}

	return response, nil
}

// Login ...
func (srv *AuthServer) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	user, err := db.GetUser(req.UserId, req.Password)
	if err != nil {
		return nil, err
	}

	user.Password = "removed"

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = &AuthClaims{
		&jwt.StandardClaims{
			// set the expire time: 1 year
			// see http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20#section-4.1.4
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
		},
		"level1",
		*user,
	}
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, err.Error())
	}

	response := &pbAuth.LoginResponse{
		Token: tokenString,
	}

	return response, nil
}
