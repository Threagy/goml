package server

import (
	"context"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	linq "github.com/ahmetb/go-linq"
	"github.com/davecgh/go-spew/spew"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/goml/gomlet/db"
	pbAuth "github.com/goml/gomlet/proto/v1/auth"
	pbCommon "github.com/goml/gomlet/proto/v1/common"
	pbGomlapi "github.com/goml/gomlet/proto/v1/gomlapi"
	pbUser "github.com/goml/gomlet/proto/v1/user"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthClaims ..
type AuthClaims struct {
	*jwt.StandardClaims
	TokenType string
	db.UserModel
}

func appendLog(log string) {
	system.Logs = append(system.Logs, log)
}

func setState(state pbGomlapi.GomlapiState) {
	system.State = state
	log.Printf("State: %v", state)
}

func getNodeSummary(nodeID string) (*pbCommon.NodeSummary, error) {
	pick := linq.From(system.Config.Nodes).WhereT(func(n *pbCommon.NodeSummary) bool {
		return n.Id == nodeID
	}).First()

	if pick == nil {
		return nil, fmt.Errorf("Node[%s] is not found", nodeID)
	}

	nodeSummary, ok := pick.(*pbCommon.NodeSummary)
	if !ok {
		return nil, fmt.Errorf("Node[%s] convert to fail", nodeID)
	}

	return nodeSummary, nil
}

func refreshDockerImages() error {
	if system.State != pbGomlapi.GomlapiState_IDLE &&
		system.State != pbGomlapi.GomlapiState_INITIALIZING &&
		system.State != pbGomlapi.GomlapiState_IMAGE_PULLING_FAIL {
		return errors.New("System State is not idle")
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	ctx := context.Background()

	setState(pbGomlapi.GomlapiState_IMAGE_PULLING)

	for _, imageName := range system.Config.Images {
		// image를 확인하고 없다면 pull
		imageID, err := checkImageAndPull(ctx, cli, imageName)
		if err != nil {
			appendLog(fmt.Sprintf("[Error] %s: %v", imageName, err))
			setState(pbGomlapi.GomlapiState_IMAGE_PULLING_FAIL)
			return err
		}

		res, _, err := cli.ImageInspectWithRaw(ctx, imageID)
		if err != nil {
			appendLog(fmt.Sprintf("[Error] %s: %v", imageName, err))
			setState(pbGomlapi.GomlapiState_IMAGE_PULLING_FAIL)
			return err
		}

		pbImageInspect, err := convertToImageInspect(res)
		if err != nil {
			appendLog(fmt.Sprintf("[Error] %s: %v", imageName, err))
			setState(pbGomlapi.GomlapiState_IMAGE_PULLING_FAIL)
			return err
		}
		system.Config.ImageInspects = append(system.Config.ImageInspects, pbImageInspect)
	}

	setState(pbGomlapi.GomlapiState_IDLE)
	system.Logs = []string{}
	return nil
}

func convertToImageSummary(image types.ImageSummary) (*pbCommon.ImageSummary, error) {
	pbImage := &pbCommon.ImageSummary{
		Containers:  image.Containers,
		Created:     image.Created,
		Id:          image.ID,
		Labels:      image.Labels,
		ParentId:    image.ParentID,
		RepoDigests: image.RepoDigests,
		RepoTags:    image.RepoTags,
		SharedSize:  image.SharedSize,
		Size:        image.Size,
		VirtualSize: image.VirtualSize,
	}

	return pbImage, nil
}

func convertToImageInspect(image types.ImageInspect) (*pbCommon.ImageInspect, error) {
	pbImage := &pbCommon.ImageInspect{
		Id:            image.ID,
		RepoTags:      image.RepoTags,
		RepoDigests:   image.RepoDigests,
		Parent:        image.Parent,
		Comment:       image.Comment,
		Created:       image.Created,
		Container:     image.Container,
		DockerVersion: image.DockerVersion,
		Author:        image.Author,
		Architecture:  image.Architecture,
		Os:            image.Os,
		OsVersion:     image.OsVersion,
		Size:          image.Size,
		VirtualSize:   image.VirtualSize,
	}

	pbImage.ContainerConfig = &pbCommon.ContainerConfig{
		Hostname:   image.ContainerConfig.Hostname,
		Domainname: image.ContainerConfig.Domainname,
		User:       image.ContainerConfig.User,
		Env:        image.ContainerConfig.Env,
		Image:      image.ContainerConfig.Image,
		WorkingDir: image.ContainerConfig.WorkingDir,
		Labels:     image.ContainerConfig.Labels,
	}

	// ports
	pbImage.ContainerConfig.ExposedPorts = []string{}
	for key := range image.ContainerConfig.ExposedPorts {
		pbImage.ContainerConfig.ExposedPorts = append(pbImage.ContainerConfig.ExposedPorts, string(key))
	}
	sort.Slice(pbImage.ContainerConfig.ExposedPorts, func(i, j int) bool {
		return pbImage.ContainerConfig.ExposedPorts[i] < pbImage.ContainerConfig.ExposedPorts[j]
	})

	// volumes
	for key := range image.ContainerConfig.Volumes {
		pbImage.ContainerConfig.Volumes = append(pbImage.ContainerConfig.Volumes, key)
	}
	sort.Slice(pbImage.ContainerConfig.Volumes, func(i, j int) bool {
		return pbImage.ContainerConfig.Volumes[i] < pbImage.ContainerConfig.Volumes[j]
	})

	return pbImage, nil
}

func convertToContainer(container types.Container, config *container.Config) (*pbCommon.Container, error) {
	pbContainer := &pbCommon.Container{
		Id:         container.ID,
		Names:      container.Names,
		Image:      container.Image,
		ImageId:    container.ImageID,
		Command:    container.Command,
		Created:    container.Created,
		SizeRw:     container.SizeRw,
		SizeRootFs: container.SizeRootFs,
		Labels:     container.Labels,
		State:      container.State,
		Status:     container.Status,
		HostConfig: &pbCommon.NetworkMode{
			NetworkMode: container.HostConfig.NetworkMode,
		},
	}

	pbContainer.Ports = []*pbCommon.Port{}
	for _, p := range container.Ports {
		pbContainer.Ports = append(pbContainer.Ports, &pbCommon.Port{
			Ip:          p.IP,
			PrivatePort: int32(p.PrivatePort),
			PublicPort:  int32(p.PublicPort),
			Type:        p.Type,
		})
	}
	sort.Slice(pbContainer.Ports, func(i, j int) bool {
		return pbContainer.Ports[i].PrivatePort < pbContainer.Ports[j].PrivatePort
	})

	pbContainer.Mounts = []*pbCommon.MountPoint{}
	for _, m := range container.Mounts {
		pbContainer.Mounts = append(pbContainer.Mounts, &pbCommon.MountPoint{
			Type:        string(m.Type),
			Name:        m.Name,
			Source:      m.Source,
			Destination: m.Destination,
			Driver:      m.Driver,
			Mode:        m.Mode,
			Rw:          m.RW,
			Propagation: string(m.Propagation),
		})
	}

	pbContainer.Env = config.Env

	sort.Slice(pbContainer.Mounts, func(i, j int) bool {
		return pbContainer.Mounts[i].Source < pbContainer.Mounts[j].Source
	})

	return pbContainer, nil
}

func convertToUser(users []*db.UserModel) []*pbUser.User {
	slice := []*pbUser.User{}

	for _, user := range users {
		u := &pbUser.User{}
		u.Name = user.Name
		u.UserId = user.UserID
		u.Role = int32(user.Role)

		slice = append(slice, u)
	}

	return slice
}

type logWriter struct {
}

func (w *logWriter) Write(p []byte) (int, error) {
	value := string(p)
	lines := strings.Split(value, "\r\n")

	linq.From(lines).WhereT(func(line string) bool {
		return line != ""
	}).ToSlice(&lines)

	system.Logs = append(system.Logs, lines...)
	log.Println(lines)

	return len(p), nil
}

func checkImageAndPull(ctx context.Context, cli *client.Client, imageName string) (string, error) {
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return "", err
	}

	isExist := false
	for _, image := range images {
		for _, repoTag := range image.RepoTags {
			if strings.Index(imageName, repoTag) != -1 {
				return image.ID, nil
			}
		}
	}

	log.Printf("%v %v", isExist, imageName)
	if isExist == false {
		var pullImageName = imageName

		// // Docker Hub Official Image는 docker.io/libaray를 추가한다.
		// if strings.Index(pullImageName, "/") == -1 {
		// 	pullImageName = fmt.Sprintf("docker.io/library/%s", imageName)
		// }
		// log.Println(pullImageName)
		out, err := cli.ImagePull(ctx, pullImageName, types.ImagePullOptions{})
		if err != nil {
			return "", err
		}

		defer out.Close()

		io.Copy(&logWriter{}, out)
	}

	images, err = cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return "", err
	}

	for _, image := range images {
		for _, repoTag := range image.RepoTags {
			if strings.Index(imageName, repoTag) != -1 {
				return image.ID, nil
			}
		}
	}

	return "", errors.New("Image is not found")
}

