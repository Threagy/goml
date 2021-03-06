// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/v1/common/docker.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ImageSummary struct {
	Containers           int64             `protobuf:"varint,1,opt,name=containers,proto3" json:"containers,omitempty"`
	Created              int64             `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Id                   string            `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Labels               map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ParentId             string            `protobuf:"bytes,5,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	RepoDigests          []string          `protobuf:"bytes,6,rep,name=repo_digests,json=repoDigests,proto3" json:"repo_digests,omitempty"`
	RepoTags             []string          `protobuf:"bytes,7,rep,name=repo_tags,json=repoTags,proto3" json:"repo_tags,omitempty"`
	SharedSize           int64             `protobuf:"varint,8,opt,name=shared_size,json=sharedSize,proto3" json:"shared_size,omitempty"`
	Size                 int64             `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	VirtualSize          int64             `protobuf:"varint,10,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ImageSummary) Reset()         { *m = ImageSummary{} }
func (m *ImageSummary) String() string { return proto.CompactTextString(m) }
func (*ImageSummary) ProtoMessage()    {}
func (*ImageSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{0}
}

func (m *ImageSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageSummary.Unmarshal(m, b)
}
func (m *ImageSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageSummary.Marshal(b, m, deterministic)
}
func (m *ImageSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageSummary.Merge(m, src)
}
func (m *ImageSummary) XXX_Size() int {
	return xxx_messageInfo_ImageSummary.Size(m)
}
func (m *ImageSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ImageSummary proto.InternalMessageInfo

func (m *ImageSummary) GetContainers() int64 {
	if m != nil {
		return m.Containers
	}
	return 0
}

func (m *ImageSummary) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ImageSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ImageSummary) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ImageSummary) GetParentId() string {
	if m != nil {
		return m.ParentId
	}
	return ""
}

func (m *ImageSummary) GetRepoDigests() []string {
	if m != nil {
		return m.RepoDigests
	}
	return nil
}

func (m *ImageSummary) GetRepoTags() []string {
	if m != nil {
		return m.RepoTags
	}
	return nil
}

func (m *ImageSummary) GetSharedSize() int64 {
	if m != nil {
		return m.SharedSize
	}
	return 0
}

func (m *ImageSummary) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ImageSummary) GetVirtualSize() int64 {
	if m != nil {
		return m.VirtualSize
	}
	return 0
}

type ImageInspect struct {
	Id                   string           `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	RepoTags             []string         `protobuf:"bytes,2,rep,name=repo_tags,json=repoTags,proto3" json:"repo_tags,omitempty"`
	RepoDigests          []string         `protobuf:"bytes,3,rep,name=repo_digests,json=repoDigests,proto3" json:"repo_digests,omitempty"`
	Parent               string           `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Comment              string           `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	Created              string           `protobuf:"bytes,6,opt,name=created,proto3" json:"created,omitempty"`
	Container            string           `protobuf:"bytes,7,opt,name=container,proto3" json:"container,omitempty"`
	ContainerConfig      *ContainerConfig `protobuf:"bytes,8,opt,name=container_config,json=containerConfig,proto3" json:"container_config,omitempty"`
	DockerVersion        string           `protobuf:"bytes,9,opt,name=docker_version,json=dockerVersion,proto3" json:"docker_version,omitempty"`
	Author               string           `protobuf:"bytes,10,opt,name=author,proto3" json:"author,omitempty"`
	Config               *ContainerConfig `protobuf:"bytes,11,opt,name=config,proto3" json:"config,omitempty"`
	Architecture         string           `protobuf:"bytes,12,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Os                   string           `protobuf:"bytes,13,opt,name=os,proto3" json:"os,omitempty"`
	OsVersion            string           `protobuf:"bytes,14,opt,name=os_version,json=osVersion,proto3" json:"os_version,omitempty"`
	Size                 int64            `protobuf:"varint,15,opt,name=size,proto3" json:"size,omitempty"`
	VirtualSize          int64            `protobuf:"varint,16,opt,name=virtual_size,json=virtualSize,proto3" json:"virtual_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ImageInspect) Reset()         { *m = ImageInspect{} }
