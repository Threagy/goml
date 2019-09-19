package server

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	linq "github.com/ahmetb/go-linq"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes/empty"
	pbCommon "github.com/goml/gomlet/proto/v1/common"
	pbGomlapi "github.com/goml/gomlet/proto/v1/gomlapi"
	pbGomlet "github.com/goml/gomlet/proto/v1/gomlet"
	"github.com/rs/xid"
	"google.golang.org/grpc"
)

var (
	hostname   string
	configFile = "./goml.conf"
	system     *pbGomlapi.System
)

// GomlapiServer ...
type GomlapiServer struct {
	GrpcServer *grpc.Server
}

// NewGomlapiServer ...
func NewGomlapiServer() *GomlapiServer {
	// get hostname
	if h, err := os.Hostname(); err == nil {
		hostname = h
	} else {
		panic(err)
	}

	// load config
	system = &pbGomlapi.System{}
	setState(pbGomlapi.GomlapiState_INITIALIZING)

	if _, err := os.Stat("./goml.conf"); os.IsNotExist(err) {
		file, err := os.Create("./goml.conf")
		defer file.Close()

		system.Config = &pbGomlapi.Config{
			NodeId: xid.New().String(),
			Nodes:  []*pbCommon.NodeSummary{},
			Images: []string{},
		}

		m := &jsonpb.Marshaler{}
		json, err := m.MarshalToString(system.Config)
		log.Print(json)
		if err != nil {
			panic(err)
		}

		w := bufio.NewWriter(file)
		_, err = w.WriteString(json)
		if err != nil {
			panic(err)
		}

		err = w.Flush()
		if err != nil {
			panic(err)
		}
	} else {
		jsonBuffer, err := ioutil.ReadFile("./goml.conf")
		if err != nil {
			panic(err)
		}

		system.Config = &pbGomlapi.Config{}
		m := &jsonpb.Unmarshaler{}
		if err = m.Unmarshal(bytes.NewReader(jsonBuffer), system.Config); err != nil {
			panic(err)
		}
	}

	// refresh docker images
	go refreshDockerImages()

	// create server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ensureValidToken))
	server := &GomlapiServer{GrpcServer: grpcServer}
	pbGomlapi.RegisterGomlapiServer(grpcServer, server)
	return server
}

// SetConfig ...
func (srv *GomlapiServer) SetConfig(ctx context.Context, req *pbGomlapi.SetConfigRequest) (*empty.Empty, error) {
	file, err := os.Create("./goml.conf")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	jsonBuffer, err := json.MarshalIndent(req.Config, "", "\t")
	if err != nil {
		return nil, err
	}

	w := bufio.NewWriter(file)
	_, err = w.WriteString(string(jsonBuffer))
	if err != nil {
		return nil, err
	}

	err = w.Flush()
	if err != nil {
		return nil, err
	}

	system.Config = req.Config

	// refresh docker images
	go refreshDockerImages()

	return &empty.Empty{}, nil
}

// GetSystem ...
func (srv *GomlapiServer) GetSystem(ctx context.Context, req *pbGomlapi.GetSystemRequest) (*pbGomlapi.GetSystemResponse, error) {
	response := &pbGomlapi.GetSystemResponse{
		System: system,
	}

	return response, nil
}

// GetNodeList ... (master)
func (srv *GomlapiServer) GetNodeList(ctx context.Context, req *pbGomlapi.GetNodeListRequest) (*pbGomlapi.GetNodeListResponse, error) {
	response := &pbGomlapi.GetNodeListResponse{}

	var addresses []string
	linq.From(system.Config.Nodes).SelectT(func(n *pbCommon.NodeSummary) string {
		return n.Address
	}).ToSlice(&addresses)

	nodes := newNodes(addresses)

	ch := make(chan *pbCommon.Node, len(addresses))
	for _, node := range nodes {
		go node.getNodeInfo(ch)
	}
	nodeInfoList := []*pbCommon.Node{}
	for range nodes {
		nodeInfoList = append(nodeInfoList, <-ch)
	}

	response.Nodes = nodeInfoList

	return response, nil
}

// GetNode ... (master)
func (srv *GomlapiServer) GetNode(ctx context.Context, req *pbGomlapi.GetNodeRequest) (*pbGomlapi.GetNodeResponse, error) {
	node, err := newNode(req.NodeAddress)
	if err != nil {
		log.Printf("[Error] %v: %v", req.NodeAddress, err)
		return nil, err
	}

	ch := make(chan *pbCommon.Node, 1)
	err = node.getNodeInfo(ch)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}

	nodeInfo := <-ch

	response := &pbGomlapi.GetNodeResponse{
		NodeId:    nodeInfo.Id,
		NodeAlias: nodeInfo.Alias,
		Address:   nodeInfo.Address,
		Hostname:  nodeInfo.Hostname,
	}

	return response, nil
}

// CreateContainer ... (master)
func (srv *GomlapiServer) CreateContainer(ctx context.Context, req *pbGomlapi.CreateContainerRequest) (*pbGomlapi.CreateContainerResponse, error) {
	nodeSummary, err := getNodeSummary(req.NodeId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}
	address := nodeSummary.Address

	node, err := newNode(address)
	if err != nil {
		log.Printf("[Error] %v: %v", address, err)
		return nil, err
	}

	createContainerRequest := &pbGomlet.CreateContainerRequest{
		ContainerName: req.ContainerName,
		ImageName:     req.ImageName,
		Volumes:       req.Volumes,
		Labels:        req.Labels,
		Envs:          req.Envs,
		Ports:         req.Ports,
		AutoPort:      req.AutoPort,
	}

	createContainerResponse, err := node.createContainer(ctx, createContainerRequest)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}

	response := &pbGomlapi.CreateContainerResponse{}
	response.Id = createContainerResponse.Id
	response.Warnings = createContainerResponse.Warnings

	return response, nil
}

// StartContainer ...
func (srv *GomlapiServer) StartContainer(ctx context.Context, req *pbGomlapi.StartContainerRequest) (*empty.Empty, error) {
	nodeSummary, err := getNodeSummary(req.NodeId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}
	address := nodeSummary.Address

	node, err := newNode(address)
	if err != nil {
		log.Printf("[Error] %v: %v", address, err)
		return nil, err
	}

	err = node.startContainer(ctx, req.ContainerId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

// StopContainer ...
func (srv *GomlapiServer) StopContainer(ctx context.Context, req *pbGomlapi.StopContainerRequest) (*empty.Empty, error) {
	nodeSummary, err := getNodeSummary(req.NodeId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}
	address := nodeSummary.Address

	node, err := newNode(address)
	if err != nil {
		log.Printf("[Error] %v: %v", address, err)
		return nil, err
	}

	err = node.stopContainer(ctx, req.ContainerId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}

// DeleteContainer ...
func (srv *GomlapiServer) DeleteContainer(ctx context.Context, req *pbGomlapi.DeleteContainerRequest) (*empty.Empty, error) {
	nodeSummary, err := getNodeSummary(req.NodeId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}
	address := nodeSummary.Address

	node, err := newNode(address)
	if err != nil {
		log.Printf("[Error] %v: %v", address, err)
		return nil, err
	}

	err = node.deleteContainer(ctx, req.ContainerId)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}

	return &empty.Empty{}, nil
}