func newGrpcMux(address string) *runtime.ServeMux {
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	ctx := context.Background()

	// auth server
	err := pbAuth.RegisterAuthHandlerFromEndpoint(ctx, gwmux, address, dopts)
	if err != nil {
		panic(err)
	}

	// user server
	err = pbUser.RegisterUserManagerHandlerFromEndpoint(ctx, gwmux, address, dopts)
	if err != nil {
		panic(err)
	}

	// gomlapi server
	err = pbGomlapi.RegisterGomlapiHandlerFromEndpoint(ctx, gwmux, address, dopts)
	if err != nil {
		panic(err)
	}

	return gwmux
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise.
// https://github.com/grpc/grpc-go/issues/555
func grpcHandlerFunc(authServer *grpc.Server, userServer *grpc.Server, gomlapiServer *grpc.Server, gomletServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				enableCors(&w)

				if r.Header.Get("Origin") != "" && r.Method == "OPTIONS" {
					return
				}

				if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
					// log.Printf("GRPC: %v", r.Header.Get("Content-Type"))
					// log.Printf("GRPC: %v, %v", r.Header.Get("Content-Type"), r.RequestURI)

					if strings.Index(r.RequestURI, "/proto.gomlapi.Gomlapi") != -1 {
						gomlapiServer.ServeHTTP(w, r)
					} else if strings.Index(r.RequestURI, "/proto.gomlet.Gomlet") != -1 {
						gomletServer.ServeHTTP(w, r)
					} else if strings.Index(r.RequestURI, "/proto.user.User") != -1 {
						userServer.ServeHTTP(w, r)
					} else if strings.Index(r.RequestURI, "/proto.auth.Auth") != -1 {
						authServer.ServeHTTP(w, r)
					}
				} else {
					// logger.Debugf("HTTP: %v", r.Header.Get("Content-Type"))
					otherHandler.ServeHTTP(w, r)
				}
			}),
		&http2.Server{})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, GET, POST, PUT, PATCH, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "X-Requested-With, origin, content-type, accept, authorization")
}