func (m *ImageInspect) String() string { return proto.CompactTextString(m) }
func (*ImageInspect) ProtoMessage()    {}
func (*ImageInspect) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{1}
}

func (m *ImageInspect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageInspect.Unmarshal(m, b)
}
func (m *ImageInspect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageInspect.Marshal(b, m, deterministic)
}
func (m *ImageInspect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageInspect.Merge(m, src)
}
func (m *ImageInspect) XXX_Size() int {
	return xxx_messageInfo_ImageInspect.Size(m)
}
func (m *ImageInspect) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageInspect.DiscardUnknown(m)
}

var xxx_messageInfo_ImageInspect proto.InternalMessageInfo

func (m *ImageInspect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ImageInspect) GetRepoTags() []string {
	if m != nil {
		return m.RepoTags
	}
	return nil
}

func (m *ImageInspect) GetRepoDigests() []string {
	if m != nil {
		return m.RepoDigests
	}
	return nil
}

func (m *ImageInspect) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ImageInspect) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *ImageInspect) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *ImageInspect) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *ImageInspect) GetContainerConfig() *ContainerConfig {
	if m != nil {
		return m.ContainerConfig
	}
	return nil
}

func (m *ImageInspect) GetDockerVersion() string {
	if m != nil {
		return m.DockerVersion
	}
	return ""
}

func (m *ImageInspect) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *ImageInspect) GetConfig() *ContainerConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ImageInspect) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *ImageInspect) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *ImageInspect) GetOsVersion() string {
	if m != nil {
		return m.OsVersion
	}
	return ""
}

func (m *ImageInspect) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ImageInspect) GetVirtualSize() int64 {
	if m != nil {
		return m.VirtualSize
	}
	return 0
}

type ContainerConfig struct {
	Hostname             string            `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Domainname           string            `protobuf:"bytes,2,opt,name=domainname,proto3" json:"domainname,omitempty"`
	User                 string            `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	ExposedPorts         []string          `protobuf:"bytes,4,rep,name=exposed_ports,json=exposedPorts,proto3" json:"exposed_ports,omitempty"`
	Env                  []string          `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty"`
	Image                string            `protobuf:"bytes,6,opt,name=image,proto3" json:"image,omitempty"`
	Volumes              []string          `protobuf:"bytes,7,rep,name=volumes,proto3" json:"volumes,omitempty"`
	WorkingDir           string            `protobuf:"bytes,8,opt,name=working_dir,json=workingDir,proto3" json:"working_dir,omitempty"`
	Labels               map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ContainerConfig) Reset()         { *m = ContainerConfig{} }
func (m *ContainerConfig) String() string { return proto.CompactTextString(m) }
func (*ContainerConfig) ProtoMessage()    {}
func (*ContainerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{2}
}

func (m *ContainerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerConfig.Unmarshal(m, b)
}
func (m *ContainerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerConfig.Marshal(b, m, deterministic)
}
func (m *ContainerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerConfig.Merge(m, src)
}
func (m *ContainerConfig) XXX_Size() int {
	return xxx_messageInfo_ContainerConfig.Size(m)
}
func (m *ContainerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerConfig proto.InternalMessageInfo

func (m *ContainerConfig) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *ContainerConfig) GetDomainname() string {
	if m != nil {
		return m.Domainname
	}
	return ""
}

func (m *ContainerConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ContainerConfig) GetExposedPorts() []string {
	if m != nil {
		return m.ExposedPorts
	}
	return nil
}

