package main

import (
	"fmt"
	"github.com/brahyanje0504/sofka_backend/application/rest/application/rest/database"
	"github.com/brahyanje0504/sofka_backend/application/rest/application/rest/server"
	"github.com/brahyanje0504/sofka_backend/domain/usecase"
	"github.com/brahyanje0504/sofka_backend/infrastructure/drivenadapters/adapters"
	"github.com/brahyanje0504/sofka_backend/infrastructure/entrypoint/rest/handlers"
	"github.com/brahyanje0504/sofka_backend/infrastructure/entrypoint/rest/routers"
)

func main() {
	s := server.NewServer()
	db, err := database.NewConectionDatabase()
	if err != nil {
		panic(fmt.Sprintf("error connection database %v", err.Error()))
	}

	drive := adapters.NewAdapter(db)
	uc := usecase.NewSpacecraftUseCase(drive)
	handler := handlers.NewHandler(uc)

	r := routers.NewRouter(s, handler)
	err = r.Start()
	if err != nil {
		panic(err.Error())
	}

}
