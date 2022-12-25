package domain

// Need - потребность
type Need struct {
	Id            int             `json:"id"`
	ClientId      int             `json:"clientId"`
	RealtorId     int             `json:"realtorId"`
	Address       PropertyAddress `json:"address"`
	PropertyType  int8            `json:"propertyType"`
	MinSquare     *int            `json:"minSquare"`
	MaxSquare     *int            `json:"maxSquare"`
	MinRooms      *int            `json:"minRooms"`
	MaxRooms      *int            `json:"maxRooms"`
	MinFloor      *int            `json:"minFloor"`
	MaxFloor      *int            `json:"maxFloor"`
	MaxFloorCount *int            `json:"maxFloorCount"`
	MinFloorCount *int            `json:"minFloorCount"`
	MinPrice      int             `json:"minPrice"`
	MaxPrice      int             `json:"maxPrice"`
}
