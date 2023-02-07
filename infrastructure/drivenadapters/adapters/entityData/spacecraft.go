package entityData

import "time"

type SpacecraftData struct {
	ID                string    `bson:"id"`
	TypeSpacecraft    string    `bson:"typeSpacecraft"`
	TonsOfThrust      float64   `bson:"tonsOfThrust"`
	TonsOfWeight      float64   `bson:"tonsOfWeight"`
	TypeFuel          string    `bson:"typeFuel"`
	Height            float64   `bson:"height"`
	Country           string    `bson:"country"`
	Objective         string    `bson:"objective"`
	Activated         bool      `bson:"activated"`
	ExpirationDate    time.Time `bson:"expirationDate"`
	PeopleCapacity    float64   `bson:"peopleCapacity"`
	TransportCapacity float64   `bson:"transportCapacity"`
	MassiveSpeed      float64   `bson:"massiveSpeed"`
	PropulsionSystem  string    `bson:"propulsionSystem"`
	Image             string    `bson:"image"`
	Name              string    `bson:"name"`
}
