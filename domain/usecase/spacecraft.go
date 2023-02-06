package usecase

import (
	"context"
	"github.com/brahyanje0504/sofka_backend/domain/model/entities"
	"github.com/brahyanje0504/sofka_backend/domain/model/gateways"
)

type SpacecraftUseCase struct {
	gt gateways.Spacecraft
}

func NewSpacecraftUseCase(gt gateways.Spacecraft) *SpacecraftUseCase {
	return &SpacecraftUseCase{
		gt: gt,
	}
}

func (h *SpacecraftUseCase) Create(ctx context.Context, input entities.SpacecraftEntities) error {
	return h.gt.CreateSpacecraft(ctx, input)
}

func (h *SpacecraftUseCase) GetAll(ctx context.Context) ([]entities.SpacecraftEntities, error) {
	return h.gt.GetAll(ctx)
}

func (h *SpacecraftUseCase) GetCountries(ctx context.Context) ([]string, error) {
	return h.gt.GetCountries(ctx)
}

func (h *SpacecraftUseCase) GetFuels(ctx context.Context) ([]string, error) {
	return h.gt.GetFuels(ctx)
}
