package domain

type PropertyType int8

const (
	Floor PropertyType = iota + 1
	House
	Field
)

type (
	Property struct {
		Id          int                 `json:"id"`
		Address     PropertyAddress     `json:"address"`
		Coordinates PropertyCoordinates `json:"coordinates"`
		Type        PropertyType        `json:"type"`
		Floor       *int8               `json:"floor"`
		RoomCount   *int8               `json:"roomCount"`
		Square      *int8               `json:"square"`
		FloorCount  *int8               `json:"floorCount"`
	}

	PropertyAddress struct {
		City     string `json:"city"`
		Street   string `json:"street"`
		Building string `json:"building"`
		Floor    string `json:"floor"`
	}

	PropertyCoordinates struct {
		Latitude  float32 `json:"latitude"`
		Longitude float32 `json:"longitude"`
	}
)