func fillContainers(ctx context.Context, node *pbCommon.Node) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		return err
	}

	node.Containers = []*pbCommon.Container{}

	for _, container := range containers {
		res, _, err := cli.ContainerInspectWithRaw(ctx, container.ID, true)
		if err != nil {
			appendLog(fmt.Sprintf("[Error] %s: %v", container.ID, err))
			return err
		}

		pbContainer, err := convertToContainer(container, res.Config)
		if err != nil {
			return err
		}

		node.Containers = append(node.Containers, pbContainer)
	}

	return nil
}

func fillGpuStatus(node *pbCommon.Node) error {
	cmder := newCommander()
	cmd := exec.Command(
		"nvidia-smi",
		"--query-gpu=index,memory.total,memory.used,temperature.gpu",
		"--format=csv")
	reader := cmder.start(cmd)

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		// Ignore header
		if strings.Contains(line, "index") {
			continue
		}
		if len(line) == 0 {
			return errors.New("Unable to fetch any events from nvidia-smi: Error " + err.Error())
		}

		// Remove units put by nvidia-smi
		line = strings.Replace(line, " %", "", -1)
		line = strings.Replace(line, " MiB", "", -1)
		line = strings.Replace(line, " P", "", -1)
		line = strings.Replace(line, " ", "", -1)

		r := csv.NewReader(strings.NewReader(line))
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		status := &pbCommon.GPUStatus{}
		if value, err := strconv.Atoi(record[0]); err == nil {
			if value > 0 {

			}
			status.Index = uint32(value)
		}
		if value, err := strconv.ParseInt(record[1], 10, 64); err == nil {
			status.MemoryTotalSize = uint64(value)
		}
		if value, err := strconv.ParseInt(record[2], 10, 64); err == nil {
			status.MemoryUsedSize = uint64(value)
		}
		if value, err := strconv.ParseInt(record[3], 10, 64); err == nil {
			status.Temperature = uint32(value)
		}

		node.GpuStatuses = append(node.GpuStatuses, status)
	}
	cmd.Wait()

	return nil
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		log.Print(authorization)
		return false
	}

	token := authorization[0]
	claims := &AuthClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		log.Print(err)
		return false
	}

	return true
}

