package adapters

import (
	"context"
	"github.com/brahyanje0504/sofka_backend/domain/model/entities"
	"github.com/brahyanje0504/sofka_backend/infrastructure/drivenadapters/adapters/entityData"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adapter struct {
	db *mongo.Collection
}

func NewAdapter(db *mongo.Database) *Adapter {
	return &Adapter{
		db: db.Collection("spaceships"),
	}
}

func (h *Adapter) CreateSpacecraft(ctx context.Context, entity entities.SpacecraftEntities) error {
	data := mapperDomainEntitiesToEntityData(entity)
	_, err := h.db.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (h *Adapter) GetAll(ctx context.Context) ([]entities.SpacecraftEntities, error) {
	filter := bson.D{}
	curs, err := h.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	responseData := []entityData.SpacecraftData{}
	err = curs.All(ctx, &responseData)

	if err != nil {
		return nil, err
	}

	r := make([]entities.SpacecraftEntities, len(responseData))
	for i := 0; i < len(responseData); i++ {
		r[i] = mapperEntityDataToDomainEntities(responseData[i])
	}

	return r, nil

}

func (h *Adapter) GetCountries(ctx context.Context) ([]string, error) {
	filter := bson.D{}
	curs, err := h.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var responseData []entityData.SpacecraftData
	err = curs.All(ctx, &responseData)

	if err != nil {
		return nil, err
	}

	r := map[string]bool{}
	for i := 0; i < len(responseData); i++ {
		if responseData[i].Country != "" {
			r[responseData[i].Country] = true
		}
	}

	var response []string
	for k, _ := range r {
		response = append(response, k)
	}

	return response, nil

}

func (h *Adapter) GetFuels(ctx context.Context) ([]string, error) {
	filter := bson.D{}
	curs, err := h.db.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var responseData []entityData.SpacecraftData
	err = curs.All(ctx, &responseData)

	if err != nil {
		return nil, err
	}

	r := map[string]bool{}
	for i := 0; i < len(responseData); i++ {
		if responseData[i].TypeFuel != "" {
			r[responseData[i].TypeFuel] = true
		}
	}

	var response []string
	for k, _ := range r {
		response = append(response, k)
	}

	return response, nil

}

func mapperDomainEntitiesToEntityData(entity entities.SpacecraftEntities) entityData.SpacecraftData {
	return entityData.SpacecraftData{
		ID:                entity.ID,
		TypeSpacecraft:    entity.TypeSpacecraft,
		TonsOfThrust:      entity.TonsOfThrust,
		TonsOfWeight:      entity.TonsOfWeight,
		TypeFuel:          entity.TypeFuel,
		Height:            entity.Height,
		Country:           entity.Country,
		Objective:         entity.Objective,
		Activated:         entity.Activated,
		ExpirationDate:    entity.ExpirationDate,
		PeopleCapacity:    entity.PeopleCapacity,
		TransportCapacity: entity.TransportCapacity,
		MassiveSpeed:      entity.MassiveSpeed,
		PropulsionSystem:  entity.PropulsionSystem,
		Image:             entity.Image,
		Name:              entity.Name,
	}
}

func mapperEntityDataToDomainEntities(entity entityData.SpacecraftData) entities.SpacecraftEntities {
	return entities.SpacecraftEntities{
		ID:                entity.ID,
		TypeSpacecraft:    entity.TypeSpacecraft,
		TonsOfThrust:      entity.TonsOfThrust,
		TonsOfWeight:      entity.TonsOfWeight,
		TypeFuel:          entity.TypeFuel,
		Height:            entity.Height,
		Country:           entity.Country,
		Objective:         entity.Objective,
		Activated:         entity.Activated,
		ExpirationDate:    entity.ExpirationDate,
		PeopleCapacity:    entity.PeopleCapacity,
		TransportCapacity: entity.TransportCapacity,
		MassiveSpeed:      entity.MassiveSpeed,
		PropulsionSystem:  entity.PropulsionSystem,
		Image:             entity.Image,
		Name:              entity.Name,
	}
}