func (m *ContainerConfig) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *ContainerConfig) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *ContainerConfig) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *ContainerConfig) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *ContainerConfig) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Container struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Names                []string          `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	Image                string            `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	ImageId              string            `protobuf:"bytes,4,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	Command              string            `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	Created              int64             `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"`
	Ports                []*Port           `protobuf:"bytes,7,rep,name=ports,proto3" json:"ports,omitempty"`
	SizeRw               int64             `protobuf:"varint,8,opt,name=size_rw,json=sizeRw,proto3" json:"size_rw,omitempty"`
	SizeRootFs           int64             `protobuf:"varint,9,opt,name=size_root_fs,json=sizeRootFs,proto3" json:"size_root_fs,omitempty"`
	Labels               map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State                string            `protobuf:"bytes,11,opt,name=state,proto3" json:"state,omitempty"`
	Status               string            `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	HostConfig           *NetworkMode      `protobuf:"bytes,13,opt,name=host_config,json=hostConfig,proto3" json:"host_config,omitempty"`
	Mounts               []*MountPoint     `protobuf:"bytes,14,rep,name=mounts,proto3" json:"mounts,omitempty"`
	Env                  []string          `protobuf:"bytes,15,rep,name=env,proto3" json:"env,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{3}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

func (m *Container) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Container) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *Container) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *Container) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Container) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Container) GetSizeRw() int64 {
	if m != nil {
		return m.SizeRw
	}
	return 0
}

func (m *Container) GetSizeRootFs() int64 {
	if m != nil {
		return m.SizeRootFs
	}
	return 0
}

func (m *Container) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Container) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Container) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Container) GetHostConfig() *NetworkMode {
	if m != nil {
		return m.HostConfig
	}
	return nil
}

func (m *Container) GetMounts() []*MountPoint {
	if m != nil {
		return m.Mounts
	}
	return nil
}

func (m *Container) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

type NetworkMode struct {
	NetworkMode          string   `protobuf:"bytes,1,opt,name=network_mode,json=networkMode,proto3" json:"network_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkMode) Reset()         { *m = NetworkMode{} }
func (m *NetworkMode) String() string { return proto.CompactTextString(m) }
func (*NetworkMode) ProtoMessage()    {}
func (*NetworkMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{4}
}

func (m *NetworkMode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkMode.Unmarshal(m, b)
}
func (m *NetworkMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkMode.Marshal(b, m, deterministic)
}
func (m *NetworkMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkMode.Merge(m, src)
}
func (m *NetworkMode) XXX_Size() int {
	return xxx_messageInfo_NetworkMode.Size(m)
}
func (m *NetworkMode) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkMode.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkMode proto.InternalMessageInfo

func (m *NetworkMode) GetNetworkMode() string {
	if m != nil {
		return m.NetworkMode
	}
	return ""
}

type Port struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	PrivatePort          int32    `protobuf:"varint,2,opt,name=private_port,json=privatePort,proto3" json:"private_port,omitempty"`
	PublicPort           int32    `protobuf:"varint,3,opt,name=public_port,json=publicPort,proto3" json:"public_port,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{5}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Port) GetPrivatePort() int32 {
	if m != nil {
		return m.PrivatePort
	}
	return 0
}

func (m *Port) GetPublicPort() int32 {
	if m != nil {
		return m.PublicPort
	}
	return 0
}

