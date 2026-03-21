package api

import (
	"context"
	"net/http"
	appentity "open-outcry/pkg/models/app_entity"
	transfer "open-outcry/pkg/models/transfer"
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
	return Response(http.StatusOK, TransferEntry{
		Id:                      entry.Id,
		Type:                    TransferType(entry.Type),
		Amount:                  entry.Amount,
		Currency:                string(entry.Currency),
		SenderAccountId:         string(entry.SenderAccountId),
		BeneficiaryAccountId:    string(entry.BeneficiaryAccountId),
		Details:                 string(entry.Details),
		ExternalReferenceNumber: string(entry.ExternalReferenceNumber),
		Status:                  entry.Status,
		DebitBalanceAmount:      entry.DebitBalanceAmount,
		CreditBalanceAmount:     entry.CreditBalanceAmount,
	}), nil
}

func (s *AdminAPIService) GetAppEntities(ctx context.Context) (ImplResponse, error) {
	entities := appentity.GetAppEntities()
	var result []AppEntity
	for _, e := range entities {
		result = append(result, AppEntity{Id: string(e.Id), ExternalId: string(e.ExternalId)})
	}
	return Response(http.StatusOK, result), nil
}

func (s *AdminAPIService) GetAppEntity(ctx context.Context, appEntityId string) (ImplResponse, error) {
	entity := appentity.GetAppEntity(appentity.AppEntityId(appEntityId))
	if entity == nil {
		return Response(http.StatusNotFound, nil), nil
	}
	return Response(http.StatusOK, AppEntity{Id: string(entity.Id), ExternalId: string(entity.ExternalId)}), nil
}
