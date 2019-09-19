package server

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/golang/protobuf/ptypes/empty"
	pbCommon "github.com/goml/gomlet/proto/v1/common"
	pbGomlet "github.com/goml/gomlet/proto/v1/gomlet"
	"google.golang.org/grpc"
)

// GomletServer ...
type GomletServer struct {
	GrpcServer *grpc.Server
}

// NewGomletServer ...
func NewGomletServer() *GomletServer {
	// get hostname
	if h, err := os.Hostname(); err == nil {
		hostname = h
	} else {
		panic(err)
	}

	// create server
	grpcServer := grpc.NewServer()
	server := &GomletServer{GrpcServer: grpcServer}
	pbGomlet.RegisterGomletServer(grpcServer, server)
	return server
}

// GetNode ...
func (srv *GomletServer) GetNode(ctx context.Context, req *pbGomlet.GetNodeRequest) (*pbGomlet.GetNodeResponse, error) {
	response := &pbGomlet.GetNodeResponse{
		Node: &pbCommon.Node{
			Id:       system.Config.NodeId,
			Hostname: hostname,
		},
	}

	if err := fillGpuStatus(response.Node); err != nil {
		return nil, err
	}

	if err := fillContainers(ctx, response.Node); err != nil {
		return nil, err
	}

	return response, nil
}

// CreateContainer ...
func (srv *GomletServer) CreateContainer(ctx context.Context, req *pbGomlet.CreateContainerRequest) (*pbGomlet.CreateContainerResponse, error) {
	imageName := req.ImageName
	containerName := req.ContainerName
	labels := req.Labels
	envs := req.Envs
	volumes := req.Volumes
	ports := req.Ports

	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	// image를 확인하고 없다면 pull
	_, err = checkImageAndPull(ctx, cli, imageName)
	if err != nil {
		return nil, err
	}

	// create container
	mounts := []mount.Mount{
		mount.Mount{Type: mount.TypeBind, Source: "/etc/timezone", Target: "/etc/timezone"},
		mount.Mount{Type: mount.TypeBind, Source: "/etc/localtime", Target: "/etc/localtime"},
	}

	for k, v := range volumes {
		m := mount.Mount{
			Type:   mount.TypeBind,
			Source: k,
			Target: v,
		}
		mounts = append(mounts, m)
	}

	portBindings := nat.PortMap{}

	for k, v := range ports {
		key := nat.Port(v)
		if req.AutoPort == true {
			portBindings[key] = []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "0"}}
		} else {
			portBindings[key] = []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: fmt.Sprintf("%d", k)}}
		}
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:  imageName,
		Labels: labels,
		Env:    envs,
	}, &container.HostConfig{
		Mounts:       mounts,
		PortBindings: portBindings,
	}, nil, containerName)
	if err != nil {
		return nil, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	response := &pbGomlet.CreateContainerResponse{
		Id:       resp.ID,
		Warnings: resp.Warnings,
	}

	return response, nil
}

// StartContainer ...
func (srv *GomletServer) StartContainer(ctx context.Context, req *pbGomlet.StartContainerRequest) (*empty.Empty, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	containerID := req.Id

	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// StopContainer ...
func (srv *GomletServer) StopContainer(ctx context.Context, req *pbGomlet.StopContainerRequest) (*empty.Empty, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	containerID := req.Id

	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// DeleteContainer ...
func (srv *GomletServer) DeleteContainer(ctx context.Context, req *pbGomlet.DeleteContainerRequest) (*empty.Empty, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	containerID := req.Id

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

// GetContainers ...
func (srv *GomletServer) GetContainers(ctx context.Context, req *pbGomlet.GetContainersRequest) (*pbGomlet.GetContainersResponse, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}

	response := &pbGomlet.GetContainersResponse{Containers: []*pbCommon.Container{}}

	for _, container := range containers {
		res, _, err := cli.ContainerInspectWithRaw(ctx, container.ID, true)
		if err != nil {
			appendLog(fmt.Sprintf("[Error] %s: %v", container.ID, err))
			return nil, err
		}

		pbContainer, err := convertToContainer(container, res.Config)
		if err != nil {
			return nil, err
		}

		response.Containers = append(response.Containers, pbContainer)
	}

	return response, nil
}
