package domain

import "github.com/bouhartsev/infinity_realty/internal/domain/errdomain"

type (
	Realtor struct {
		Id         int      `json:"id"`                   // ID риэлтора.
		Name       string   `json:"name"`                 // Имя риэлтора.
		Surname    string   `json:"surname"`              // Фамилия риэлтора.
		Patronymic string   `json:"patronymic"`           // Отчество риэлтора.
		Commission *float32 `json:"commission,omitempty"` // Комиссия риэлтора.
	}

	CreateRealtorRequest struct {
		Name       string   `json:"name"`                 // Имя риэлтора.
		Surname    string   `json:"surname"`              // Фамилия риэлтора.
		Patronymic string   `json:"patronymic"`           // Отчество риэлтора.
		Commission *float32 `json:"commission,omitempty"` // Комиссия риэлтора.
	}

	UpdateRealtorRequest struct {
		RealtorId  int      `json:"-"`
		Name       *string  `json:"name"`                 // Имя риэлтора.
		Surname    *string  `json:"surname"`              // Фамилия риэлтора.
		Patronymic *string  `json:"patronymic"`           // Отчество риэлтора.
		Commission *float32 `json:"commission,omitempty"` // Комиссия риэлтора.
	}

	GetRealtorResponse struct {
		Realtor Realtor `json:"realtor"`
	}
)

func (req CreateRealtorRequest) Validate() error {
	if req.Name == "" {
		return errdomain.NewUserError("Name cannot be empty.", "email")
	}

	if req.Surname == "" {
		return errdomain.NewUserError("Surname cannot be empty.", "surname")
	}

	if req.Patronymic == "" {
		return errdomain.NewUserError("Patronymic cannot be empty.", "patronymic")
	}

	if req.Commission != nil && *req.Commission < 0 {
		return errdomain.NewUserError("Commission cannot be negative.", "patronymic")
	}

	return nil
}

func (req UpdateRealtorRequest) Validate() error {
	if req.Name != nil && *req.Name == "" {
		return errdomain.NewUserError("Name cannot be empty.", "email")
	}

	if req.Surname != nil && *req.Surname == "" {
		return errdomain.NewUserError("Surname cannot be empty.", "surname")
	}

	if req.Patronymic != nil && *req.Patronymic == "" {
		return errdomain.NewUserError("Patronymic cannot be empty.", "patronymic")
	}

	if req.Commission != nil && *req.Commission < 0 {
		return errdomain.NewUserError("Commission cannot be negative.", "patronymic")
	}

	return nil
}