func (m *Port) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type MountPoint struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Destination          string   `protobuf:"bytes,4,opt,name=destination,proto3" json:"destination,omitempty"`
	Driver               string   `protobuf:"bytes,5,opt,name=driver,proto3" json:"driver,omitempty"`
	Mode                 string   `protobuf:"bytes,6,opt,name=mode,proto3" json:"mode,omitempty"`
	Rw                   bool     `protobuf:"varint,7,opt,name=rw,proto3" json:"rw,omitempty"`
	Propagation          string   `protobuf:"bytes,8,opt,name=propagation,proto3" json:"propagation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MountPoint) Reset()         { *m = MountPoint{} }
func (m *MountPoint) String() string { return proto.CompactTextString(m) }
func (*MountPoint) ProtoMessage()    {}
func (*MountPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_00285085086f4906, []int{6}
}

func (m *MountPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MountPoint.Unmarshal(m, b)
}
func (m *MountPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MountPoint.Marshal(b, m, deterministic)
}
func (m *MountPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MountPoint.Merge(m, src)
}
func (m *MountPoint) XXX_Size() int {
	return xxx_messageInfo_MountPoint.Size(m)
}
func (m *MountPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_MountPoint.DiscardUnknown(m)
}

var xxx_messageInfo_MountPoint proto.InternalMessageInfo

func (m *MountPoint) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MountPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MountPoint) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *MountPoint) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *MountPoint) GetDriver() string {
	if m != nil {
		return m.Driver
	}
	return ""
}

func (m *MountPoint) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *MountPoint) GetRw() bool {
	if m != nil {
		return m.Rw
	}
	return false
}

func (m *MountPoint) GetPropagation() string {
	if m != nil {
		return m.Propagation
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageSummary)(nil), "proto.common.ImageSummary")
	proto.RegisterMapType((map[string]string)(nil), "proto.common.ImageSummary.LabelsEntry")
	proto.RegisterType((*ImageInspect)(nil), "proto.common.ImageInspect")
	proto.RegisterType((*ContainerConfig)(nil), "proto.common.ContainerConfig")
	proto.RegisterMapType((map[string]string)(nil), "proto.common.ContainerConfig.LabelsEntry")
	proto.RegisterType((*Container)(nil), "proto.common.Container")
	proto.RegisterMapType((map[string]string)(nil), "proto.common.Container.LabelsEntry")
	proto.RegisterType((*NetworkMode)(nil), "proto.common.NetworkMode")
	proto.RegisterType((*Port)(nil), "proto.common.Port")
	proto.RegisterType((*MountPoint)(nil), "proto.common.MountPoint")
}

func init() { proto.RegisterFile("proto/v1/common/docker.proto", fileDescriptor_00285085086f4906) }

var fileDescriptor_00285085086f4906 = []byte{
	// 1021 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x8e, 0x1b, 0x35,
	0x14, 0x56, 0xfe, 0x33, 0x67, 0xb2, 0x3f, 0xb2, 0x56, 0x30, 0x5d, 0xb6, 0x34, 0xa4, 0xa2, 0x0a,
	0x12, 0xdd, 0x94, 0x22, 0x24, 0x28, 0x12, 0x12, 0xb4, 0x20, 0x22, 0x51, 0x54, 0x4d, 0x11, 0x17,
	0xdc, 0x44, 0xde, 0xb1, 0x3b, 0x6b, 0x6d, 0xc6, 0x1e, 0xd9, 0x9e, 0x2c, 0xdb, 0x27, 0xe0, 0x9a,
	0x37, 0xe0, 0x7d, 0x78, 0x0f, 0x5e, 0x03, 0xf9, 0xd8, 0x33, 0x99, 0x64, 0x4b, 0x85, 0xd4, 0x9b,
	0xae, 0xcf, 0x77, 0x8e, 0xc7, 0xf6, 0x77, 0xbe, 0xef, 0x34, 0x70, 0x56, 0x6a, 0x65, 0xd5, 0x62,
	0xf3, 0xd9, 0x22, 0x53, 0x45, 0xa1, 0xe4, 0x82, 0xa9, 0xec, 0x8a, 0xeb, 0x73, 0x84, 0xc9, 0x04,
	0xff, 0x9c, 0xfb, 0xd4, 0xe9, 0x59, 0xae, 0x54, 0xbe, 0xe6, 0x0b, 0x5a, 0x8a, 0x05, 0x95, 0x52,
	0x59, 0x6a, 0x85, 0x92, 0xc6, 0xd7, 0x9e, 0x7e, 0x8a, 0x7f, 0xb2, 0x87, 0x39, 0x97, 0x0f, 0xcd,
	0x35, 0xcd, 0x73, 0xae, 0x17, 0xaa, 0xc4, 0x8a, 0xdb, 0xd5, 0xb3, 0x3f, 0x7a, 0x30, 0x59, 0x16,
	0x34, 0xe7, 0x2f, 0xab, 0xa2, 0xa0, 0xfa, 0x86, 0x7c, 0x08, 0x90, 0x29, 0x69, 0xa9, 0x90, 0x5c,
	0x9b, 0xa4, 0x33, 0xed, 0xcc, 0x7b, 0x69, 0x0b, 0x21, 0x09, 0x8c, 0x32, 0xcd, 0xa9, 0xe5, 0x2c,
	0xe9, 0x62, 0xb2, 0x0e, 0xc9, 0x21, 0x74, 0x05, 0x4b, 0x7a, 0xd3, 0xce, 0x3c, 0x4a, 0xbb, 0x82,
	0x91, 0x6f, 0x60, 0xb8, 0xa6, 0x17, 0x7c, 0x6d, 0x92, 0xfe, 0xb4, 0x37, 0x8f, 0x1f, 0x3f, 0x38,
	0x6f, 0xbf, 0xe2, 0xbc, 0x7d, 0xea, 0xf9, 0x4f, 0x58, 0xf8, 0xbd, 0xb4, 0xfa, 0x26, 0x0d, 0xbb,
	0xc8, 0x07, 0x10, 0x95, 0x54, 0x73, 0x69, 0x57, 0x82, 0x25, 0x03, 0xfc, 0xec, 0xd8, 0x03, 0x4b,
	0x46, 0x3e, 0x82, 0x89, 0xe6, 0xa5, 0x5a, 0x31, 0x91, 0x73, 0x63, 0x4d, 0x32, 0x9c, 0xf6, 0xe6,
	0x51, 0x1a, 0x3b, 0xec, 0x99, 0x87, 0xdc, 0x7e, 0x2c, 0xb1, 0x34, 0x37, 0xc9, 0x08, 0xf3, 0x63,
	0x07, 0xfc, 0x42, 0x73, 0x43, 0xee, 0x41, 0x6c, 0x2e, 0xa9, 0xe6, 0x6c, 0x65, 0xc4, 0x6b, 0x9e,
	0x8c, 0xfd, 0x3b, 0x3d, 0xf4, 0x52, 0xbc, 0xe6, 0x84, 0x40, 0x1f, 0x33, 0x11, 0x66, 0x70, 0xed,
	0x0e, 0xdd, 0x08, 0x6d, 0x2b, 0xba, 0xf6, 0xbb, 0x00, 0x73, 0x71, 0xc0, 0xdc, 0xb6, 0xd3, 0xaf,
	0x20, 0x6e, 0xbd, 0x85, 0x1c, 0x43, 0xef, 0x8a, 0xdf, 0x20, 0x8d, 0x51, 0xea, 0x96, 0xe4, 0x04,
	0x06, 0x1b, 0xba, 0xae, 0x38, 0xb2, 0x17, 0xa5, 0x3e, 0x78, 0xd2, 0xfd, 0xb2, 0x33, 0xfb, 0xb3,
	0x1f, 0x5a, 0xb1, 0x94, 0xa6, 0xe4, 0x99, 0x75, 0x84, 0x2e, 0x59, 0xd8, 0xdb, 0x5d, 0xb2, 0xdd,
	0x07, 0x75, 0xf7, 0x1e, 0xb4, 0x4f, 0x48, 0xef, 0x36, 0x21, 0xef, 0xc1, 0xd0, 0xf3, 0x97, 0xf4,
	0xf1, 0x9b, 0x21, 0xc2, 0x96, 0xaa, 0xa2, 0x70, 0x09, 0x4f, 0x73, 0x1d, 0xb6, 0x9b, 0x3d, 0x0c,
	0x99, 0xd0, 0xec, 0x33, 0x88, 0x1a, 0x51, 0x24, 0x23, 0xcc, 0x6d, 0x01, 0xf2, 0x23, 0x1c, 0x37,
	0xc1, 0x2a, 0x53, 0xf2, 0x95, 0xc8, 0x91, 0xe2, 0xf8, 0xf1, 0xdd, 0x5d, 0x11, 0x3c, 0xad, 0xab,
	0x9e, 0x62, 0x51, 0x7a, 0x94, 0xed, 0x02, 0xe4, 0x63, 0x38, 0xf4, 0x4e, 0x58, 0x6d, 0xb8, 0x36,
	0x42, 0x49, 0x6c, 0x48, 0x94, 0x1e, 0x78, 0xf4, 0x57, 0x0f, 0xba, 0xa7, 0xd1, 0xca, 0x5e, 0x2a,
	0x8d, 0x3d, 0x89, 0xd2, 0x10, 0x91, 0x2f, 0x60, 0x18, 0x8e, 0x8f, 0xff, 0xcf, 0xf1, 0xa1, 0x98,
	0xcc, 0x60, 0x42, 0x75, 0x76, 0x29, 0x2c, 0xcf, 0x6c, 0xa5, 0x79, 0x32, 0xc1, 0x8f, 0xee, 0x60,
	0xae, 0x3b, 0xca, 0x24, 0x07, 0xbe, 0x3b, 0xca, 0x90, 0xbb, 0x00, 0xca, 0x34, 0xb7, 0x3c, 0xf4,
	0x94, 0x28, 0x53, 0xdf, 0xb0, 0xd6, 0xd3, 0xd1, 0x5b, 0xf4, 0x74, 0x7c, 0x4b, 0x4f, 0xb3, 0x7f,
	0xba, 0x70, 0xb4, 0x77, 0x4b, 0x72, 0x0a, 0xe3, 0x4b, 0x65, 0xac, 0xa4, 0x05, 0x0f, 0xea, 0x68,
	0x62, 0x67, 0x5f, 0xa6, 0x0a, 0x2a, 0x24, 0x66, 0xbd, 0xc6, 0x5a, 0x88, 0xbb, 0x46, 0x65, 0xb8,
	0x0e, 0x36, 0xc5, 0x35, 0xb9, 0x0f, 0x07, 0xfc, 0xf7, 0x52, 0x19, 0xce, 0x56, 0xa5, 0xd2, 0xd6,
	0xfb, 0x35, 0x4a, 0x27, 0x01, 0x7c, 0xe1, 0x30, 0xa7, 0x64, 0x2e, 0x37, 0xc9, 0x00, 0x53, 0x6e,
	0xe9, 0x94, 0x2c, 0x9c, 0x5c, 0x83, 0x34, 0x7c, 0xe0, 0x24, 0xb3, 0x51, 0xeb, 0xaa, 0xe0, 0xb5,
	0xe7, 0xea, 0xd0, 0x59, 0xee, 0x5a, 0xe9, 0x2b, 0x21, 0xf3, 0x15, 0x13, 0x1a, 0xf5, 0x10, 0xa5,
	0x10, 0xa0, 0x67, 0x42, 0x93, 0x6f, 0x9b, 0x81, 0x01, 0x38, 0x30, 0x3e, 0x79, 0x6b, 0xb3, 0xde,
	0x34, 0x33, 0xde, 0xc5, 0x7e, 0x7f, 0xf5, 0x21, 0x6a, 0x8e, 0x08, 0xc3, 0xac, 0xd3, 0x0c, 0xb3,
	0x13, 0x18, 0x38, 0xfe, 0x6a, 0xdf, 0xf9, 0x60, 0x4b, 0x41, 0xaf, 0x4d, 0xc1, 0x1d, 0x18, 0xe3,
	0xc2, 0xcd, 0x2d, 0xef, 0xb4, 0x11, 0xc6, 0x4b, 0x56, 0x5b, 0x8d, 0x4a, 0xd6, 0xb6, 0x1a, 0x95,
	0x6c, 0xdf, 0x6a, 0xad, 0xb9, 0x3a, 0x87, 0x81, 0x6f, 0xcb, 0x08, 0x59, 0x21, 0xbb, 0xac, 0xb8,
	0xee, 0xa4, 0xbe, 0x80, 0xbc, 0x0f, 0x23, 0xa7, 0xa3, 0x95, 0xbe, 0x0e, 0x03, 0x6d, 0xe8, 0xc2,
	0xf4, 0x9a, 0x4c, 0x61, 0xe2, 0x13, 0x4a, 0xd9, 0xd5, 0x2b, 0x13, 0x86, 0x1a, 0x60, 0x56, 0x29,
	0xfb, 0x83, 0x21, 0x5f, 0xef, 0x71, 0x7f, 0xff, 0x3f, 0xb8, 0x7f, 0xe3, 0xa4, 0x3e, 0x81, 0x81,
	0xb1, 0xd4, 0x72, 0x34, 0x59, 0x94, 0xfa, 0xc0, 0x79, 0xd2, 0x2d, 0x2a, 0x13, 0xec, 0x13, 0x22,
	0xf2, 0x04, 0x62, 0x27, 0xd7, 0x7a, 0x2e, 0x1c, 0xa0, 0x31, 0xef, 0xec, 0x9e, 0xf7, 0x33, 0xb7,
	0x4e, 0x18, 0xcf, 0x15, 0xe3, 0x29, 0xb8, 0xea, 0x20, 0xfd, 0x47, 0x30, 0x2c, 0x54, 0x25, 0xad,
	0x49, 0x0e, 0xf1, 0x9a, 0xc9, 0xee, 0xb6, 0xe7, 0x2e, 0xf7, 0x42, 0x09, 0x69, 0xd3, 0x50, 0x57,
	0xeb, 0xf6, 0xa8, 0xd1, 0xed, 0xbb, 0x68, 0xe4, 0x11, 0xc4, 0xad, 0x9b, 0x39, 0xff, 0x4a, 0x1f,
	0xae, 0x0a, 0xc5, 0x6a, 0x33, 0xc6, 0x72, 0x5b, 0x32, 0x93, 0xd0, 0x77, 0x1d, 0x42, 0x3d, 0x95,
	0x8d, 0x9e, 0x4a, 0xb7, 0xb5, 0xd4, 0x62, 0x43, 0x2d, 0x47, 0xcf, 0xe1, 0x51, 0x83, 0x34, 0x0e,
	0x18, 0x6e, 0xb9, 0x07, 0x71, 0x59, 0x5d, 0xac, 0x45, 0xe6, 0x2b, 0x7a, 0x58, 0x01, 0x1e, 0xc2,
	0x02, 0x02, 0x7d, 0x7b, 0x53, 0xf2, 0xa0, 0x31, 0x5c, 0xcf, 0xfe, 0xee, 0x00, 0x6c, 0x59, 0x68,
	0x4a, 0x3a, 0xdb, 0x12, 0x87, 0xb5, 0x86, 0x03, 0xae, 0xb1, 0x57, 0xaa, 0xd2, 0x59, 0xad, 0xe4,
	0x10, 0x91, 0x29, 0xc4, 0x8c, 0x1b, 0x2b, 0x24, 0xfe, 0x68, 0x08, 0x27, 0xb5, 0x21, 0xb7, 0x93,
	0x69, 0xb1, 0xe1, 0x3a, 0x08, 0x3a, 0x44, 0xee, 0x14, 0xe4, 0xc4, 0x0f, 0x07, 0x5c, 0x3b, 0x12,
	0xf4, 0x35, 0xfe, 0x6f, 0x31, 0x4e, 0xbb, 0xda, 0xc9, 0x32, 0x2e, 0xb5, 0x2a, 0x69, 0xee, 0xbf,
	0xee, 0x27, 0x42, 0x1b, 0xfa, 0x6e, 0xfe, 0xdb, 0x83, 0x5c, 0xd8, 0xcb, 0xea, 0xc2, 0x75, 0x78,
	0x91, 0xab, 0x62, 0x8d, 0xff, 0x70, 0xbb, 0xd8, 0xfb, 0xbd, 0x74, 0x31, 0x44, 0xe0, 0xf3, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xb5, 0x5c, 0x5e, 0x49, 0x09, 0x00, 0x00,
}
