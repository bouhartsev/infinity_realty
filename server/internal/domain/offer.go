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

	UpdateOfferRequest struct {
		OfferId    int  `json:"-"`
		ClientId   *int `json:"clientId"`
		RealtorId  *int `json:"realtorId"`
		PropertyId *int `json:"propertyId"`
		Price      *int `json:"price" minimum:"1"`
	}

	GetOfferResponse struct {
		Offer Offer `json:"offer"`
	}
)

func (req CreateOfferRequest) Validate() error {
	if req.Price <= 0 {
		return errdomain.NewUserError("Price must be positive number.", "price")
	}
	return nil
}

func (req UpdateOfferRequest) Validate() error {
	if req.Price != nil && *req.Price <= 0 {
		return errdomain.NewUserError("Price must be positive number.", "price")
	}
	return nil
}
