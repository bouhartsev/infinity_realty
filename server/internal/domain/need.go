package domain

type (
	Need struct {
		Id            int             `json:"id"`
		ClientId      int             `json:"clientId"`
		RealtorId     int             `json:"realtorId"`
		PropertyType  int8            `json:"propertyType"`
		Address       PropertyAddress `json:"address,omitempty"`
		MinSquare     *int            `json:"minSquare,omitempty"`
		MaxSquare     *int            `json:"maxSquare,omitempty"`
		MinRooms      *int            `json:"minRooms,omitempty"`
		MaxRooms      *int            `json:"maxRooms,omitempty"`
		MinFloor      *int            `json:"minFloor,omitempty"`
		MaxFloor      *int            `json:"maxFloor,omitempty"`
		MaxFloorCount *int            `json:"maxFloorCount,omitempty"`
		MinFloorCount *int            `json:"minFloorCount,omitempty"`
		MinPrice      *int            `json:"minPrice,omitempty"`
		MaxPrice      *int            `json:"maxPrice,omitempty"`
	}

	CreateNeedRequest struct {
		ClientId      int             `json:"clientId"`
		RealtorId     int             `json:"realtorId"`
		PropertyType  int8            `json:"propertyType"`
		Address       PropertyAddress `json:"address"`
		MinSquare     *int            `json:"minSquare,omitempty"`
		MaxSquare     *int            `json:"maxSquare,omitempty"`
		MinRooms      *int            `json:"minRooms,omitempty"`
		MaxRooms      *int            `json:"maxRooms,omitempty"`
		MinFloor      *int            `json:"minFloor,omitempty"`
		MaxFloor      *int            `json:"maxFloor,omitempty"`
		MaxFloorCount *int            `json:"maxFloorCount,omitempty"`
		MinFloorCount *int            `json:"minFloorCount,omitempty"`
		MinPrice      *int            `json:"minPrice,omitempty"`
		MaxPrice      *int            `json:"maxPrice,omitempty"`
	}

	UpdateNeedRequest struct {
		NeedId        int              `json:"-"`
		ClientId      *int             `json:"clientId"`
		RealtorId     *int             `json:"realtorId"`
		PropertyType  *int8            `json:"propertyType"`
		Address       *PropertyAddress `json:"address"`
		MinSquare     *int             `json:"minSquare,omitempty"`
		MaxSquare     *int             `json:"maxSquare,omitempty"`
		MinRooms      *int             `json:"minRooms,omitempty"`
		MaxRooms      *int             `json:"maxRooms,omitempty"`
		MinFloor      *int             `json:"minFloor,omitempty"`
		MaxFloor      *int             `json:"maxFloor,omitempty"`
		MaxFloorCount *int             `json:"maxFloorCount,omitempty"`
		MinFloorCount *int             `json:"minFloorCount,omitempty"`
		MinPrice      *int             `json:"minPrice,omitempty"`
		MaxPrice      *int             `json:"maxPrice,omitempty"`
	}
)
