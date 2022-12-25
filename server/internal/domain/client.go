package domain

import (
	"regexp"

	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

type (
	Client struct {
		Id         int     `json:"id"`                  // ID клиента.
		Name       string  `json:"name"`                // Имя клиента.
		Surname    string  `json:"surname"`             // Фамилия клиента.
		Patronymic string  `json:"patronymic"`          // Отчество клиента.
		Email      *string `json:"email,omitempty"`     // Email клиента.
		Telephone  *string `json:"telephone,omitempty"` // Телефон клиента.
	}

	CreateClientRequest struct {
		Name       string  `json:"name" validate:"required"`       // Имя клиента.
		Surname    string  `json:"surname" validate:"required"`    // Фамилия клиента.
		Patronymic string  `json:"patronymic" validate:"required"` // Отчество клиента.
		Email      *string `json:"email,omitempty"`                // Email клиента.
		Telephone  *string `json:"telephone,omitempty"`            // Телефон клиента.
	}

	UpdateClientRequest struct {
		ClientId   int     `json:"-"`
		Name       *string `json:"name,omitempty"`       // Имя клиента.
		Surname    *string `json:"surname,omitempty"`    // Фамилия клиента.
		Patronymic *string `json:"patronymic,omitempty"` // Отчество клиента.
		Email      *string `json:"email,omitempty"`      // Email клиента.
		Telephone  *string `json:"telephone,omitempty"`  // Телефон клиента.
	}

	GetClientResponse struct {
		Client *Client `json:"client"`
	}
)

func (req CreateClientRequest) Validate() error {
	if req.Name == "" {
		return errdomain.NewUserError("Name cannot be empty.", "email")
	}

	if req.Surname == "" {
		return errdomain.NewUserError("Surname cannot be empty.", "surname")
	}

	if req.Patronymic == "" {
		return errdomain.NewUserError("Patronymic cannot be empty.", "patronymic")
	}

	if req.Email == nil && req.Telephone == nil {
		return errdomain.NewInvalidRequestError("Either email or telephone must be specified.")
	} else {
		if req.Email != nil && !regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(*req.Email) {
			return errdomain.NewUserError("Email specified in invalid format. Valid format: something@mail.com.", "email")
		}
		if req.Telephone != nil && !regexp.MustCompile("8[0-9]{10}").MatchString(*req.Telephone) {
			return errdomain.NewUserError("Telephone specified in invalid format. Valid format: 8XXXXXXXXX.", "telephone")
		}
	}

	return nil
}

func (req UpdateClientRequest) Validate() error {
	if req.Name != nil && *req.Name == "" {
		return errdomain.NewUserError("Name cannot be empty.", "email")
	}

	if req.Surname != nil && *req.Surname == "" {
		return errdomain.NewUserError("Surname cannot be empty.", "surname")
	}

	if req.Patronymic != nil && *req.Patronymic == "" {
		return errdomain.NewUserError("Patronymic cannot be empty.", "patronymic")
	}

	if req.Email != nil && !regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(*req.Email) {
		return errdomain.NewUserError("Email specified in invalid format. Valid format: something@mail.com.", "email")
	}

	if req.Telephone != nil && !regexp.MustCompile("8[0-9]{10}").MatchString(*req.Telephone) {
		return errdomain.NewUserError("Telephone specified in invalid format. Valid format: 8XXXXXXXXX.", "telephone")
	}

	return nil
}
