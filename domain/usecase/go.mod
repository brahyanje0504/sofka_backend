module github.com/brahyanje0504/sofka_backend/domain/usecase

replace (
	github.com/brahyanje0504/sofka_backend/domain/model => ../model
)

require (
	github.com/brahyanje0504/sofka_backend/domain/model v0.0.0
)

go 1.19
