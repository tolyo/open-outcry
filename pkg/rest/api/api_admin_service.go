package api

import (
	"context"
	"net/http"

	appentity "open-outcry/pkg/models/app_entity"
	"open-outcry/pkg/models/transfer"
)

type AdminAPIService struct{}

func NewAdminAPIService() AdminAPIServicer {
	return &AdminAPIService{}
}

func (s *AdminAPIService) CreateAdminTransfer(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusNotImplemented, nil), nil
}

func (s *AdminAPIService) GetAdminTransferById(ctx context.Context, transferId string) (ImplResponse, error) {
	entry := transfer.GetTransfer(transferId)
	if entry == nil {
		return Response(http.StatusNotFound, nil), nil
	}
	return Response(http.StatusOK, mapTransferEntry(*entry)), nil
}

func (s *AdminAPIService) GetAppEntities(ctx context.Context) (ImplResponse, error) {
	return Response(http.StatusOK, mapAppEntities(appentity.GetAppEntities())), nil
}

func (s *AdminAPIService) GetAppEntity(ctx context.Context, appEntityId string) (ImplResponse, error) {
	entity := appentity.GetAppEntity(appentity.AppEntityId(appEntityId))
	if entity == nil {
		return Response(http.StatusNotFound, nil), nil
	}
	return Response(http.StatusOK, mapAppEntity(*entity)), nil
}
