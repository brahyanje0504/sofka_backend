package gateways

import (
	"context"
	"github.com/brahyanje0504/sofka_backend/domain/model/entities"
)

type Spacecraft interface {
	CreateSpacecraft(ctx context.Context, entity entities.SpacecraftEntities) error
	GetAll(ctx context.Context) ([]entities.SpacecraftEntities, error)
	GetCountries(ctx context.Context) ([]string, error)
	GetFuels(ctx context.Context) ([]string, error)
}