func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	// log.Print(md)

	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

func execResizeTty(w http.ResponseWriter, r *http.Request) {
	execIDs, _ := r.URL.Query()["execID"]
	widths, _ := r.URL.Query()["width"]
	heights, _ := r.URL.Query()["height"]

	execID := execIDs[0]
	width, _ := strconv.Atoi(widths[0])
	height, _ := strconv.Atoi(heights[0])

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print(err)
		return
	}

	ctx := context.Background()

	err = cli.ContainerExecResize(ctx, execID, types.ResizeOptions{
		Width:  uint(width),
		Height: uint(height),
	})
	if err != nil {
		log.Print(err)
	}
}

func execContainer(ws *websocket.Conn) {
	wsParams := strings.Split(ws.Request().URL.Path[len("/v1/exec/"):], ",")
	container := wsParams[0]
	cmd, _ := base64.StdEncoding.DecodeString(wsParams[1])

	if container == "" {
		ws.Write([]byte("Container does not exist"))
		return
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		log.Print(err)
		return
	}

	ctx := context.Background()

	resp, err := cli.ContainerExecCreate(ctx, container, types.ExecConfig{
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          []string{string(cmd)},
		Tty:          true,
	})

	if err != nil {
		log.Print(err)
		return
	}

	execID := resp.ID
	ws.Write([]byte(execID))
	ws.Write([]byte("\r\n"))

	if err := hijack(ctx, cli, resp.ID, true, ws, ws, ws, nil, nil); err != nil {
		panic(err)
	}

	// fmt.Println("Connection!")
	// fmt.Println(ws)
	// spew.Dump(ws)
}

func hijack(ctx context.Context, cli *client.Client, execID string, setRawTerminal bool, in io.ReadCloser, stdout, stderr io.Writer, started chan io.Closer, data interface{}) error {
	// err := cli.ContainerExecStart(ctx, execID, types.ExecStartCheck{
	// 	Detach: false,
	// 	Tty:    true,
	// })

	// if err != nil {
	// 	log.Print(err)
	// 	return err
	// }

	// log.Printf("hijack: %v", execID)
	hijack, err := cli.ContainerExecAttach(ctx, execID, types.ExecConfig{
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		// Cmd:          []string{"/bin/bash"},
		Tty: true,
	})

	if err != nil {
		log.Print(err)
		return err
	}

	rwc := hijack.Conn
	br := hijack.Reader
	defer rwc.Close()

	if started != nil {
		started <- rwc
	}

	var receiveStdout chan error

	if stdout != nil || stderr != nil {
		go func() (err error) {
			if setRawTerminal && stdout != nil {
				_, err = io.Copy(stdout, br)
			}
			return err
		}()
	}

	go func() error {
		if in != nil {
			io.Copy(rwc, in)
		}

		if conn, ok := rwc.(interface {
			CloseWrite() error
		}); ok {
			if err := conn.CloseWrite(); err != nil {
			}
		}
		return nil
	}()

	if stdout != nil || stderr != nil {
		if err := <-receiveStdout; err != nil {
			return err
		}
	}
	spew.Dump(br)
	go func() {
		for {
			fmt.Println(br)
			spew.Dump(br)
		}
	}()

	return nil
}
