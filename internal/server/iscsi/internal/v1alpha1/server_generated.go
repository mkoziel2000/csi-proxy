// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	"github.com/kubernetes-csi/csi-proxy/client/api/iscsi/v1alpha1"
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	"github.com/kubernetes-csi/csi-proxy/internal/server/iscsi/internal"
	"google.golang.org/grpc"
)

var version = apiversion.NewVersionOrPanic("v1alpha1")

type versionedAPI struct {
	apiGroupServer internal.ServerInterface
}

func NewVersionedServer(apiGroupServer internal.ServerInterface) internal.VersionedAPI {
	return &versionedAPI{
		apiGroupServer: apiGroupServer,
	}
}

func (s *versionedAPI) Register(grpcServer *grpc.Server) {
	v1alpha1.RegisterIscsiServer(grpcServer, s)
}

func (s *versionedAPI) AddTargetPortal(context context.Context, versionedRequest *v1alpha1.AddTargetPortalRequest) (*v1alpha1.AddTargetPortalResponse, error) {
	request := &internal.AddTargetPortalRequest{}
	if err := Convert_v1alpha1_AddTargetPortalRequest_To_internal_AddTargetPortalRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.AddTargetPortal(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.AddTargetPortalResponse{}
	if err := Convert_internal_AddTargetPortalResponse_To_v1alpha1_AddTargetPortalResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) ConnectTarget(context context.Context, versionedRequest *v1alpha1.ConnectTargetRequest) (*v1alpha1.ConnectTargetResponse, error) {
	request := &internal.ConnectTargetRequest{}
	if err := Convert_v1alpha1_ConnectTargetRequest_To_internal_ConnectTargetRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ConnectTarget(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.ConnectTargetResponse{}
	if err := Convert_internal_ConnectTargetResponse_To_v1alpha1_ConnectTargetResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) DisconnectTarget(context context.Context, versionedRequest *v1alpha1.DisconnectTargetRequest) (*v1alpha1.DisconnectTargetResponse, error) {
	request := &internal.DisconnectTargetRequest{}
	if err := Convert_v1alpha1_DisconnectTargetRequest_To_internal_DisconnectTargetRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.DisconnectTarget(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.DisconnectTargetResponse{}
	if err := Convert_internal_DisconnectTargetResponse_To_v1alpha1_DisconnectTargetResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) DiscoverTargetPortal(context context.Context, versionedRequest *v1alpha1.DiscoverTargetPortalRequest) (*v1alpha1.DiscoverTargetPortalResponse, error) {
	request := &internal.DiscoverTargetPortalRequest{}
	if err := Convert_v1alpha1_DiscoverTargetPortalRequest_To_internal_DiscoverTargetPortalRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.DiscoverTargetPortal(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.DiscoverTargetPortalResponse{}
	if err := Convert_internal_DiscoverTargetPortalResponse_To_v1alpha1_DiscoverTargetPortalResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) GetTargetDisks(context context.Context, versionedRequest *v1alpha1.GetTargetDisksRequest) (*v1alpha1.GetTargetDisksResponse, error) {
	request := &internal.GetTargetDisksRequest{}
	if err := Convert_v1alpha1_GetTargetDisksRequest_To_internal_GetTargetDisksRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.GetTargetDisks(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.GetTargetDisksResponse{}
	if err := Convert_internal_GetTargetDisksResponse_To_v1alpha1_GetTargetDisksResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) ListTargetPortals(context context.Context, versionedRequest *v1alpha1.ListTargetPortalsRequest) (*v1alpha1.ListTargetPortalsResponse, error) {
	request := &internal.ListTargetPortalsRequest{}
	if err := Convert_v1alpha1_ListTargetPortalsRequest_To_internal_ListTargetPortalsRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.ListTargetPortals(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.ListTargetPortalsResponse{}
	if err := Convert_internal_ListTargetPortalsResponse_To_v1alpha1_ListTargetPortalsResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}

func (s *versionedAPI) RemoveTargetPortal(context context.Context, versionedRequest *v1alpha1.RemoveTargetPortalRequest) (*v1alpha1.RemoveTargetPortalResponse, error) {
	request := &internal.RemoveTargetPortalRequest{}
	if err := Convert_v1alpha1_RemoveTargetPortalRequest_To_internal_RemoveTargetPortalRequest(versionedRequest, request); err != nil {
		return nil, err
	}

	response, err := s.apiGroupServer.RemoveTargetPortal(context, request, version)
	if err != nil {
		return nil, err
	}

	versionedResponse := &v1alpha1.RemoveTargetPortalResponse{}
	if err := Convert_internal_RemoveTargetPortalResponse_To_v1alpha1_RemoveTargetPortalResponse(response, versionedResponse); err != nil {
		return nil, err
	}

	return versionedResponse, err
}
