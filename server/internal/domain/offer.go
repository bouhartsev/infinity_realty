package domain

import "github.com/bouhartsev/infinity_realty/internal/domain/errdomain"

type (
	Offer struct {
		Id         int `json:"id"`
		ClientId   int `json:"clientId"`
		RealtorId  int `json:"realtorId"`
		PropertyId int `json:"propertyId"`
		Price      int `json:"price"`
	}

	CreateOfferRequest struct {
		ClientId   int `json:"clientId"`
		RealtorId  int `json:"realtorId"`
		PropertyId int `json:"propertyId"`
		Price      int `json:"price" minimum:"1"`
	}
)

func (req CreateOfferRequest) Validate() error {
	if req.Price <= 0 {
		return errdomain.NewUserError("Price must be positive number.", "price")
	}
	return nil
}
