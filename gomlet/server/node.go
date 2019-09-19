package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	pbCommon "github.com/goml/gomlet/proto/v1/common"
	pbGomlet "github.com/goml/gomlet/proto/v1/gomlet"
	"google.golang.org/grpc"
)

type node struct {
	address string
}

func (n node) getURL() string {
	url := fmt.Sprintf("%s:7777", n.address)
	return url
}

func (n node) getNodeInfo(ch chan<- *pbCommon.Node) error {
	conn, err := grpc.Dial(n.getURL(), grpc.WithInsecure())
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	client := pbGomlet.NewGomletClient(conn)

	request := &pbGomlet.GetNodeRequest{}

	response, err := client.GetNode(ctx, request)
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}

	node := response.Node

	tokens := strings.Split(n.getURL(), ":")
	node.Address = tokens[0]
	if port, err := strconv.Atoi(tokens[1]); err == nil {
		node.Port = int32(port)
	} else {
		return err
	}

	nodeSummary, err := getNodeSummary(node.Id)
	if err == nil {
		node.Alias = nodeSummary.Alias
	}

	ch <- node
	return nil
}

func (n node) createContainer(ctx context.Context, req *pbGomlet.CreateContainerRequest) (*pbGomlet.CreateContainerResponse, error) {
	conn, err := grpc.Dial(n.getURL(), grpc.WithInsecure())
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}
	defer conn.Close()

	client := pbGomlet.NewGomletClient(conn)

	response, err := client.CreateContainer(ctx, req)
	if err != nil {
		log.Printf("[Error] %v", err)
		return nil, err
	}

	return response, nil
}

func (n node) stopContainer(ctx context.Context, id string) error {
	conn, err := grpc.Dial(n.getURL(), grpc.WithInsecure())
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}
	defer conn.Close()

	client := pbGomlet.NewGomletClient(conn)

	request := &pbGomlet.StopContainerRequest{
		Id: id,
	}
	_, err = client.StopContainer(ctx, request)
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}

	return nil
}

func (n node) startContainer(ctx context.Context, id string) error {
	conn, err := grpc.Dial(n.getURL(), grpc.WithInsecure())
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}
	defer conn.Close()

	client := pbGomlet.NewGomletClient(conn)

	request := &pbGomlet.StartContainerRequest{
		Id: id,
	}
	_, err = client.StartContainer(ctx, request)
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}

	return nil
}

func (n node) deleteContainer(ctx context.Context, id string) error {
	conn, err := grpc.Dial(n.getURL(), grpc.WithInsecure())
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}
	defer conn.Close()

	client := pbGomlet.NewGomletClient(conn)

	request := &pbGomlet.DeleteContainerRequest{
		Id: id,
	}
	_, err = client.DeleteContainer(ctx, request)
	if err != nil {
		log.Printf("[Error] %v", err)
		return err
	}

	return nil
}

func newNodes(addresses []string) []*node {
	nodes := []*node{}
	for _, address := range addresses {
		if address == "" {
			continue
		}
		nodes = append(nodes, &node{address: address})
	}

	return nodes
}

func newNode(address string) (*node, error) {
	if address == "" {
		return nil, errors.New("Address is not set")
	}

	return &node{address: address}, nil
}
