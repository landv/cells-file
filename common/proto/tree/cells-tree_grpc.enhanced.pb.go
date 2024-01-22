// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             (unknown)
// source: cells-tree.proto

package tree

import (
	context "context"
	fmt "fmt"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	metadata "google.golang.org/grpc/metadata"
	status "google.golang.org/grpc/status"
	sync "sync"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var (
	enhancedNodeProviderServers     = make(map[string]NodeProviderEnhancedServer)
	enhancedNodeProviderServersLock = sync.RWMutex{}
)

type NamedNodeProviderServer interface {
	NodeProviderServer
	Name() string
}
type NodeProviderEnhancedServer map[string]NamedNodeProviderServer

func (m NodeProviderEnhancedServer) ReadNode(ctx context.Context, r *ReadNodeRequest) (*ReadNodeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method ReadNode should have a context")
	}
	enhancedNodeProviderServersLock.RLock()
	defer enhancedNodeProviderServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ReadNode(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method ReadNode not implemented")
}

func (m NodeProviderEnhancedServer) ListNodes(r *ListNodesRequest, s NodeProvider_ListNodesServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method ListNodes should have a context")
	}
	enhancedNodeProviderServersLock.RLock()
	defer enhancedNodeProviderServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ListNodes(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (m NodeProviderEnhancedServer) mustEmbedUnimplementedNodeProviderServer() {}
func RegisterNodeProviderEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeProviderServer) {
	enhancedNodeProviderServersLock.Lock()
	defer enhancedNodeProviderServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeProviderServers[addr]
	if !ok {
		m = NodeProviderEnhancedServer{}
		enhancedNodeProviderServers[addr] = m
		RegisterNodeProviderServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeProviderEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeProviderServersLock.Lock()
	defer enhancedNodeProviderServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeProviderServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeProviderStreamerServers     = make(map[string]NodeProviderStreamerEnhancedServer)
	enhancedNodeProviderStreamerServersLock = sync.RWMutex{}
)

type NamedNodeProviderStreamerServer interface {
	NodeProviderStreamerServer
	Name() string
}
type NodeProviderStreamerEnhancedServer map[string]NamedNodeProviderStreamerServer

func (m NodeProviderStreamerEnhancedServer) ReadNodeStream(s NodeProviderStreamer_ReadNodeStreamServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method ReadNodeStream should have a context")
	}
	enhancedNodeProviderStreamerServersLock.RLock()
	defer enhancedNodeProviderStreamerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ReadNodeStream(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method ReadNodeStream not implemented")
}
func (m NodeProviderStreamerEnhancedServer) mustEmbedUnimplementedNodeProviderStreamerServer() {}
func RegisterNodeProviderStreamerEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeProviderStreamerServer) {
	enhancedNodeProviderStreamerServersLock.Lock()
	defer enhancedNodeProviderStreamerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeProviderStreamerServers[addr]
	if !ok {
		m = NodeProviderStreamerEnhancedServer{}
		enhancedNodeProviderStreamerServers[addr] = m
		RegisterNodeProviderStreamerServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeProviderStreamerEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeProviderStreamerServersLock.Lock()
	defer enhancedNodeProviderStreamerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeProviderStreamerServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeChangesStreamerServers     = make(map[string]NodeChangesStreamerEnhancedServer)
	enhancedNodeChangesStreamerServersLock = sync.RWMutex{}
)

type NamedNodeChangesStreamerServer interface {
	NodeChangesStreamerServer
	Name() string
}
type NodeChangesStreamerEnhancedServer map[string]NamedNodeChangesStreamerServer

func (m NodeChangesStreamerEnhancedServer) StreamChanges(r *StreamChangesRequest, s NodeChangesStreamer_StreamChangesServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method StreamChanges should have a context")
	}
	enhancedNodeChangesStreamerServersLock.RLock()
	defer enhancedNodeChangesStreamerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StreamChanges(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method StreamChanges not implemented")
}
func (m NodeChangesStreamerEnhancedServer) mustEmbedUnimplementedNodeChangesStreamerServer() {}
func RegisterNodeChangesStreamerEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeChangesStreamerServer) {
	enhancedNodeChangesStreamerServersLock.Lock()
	defer enhancedNodeChangesStreamerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeChangesStreamerServers[addr]
	if !ok {
		m = NodeChangesStreamerEnhancedServer{}
		enhancedNodeChangesStreamerServers[addr] = m
		RegisterNodeChangesStreamerServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeChangesStreamerEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeChangesStreamerServersLock.Lock()
	defer enhancedNodeChangesStreamerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeChangesStreamerServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeChangesReceiverStreamerServers     = make(map[string]NodeChangesReceiverStreamerEnhancedServer)
	enhancedNodeChangesReceiverStreamerServersLock = sync.RWMutex{}
)

type NamedNodeChangesReceiverStreamerServer interface {
	NodeChangesReceiverStreamerServer
	Name() string
}
type NodeChangesReceiverStreamerEnhancedServer map[string]NamedNodeChangesReceiverStreamerServer

func (m NodeChangesReceiverStreamerEnhancedServer) PostNodeChanges(s NodeChangesReceiverStreamer_PostNodeChangesServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method PostNodeChanges should have a context")
	}
	enhancedNodeChangesReceiverStreamerServersLock.RLock()
	defer enhancedNodeChangesReceiverStreamerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.PostNodeChanges(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method PostNodeChanges not implemented")
}
func (m NodeChangesReceiverStreamerEnhancedServer) mustEmbedUnimplementedNodeChangesReceiverStreamerServer() {
}
func RegisterNodeChangesReceiverStreamerEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeChangesReceiverStreamerServer) {
	enhancedNodeChangesReceiverStreamerServersLock.Lock()
	defer enhancedNodeChangesReceiverStreamerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeChangesReceiverStreamerServers[addr]
	if !ok {
		m = NodeChangesReceiverStreamerEnhancedServer{}
		enhancedNodeChangesReceiverStreamerServers[addr] = m
		RegisterNodeChangesReceiverStreamerServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeChangesReceiverStreamerEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeChangesReceiverStreamerServersLock.Lock()
	defer enhancedNodeChangesReceiverStreamerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeChangesReceiverStreamerServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeReceiverServers     = make(map[string]NodeReceiverEnhancedServer)
	enhancedNodeReceiverServersLock = sync.RWMutex{}
)

type NamedNodeReceiverServer interface {
	NodeReceiverServer
	Name() string
}
type NodeReceiverEnhancedServer map[string]NamedNodeReceiverServer

func (m NodeReceiverEnhancedServer) CreateNode(ctx context.Context, r *CreateNodeRequest) (*CreateNodeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CreateNode should have a context")
	}
	enhancedNodeReceiverServersLock.RLock()
	defer enhancedNodeReceiverServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateNode(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CreateNode not implemented")
}

func (m NodeReceiverEnhancedServer) UpdateNode(ctx context.Context, r *UpdateNodeRequest) (*UpdateNodeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method UpdateNode should have a context")
	}
	enhancedNodeReceiverServersLock.RLock()
	defer enhancedNodeReceiverServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.UpdateNode(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNode not implemented")
}

func (m NodeReceiverEnhancedServer) DeleteNode(ctx context.Context, r *DeleteNodeRequest) (*DeleteNodeResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method DeleteNode should have a context")
	}
	enhancedNodeReceiverServersLock.RLock()
	defer enhancedNodeReceiverServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteNode(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNode not implemented")
}
func (m NodeReceiverEnhancedServer) mustEmbedUnimplementedNodeReceiverServer() {}
func RegisterNodeReceiverEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeReceiverServer) {
	enhancedNodeReceiverServersLock.Lock()
	defer enhancedNodeReceiverServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeReceiverServers[addr]
	if !ok {
		m = NodeReceiverEnhancedServer{}
		enhancedNodeReceiverServers[addr] = m
		RegisterNodeReceiverServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeReceiverEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeReceiverServersLock.Lock()
	defer enhancedNodeReceiverServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeReceiverServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeReceiverStreamServers     = make(map[string]NodeReceiverStreamEnhancedServer)
	enhancedNodeReceiverStreamServersLock = sync.RWMutex{}
)

type NamedNodeReceiverStreamServer interface {
	NodeReceiverStreamServer
	Name() string
}
type NodeReceiverStreamEnhancedServer map[string]NamedNodeReceiverStreamServer

func (m NodeReceiverStreamEnhancedServer) CreateNodeStream(s NodeReceiverStream_CreateNodeStreamServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method CreateNodeStream should have a context")
	}
	enhancedNodeReceiverStreamServersLock.RLock()
	defer enhancedNodeReceiverStreamServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateNodeStream(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method CreateNodeStream not implemented")
}

func (m NodeReceiverStreamEnhancedServer) UpdateNodeStream(s NodeReceiverStream_UpdateNodeStreamServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method UpdateNodeStream should have a context")
	}
	enhancedNodeReceiverStreamServersLock.RLock()
	defer enhancedNodeReceiverStreamServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.UpdateNodeStream(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method UpdateNodeStream not implemented")
}

func (m NodeReceiverStreamEnhancedServer) DeleteNodeStream(s NodeReceiverStream_DeleteNodeStreamServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method DeleteNodeStream should have a context")
	}
	enhancedNodeReceiverStreamServersLock.RLock()
	defer enhancedNodeReceiverStreamServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.DeleteNodeStream(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method DeleteNodeStream not implemented")
}
func (m NodeReceiverStreamEnhancedServer) mustEmbedUnimplementedNodeReceiverStreamServer() {}
func RegisterNodeReceiverStreamEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeReceiverStreamServer) {
	enhancedNodeReceiverStreamServersLock.Lock()
	defer enhancedNodeReceiverStreamServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeReceiverStreamServers[addr]
	if !ok {
		m = NodeReceiverStreamEnhancedServer{}
		enhancedNodeReceiverStreamServers[addr] = m
		RegisterNodeReceiverStreamServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeReceiverStreamEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeReceiverStreamServersLock.Lock()
	defer enhancedNodeReceiverStreamServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeReceiverStreamServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedSessionIndexerServers     = make(map[string]SessionIndexerEnhancedServer)
	enhancedSessionIndexerServersLock = sync.RWMutex{}
)

type NamedSessionIndexerServer interface {
	SessionIndexerServer
	Name() string
}
type SessionIndexerEnhancedServer map[string]NamedSessionIndexerServer

func (m SessionIndexerEnhancedServer) OpenSession(ctx context.Context, r *OpenSessionRequest) (*OpenSessionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method OpenSession should have a context")
	}
	enhancedSessionIndexerServersLock.RLock()
	defer enhancedSessionIndexerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.OpenSession(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method OpenSession not implemented")
}

func (m SessionIndexerEnhancedServer) FlushSession(ctx context.Context, r *FlushSessionRequest) (*FlushSessionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method FlushSession should have a context")
	}
	enhancedSessionIndexerServersLock.RLock()
	defer enhancedSessionIndexerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.FlushSession(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method FlushSession not implemented")
}

func (m SessionIndexerEnhancedServer) CloseSession(ctx context.Context, r *CloseSessionRequest) (*CloseSessionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CloseSession should have a context")
	}
	enhancedSessionIndexerServersLock.RLock()
	defer enhancedSessionIndexerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CloseSession(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CloseSession not implemented")
}
func (m SessionIndexerEnhancedServer) mustEmbedUnimplementedSessionIndexerServer() {}
func RegisterSessionIndexerEnhancedServer(s grpc.ServiceRegistrar, srv NamedSessionIndexerServer) {
	enhancedSessionIndexerServersLock.Lock()
	defer enhancedSessionIndexerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedSessionIndexerServers[addr]
	if !ok {
		m = SessionIndexerEnhancedServer{}
		enhancedSessionIndexerServers[addr] = m
		RegisterSessionIndexerServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterSessionIndexerEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedSessionIndexerServersLock.Lock()
	defer enhancedSessionIndexerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedSessionIndexerServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeEventsProviderServers     = make(map[string]NodeEventsProviderEnhancedServer)
	enhancedNodeEventsProviderServersLock = sync.RWMutex{}
)

type NamedNodeEventsProviderServer interface {
	NodeEventsProviderServer
	Name() string
}
type NodeEventsProviderEnhancedServer map[string]NamedNodeEventsProviderServer

func (m NodeEventsProviderEnhancedServer) WatchNode(r *WatchNodeRequest, s NodeEventsProvider_WatchNodeServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method WatchNode should have a context")
	}
	enhancedNodeEventsProviderServersLock.RLock()
	defer enhancedNodeEventsProviderServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.WatchNode(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method WatchNode not implemented")
}
func (m NodeEventsProviderEnhancedServer) mustEmbedUnimplementedNodeEventsProviderServer() {}
func RegisterNodeEventsProviderEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeEventsProviderServer) {
	enhancedNodeEventsProviderServersLock.Lock()
	defer enhancedNodeEventsProviderServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeEventsProviderServers[addr]
	if !ok {
		m = NodeEventsProviderEnhancedServer{}
		enhancedNodeEventsProviderServers[addr] = m
		RegisterNodeEventsProviderServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeEventsProviderEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeEventsProviderServersLock.Lock()
	defer enhancedNodeEventsProviderServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeEventsProviderServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedSearcherServers     = make(map[string]SearcherEnhancedServer)
	enhancedSearcherServersLock = sync.RWMutex{}
)

type NamedSearcherServer interface {
	SearcherServer
	Name() string
}
type SearcherEnhancedServer map[string]NamedSearcherServer

func (m SearcherEnhancedServer) Search(r *SearchRequest, s Searcher_SearchServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method Search should have a context")
	}
	enhancedSearcherServersLock.RLock()
	defer enhancedSearcherServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.Search(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (m SearcherEnhancedServer) mustEmbedUnimplementedSearcherServer() {}
func RegisterSearcherEnhancedServer(s grpc.ServiceRegistrar, srv NamedSearcherServer) {
	enhancedSearcherServersLock.Lock()
	defer enhancedSearcherServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedSearcherServers[addr]
	if !ok {
		m = SearcherEnhancedServer{}
		enhancedSearcherServers[addr] = m
		RegisterSearcherServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterSearcherEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedSearcherServersLock.Lock()
	defer enhancedSearcherServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedSearcherServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeContentReaderServers     = make(map[string]NodeContentReaderEnhancedServer)
	enhancedNodeContentReaderServersLock = sync.RWMutex{}
)

type NamedNodeContentReaderServer interface {
	NodeContentReaderServer
	Name() string
}
type NodeContentReaderEnhancedServer map[string]NamedNodeContentReaderServer

func (m NodeContentReaderEnhancedServer) mustEmbedUnimplementedNodeContentReaderServer() {}
func RegisterNodeContentReaderEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeContentReaderServer) {
	enhancedNodeContentReaderServersLock.Lock()
	defer enhancedNodeContentReaderServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeContentReaderServers[addr]
	if !ok {
		m = NodeContentReaderEnhancedServer{}
		enhancedNodeContentReaderServers[addr] = m
		RegisterNodeContentReaderServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeContentReaderEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeContentReaderServersLock.Lock()
	defer enhancedNodeContentReaderServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeContentReaderServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeContentWriterServers     = make(map[string]NodeContentWriterEnhancedServer)
	enhancedNodeContentWriterServersLock = sync.RWMutex{}
)

type NamedNodeContentWriterServer interface {
	NodeContentWriterServer
	Name() string
}
type NodeContentWriterEnhancedServer map[string]NamedNodeContentWriterServer

func (m NodeContentWriterEnhancedServer) mustEmbedUnimplementedNodeContentWriterServer() {}
func RegisterNodeContentWriterEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeContentWriterServer) {
	enhancedNodeContentWriterServersLock.Lock()
	defer enhancedNodeContentWriterServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeContentWriterServers[addr]
	if !ok {
		m = NodeContentWriterEnhancedServer{}
		enhancedNodeContentWriterServers[addr] = m
		RegisterNodeContentWriterServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeContentWriterEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeContentWriterServersLock.Lock()
	defer enhancedNodeContentWriterServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeContentWriterServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedNodeVersionerServers     = make(map[string]NodeVersionerEnhancedServer)
	enhancedNodeVersionerServersLock = sync.RWMutex{}
)

type NamedNodeVersionerServer interface {
	NodeVersionerServer
	Name() string
}
type NodeVersionerEnhancedServer map[string]NamedNodeVersionerServer

func (m NodeVersionerEnhancedServer) CreateVersion(ctx context.Context, r *CreateVersionRequest) (*CreateVersionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method CreateVersion should have a context")
	}
	enhancedNodeVersionerServersLock.RLock()
	defer enhancedNodeVersionerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.CreateVersion(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method CreateVersion not implemented")
}

func (m NodeVersionerEnhancedServer) StoreVersion(ctx context.Context, r *StoreVersionRequest) (*StoreVersionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method StoreVersion should have a context")
	}
	enhancedNodeVersionerServersLock.RLock()
	defer enhancedNodeVersionerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.StoreVersion(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method StoreVersion not implemented")
}

func (m NodeVersionerEnhancedServer) ListVersions(r *ListVersionsRequest, s NodeVersioner_ListVersionsServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method ListVersions should have a context")
	}
	enhancedNodeVersionerServersLock.RLock()
	defer enhancedNodeVersionerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.ListVersions(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method ListVersions not implemented")
}

func (m NodeVersionerEnhancedServer) HeadVersion(ctx context.Context, r *HeadVersionRequest) (*HeadVersionResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method HeadVersion should have a context")
	}
	enhancedNodeVersionerServersLock.RLock()
	defer enhancedNodeVersionerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.HeadVersion(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method HeadVersion not implemented")
}

func (m NodeVersionerEnhancedServer) PruneVersions(ctx context.Context, r *PruneVersionsRequest) (*PruneVersionsResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method PruneVersions should have a context")
	}
	enhancedNodeVersionerServersLock.RLock()
	defer enhancedNodeVersionerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.PruneVersions(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method PruneVersions not implemented")
}
func (m NodeVersionerEnhancedServer) mustEmbedUnimplementedNodeVersionerServer() {}
func RegisterNodeVersionerEnhancedServer(s grpc.ServiceRegistrar, srv NamedNodeVersionerServer) {
	enhancedNodeVersionerServersLock.Lock()
	defer enhancedNodeVersionerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeVersionerServers[addr]
	if !ok {
		m = NodeVersionerEnhancedServer{}
		enhancedNodeVersionerServers[addr] = m
		RegisterNodeVersionerServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterNodeVersionerEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedNodeVersionerServersLock.Lock()
	defer enhancedNodeVersionerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedNodeVersionerServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedFileKeyManagerServers     = make(map[string]FileKeyManagerEnhancedServer)
	enhancedFileKeyManagerServersLock = sync.RWMutex{}
)

type NamedFileKeyManagerServer interface {
	FileKeyManagerServer
	Name() string
}
type FileKeyManagerEnhancedServer map[string]NamedFileKeyManagerServer

func (m FileKeyManagerEnhancedServer) GetEncryptionKey(ctx context.Context, r *GetEncryptionKeyRequest) (*GetEncryptionKeyResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("targetname")) == 0 {
		return nil, status.Errorf(codes.FailedPrecondition, "method GetEncryptionKey should have a context")
	}
	enhancedFileKeyManagerServersLock.RLock()
	defer enhancedFileKeyManagerServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.GetEncryptionKey(ctx, r)
		}
	}
	return nil, status.Errorf(codes.Unimplemented, "method GetEncryptionKey not implemented")
}
func (m FileKeyManagerEnhancedServer) mustEmbedUnimplementedFileKeyManagerServer() {}
func RegisterFileKeyManagerEnhancedServer(s grpc.ServiceRegistrar, srv NamedFileKeyManagerServer) {
	enhancedFileKeyManagerServersLock.Lock()
	defer enhancedFileKeyManagerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedFileKeyManagerServers[addr]
	if !ok {
		m = FileKeyManagerEnhancedServer{}
		enhancedFileKeyManagerServers[addr] = m
		RegisterFileKeyManagerServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterFileKeyManagerEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedFileKeyManagerServersLock.Lock()
	defer enhancedFileKeyManagerServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedFileKeyManagerServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}

var (
	enhancedSyncChangesServers     = make(map[string]SyncChangesEnhancedServer)
	enhancedSyncChangesServersLock = sync.RWMutex{}
)

type NamedSyncChangesServer interface {
	SyncChangesServer
	Name() string
}
type SyncChangesEnhancedServer map[string]NamedSyncChangesServer

func (m SyncChangesEnhancedServer) Put(s SyncChanges_PutServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method Put should have a context")
	}
	enhancedSyncChangesServersLock.RLock()
	defer enhancedSyncChangesServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.Put(s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method Put not implemented")
}

func (m SyncChangesEnhancedServer) Search(r *SearchSyncChangeRequest, s SyncChanges_SearchServer) error {
	md, ok := metadata.FromIncomingContext(s.Context())
	if !ok || len(md.Get("targetname")) == 0 {
		return status.Errorf(codes.FailedPrecondition, "method Search should have a context")
	}
	enhancedSyncChangesServersLock.RLock()
	defer enhancedSyncChangesServersLock.RUnlock()
	for _, mm := range m {
		if mm.Name() == md.Get("targetname")[0] {
			return mm.Search(r, s)
		}
	}
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (m SyncChangesEnhancedServer) mustEmbedUnimplementedSyncChangesServer() {}
func RegisterSyncChangesEnhancedServer(s grpc.ServiceRegistrar, srv NamedSyncChangesServer) {
	enhancedSyncChangesServersLock.Lock()
	defer enhancedSyncChangesServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedSyncChangesServers[addr]
	if !ok {
		m = SyncChangesEnhancedServer{}
		enhancedSyncChangesServers[addr] = m
		RegisterSyncChangesServer(s, m)
	}
	m[srv.Name()] = srv
}
func DeregisterSyncChangesEnhancedServer(s grpc.ServiceRegistrar, name string) {
	enhancedSyncChangesServersLock.Lock()
	defer enhancedSyncChangesServersLock.Unlock()
	addr := fmt.Sprintf("%p", s)
	m, ok := enhancedSyncChangesServers[addr]
	if !ok {
		return
	}
	delete(m, name)
}
