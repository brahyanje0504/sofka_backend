package entities

import "time"

type SpacecraftEntities struct {
	ID                string
	TypeSpacecraft    string
	TonsOfThrust      float64
	TonsOfWeight      float64
	TypeFuel          string
	Height            float64
	Country           string
	Objective         string
	Activated         bool
	ExpirationDate    time.Time
	PeopleCapacity    float64
	TransportCapacity float64
	MassiveSpeed      float64
	PropulsionSystem  string
	Image             string
	Name              string
}
