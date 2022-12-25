package domain

import "github.com/bouhartsev/infinity_realty/internal/domain/errdomain"

type (
	Property struct {
		Id          int                 `json:"id"`
		Type        int8                `json:"type"` // 1 - Apartment, 2 - House, 3 - Field
		Address     PropertyAddress     `json:"address"`
		Coordinates PropertyCoordinates `json:"coordinates"`
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

	UpdatePropertyAddress struct {
		City     *string `json:"city"`
		Street   *string `json:"street"`
		Building *string `json:"building"`
		Floor    *string `json:"floor"`
	}

	PropertyCoordinates struct {
		Latitude  float32 `json:"latitude"  minimum:"-90" maximum:"90"`
		Longitude float32 `json:"longitude"  minimum:"-180" maximum:"180"`
	}

	UpdatePropertyCoordinates struct {
		Latitude  *float32 `json:"latitude"  minimum:"-90" maximum:"90"`
		Longitude *float32 `json:"longitude"  minimum:"-180" maximum:"180"`
	}

	CreatePropertyRequest struct {
		Type        int8                 `json:"type" validate:"required" enums:"1,2,3"` // 1 - Floor, 2 - House, 3 - Field
		Address     *PropertyAddress     `json:"address"`
		Coordinates *PropertyCoordinates `json:"coordinates"`
		Floor       *int8                `json:"floor" minimum:"1"`
		RoomCount   *int8                `json:"roomCount"  minimum:"1"`
		Square      *int8                `json:"square"  minimum:"1"`
		FloorCount  *int8                `json:"floorCount"  minimum:"0"`
	}

	UpdatePropertyRequest struct {
		PropertyId  int                        `json:"-"`
		Type        *int8                      `json:"type" enums:"1,2,3"` // 1 - Floor, 2 - House, 3 - Field
		Address     *UpdatePropertyAddress     `json:"address"`
		Coordinates *UpdatePropertyCoordinates `json:"coordinates"`
		Floor       *int8                      `json:"floor" minimum:"1"`
		RoomCount   *int8                      `json:"roomCount"  minimum:"1"`
		Square      *int8                      `json:"square"  minimum:"1"`
		FloorCount  *int8                      `json:"floorCount"  minimum:"0"`
	}

	GetPropertyResponse struct {
		Property Property `json:"property"`
	}
)

func (req CreatePropertyRequest) Validate() error {
	if req.Coordinates.Latitude < -90 || req.Coordinates.Latitude > 90 {
		return errdomain.NewUserError("Latitude must be between -90 and 90.", "coordinates.latitude")
	}

	if req.Coordinates.Longitude < -180 || req.Coordinates.Longitude > 180 {
		return errdomain.NewUserError("Longitude must be between -180 and 180.", "coordinates.longitude")
	}

	switch req.Type {
	case 1: // Floor
		if req.FloorCount != nil {
			return errdomain.NewUserError("Floor count is not needed for Apartment property type.", "floor_count")
		}
	case 2: // House
		if req.Floor != nil {
			return errdomain.NewUserError("Floor is not needed for House property type.", "floor")
		}
	case 3: // Field
		if req.Floor != nil {
			return errdomain.NewUserError("Floor is not needed for Field property type.", "floor")
		}
		if req.FloorCount != nil {
			return errdomain.NewUserError("Floor count is not needed for Field property type.", "floor_count")
		}
		if req.RoomCount != nil {
			return errdomain.NewUserError("Room count is not needed for Field property type.", "room_count")
		}
	default:
		return errdomain.NewUserError("Unknown property type.", "type")
	}

	if req.Floor != nil && *req.Floor < 1 {
		return errdomain.NewUserError("Invalid floor value.", "floor")
	}

	if req.FloorCount != nil && *req.FloorCount < 1 {
		return errdomain.NewUserError("Invalid floor count value.", "floor_count")
	}

	if req.RoomCount != nil && *req.RoomCount < 1 {
		return errdomain.NewUserError("Invalid room count value.", "room_count")
	}

	if req.Square != nil && *req.Square < 0 {
		return errdomain.NewUserError("Invalid square value.", "square")
	}

	return nil
}

func (req UpdatePropertyRequest) Validate() error {
	if req.Coordinates.Latitude != nil && (*req.Coordinates.Latitude < -90 || *req.Coordinates.Latitude > 90) {
		return errdomain.NewUserError("Latitude must be between -90 and 90.", "coordinates.latitude")
	}

	if req.Coordinates.Longitude != nil && (*req.Coordinates.Longitude < -180 || *req.Coordinates.Longitude > 180) {
		return errdomain.NewUserError("Longitude must be between -180 and 180.", "coordinates.longitude")
	}

	if req.Type != nil {
		switch *req.Type {
		case 1: // Floor
			if req.FloorCount != nil {
				return errdomain.NewUserError("Floor count is not needed for Apartment property type.", "floor_count")
			}
		case 2: // House
			if req.Floor != nil {
				return errdomain.NewUserError("Floor is not needed for House property type.", "floor")
			}
		case 3: // Field
			if req.Floor != nil {
				return errdomain.NewUserError("Floor is not needed for Field property type.", "floor")
			}
			if req.FloorCount != nil {
				return errdomain.NewUserError("Floor count is not needed for Field property type.", "floor_count")
			}
			if req.RoomCount != nil {
				return errdomain.NewUserError("Room count is not needed for Field property type.", "room_count")
			}
		default:
			return errdomain.NewUserError("Unknown property type.", "type")
		}
	}

	if req.Floor != nil && *req.Floor < 1 {
		return errdomain.NewUserError("Invalid floor value.", "floor")
	}

	if req.FloorCount != nil && *req.FloorCount < 1 {
		return errdomain.NewUserError("Invalid floor count value.", "floor_count")
	}

	if req.RoomCount != nil && *req.RoomCount < 1 {
		return errdomain.NewUserError("Invalid room count value.", "room_count")
	}

	if req.Square != nil && *req.Square < 0 {
		return errdomain.NewUserError("Invalid square value.", "square")
	}

	return nil
}
