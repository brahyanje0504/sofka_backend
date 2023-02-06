package handlers

import (
	"fmt"
	"github.com/brahyanje0504/sofka_backend/domain/model/entities"
	"github.com/brahyanje0504/sofka_backend/domain/usecase"
	"github.com/brahyanje0504/sofka_backend/infrastructure/entrypoint/rest/dtos"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type SpacecraftHandler struct {
	uc *usecase.SpacecraftUseCase
}

func NewHandler(uc *usecase.SpacecraftUseCase) *SpacecraftHandler {
	return &SpacecraftHandler{
		uc: uc,
	}
}

func (h *SpacecraftHandler) CreateSpacecraft(ctx *fiber.Ctx) error {

	rq := dtos.SpacecraftRQ{}

	err := ctx.BodyParser(&rq)
	if err != nil {
		return ctx.JSON(dtos.RS{
			Status:  false,
			Message: fmt.Sprintf("errror parsing data: %v", err.Error()),
			Data:    nil,
		})
	}

	entitie := mapperDtoTOEntities(rq)

	err = h.uc.Create(ctx.Context(), entitie)
	if err != nil {
		return ctx.JSON(dtos.RS{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return ctx.JSON(dtos.RS{
		Status:  true,
		Message: "successful creation",
		Data:    nil,
	})
}

func (h *SpacecraftHandler) GetAll(ctx *fiber.Ctx) error {
	response, err := h.uc.GetAll(ctx.Context())
	if err != nil {
		return ctx.JSON(dtos.RS{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	r := make([]dtos.SpacecraftRQ, len(response))
	for i := 0; i < len(response); i++ {
		r[i] = mapperEntitiesTODto(response[i])
	}

	return ctx.JSON(dtos.RS{
		Status: true,
		Data:   r,
	})
}

func (h *SpacecraftHandler) GetCountries(ctx *fiber.Ctx) error {
	response, err := h.uc.GetCountries(ctx.Context())
	if err != nil {
		return ctx.JSON(dtos.RS{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(dtos.RS{
		Status: true,
		Data:   response,
	})
}

func (h *SpacecraftHandler) GetFuels(ctx *fiber.Ctx) error {
	response, err := h.uc.GetFuels(ctx.Context())
	if err != nil {
		return ctx.JSON(dtos.RS{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(dtos.RS{
		Status: true,
		Data:   response,
	})
}

func mapperDtoTOEntities(input dtos.SpacecraftRQ) entities.SpacecraftEntities {
	expirationDate := time.Time{}

	if input.ExpirationDate != "" {
		expirationDate, _ = time.Parse("2006-01-02", input.ExpirationDate)
	}

	return entities.SpacecraftEntities{
		ID:                uuid.New().String(),
		TypeSpacecraft:    input.TypeSpacecraft,
		TonsOfThrust:      input.TonsOfThrust,
		TonsOfWeight:      input.TonsOfWeight,
		TypeFuel:          input.TypeFuel,
		Height:            input.Height,
		Country:           input.Country,
		Objective:         input.Objective,
		Activated:         input.Activated,
		PeopleCapacity:    input.PeopleCapacity,
		TransportCapacity: input.TransportCapacity,
		MassiveSpeed:      input.MassiveSpeed,
		PropulsionSystem:  input.PropulsionSystem,
		ExpirationDate:    expirationDate,
		Image:             input.Image,
		Name:              input.Name,
	}
}

func mapperEntitiesTODto(input entities.SpacecraftEntities) dtos.SpacecraftRQ {
	return dtos.SpacecraftRQ{
		TypeSpacecraft:    input.TypeSpacecraft,
		TonsOfThrust:      input.TonsOfThrust,
		TonsOfWeight:      input.TonsOfWeight,
		TypeFuel:          input.TypeFuel,
		Height:            input.Height,
		Country:           input.Country,
		Objective:         input.Objective,
		Activated:         input.Activated,
		PeopleCapacity:    input.PeopleCapacity,
		TransportCapacity: input.TransportCapacity,
		MassiveSpeed:      input.MassiveSpeed,
		PropulsionSystem:  input.PropulsionSystem,
		ExpirationDate:    input.ExpirationDate.String(),
		Image:             input.Image,
		Name:              input.Name,
	}
}
