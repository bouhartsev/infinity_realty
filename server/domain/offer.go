package domain

type Offer struct {
	Id         int `json:"id"`
	ClientId   int `json:"clientId"`
	RealtorId  int `json:"realtorId"`
	PropertyId int `json:"propertyId"`
	Price      int `json:"price"`
}
