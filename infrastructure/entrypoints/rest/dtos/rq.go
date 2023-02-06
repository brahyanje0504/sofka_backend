package dtos

type SpacecraftRQ struct {
	TypeSpacecraft    string  `json:"type_spacecraft"`
	TonsOfThrust      float64 `json:"tons_of_thrust"`
	TonsOfWeight      float64 `json:"tons_of_weight"`
	TypeFuel          string  `json:"type_fuel"`
	Height            int     `json:"height"`
	Country           string  `json:"country"`
	Objective         string  `json:"objective"`
	Activated         bool    `json:"activated"`
	ExpirationDate    string  `json:"expiration_date"`
	PeopleCapacity    int     `json:"people_capacity"`
	TransportCapacity float64 `json:"transport_capacity"`
	MassiveSpeed      float64 `json:"massive_speed"`
	PropulsionSystem  string  `json:"propulsion_system"`
	Image             string  `json:"image"`
	Name              string  `json:"name"`
}
