package domain

import (
	"regexp"

	"github.com/bouhartsev/infinity_realty/internal/domain/errdomain"
)

type (
	User struct {
		Id         int      `json:"id"`
		Role       int      `json:"role"`
		Name       string   `json:"name"`
		Surname    string   `json:"surname"`
		Patronymic string   `json:"patronymic"`
		Email      *string  `json:"email,omitempty"`
		Telephone  *string  `json:"telephone,omitempty"`
		Commission *float32 `json:"commission,omitempty"`
		Password   string   `json:"-"`
	}

	CreateUserRequest struct {
		Name       string   `json:"name"`
		Surname    string   `json:"surname"`
		Patronymic string   `json:"patronymic"`
		Password   string   `json:"password"`
		Role       int      `json:"role"`
		Email      *string  `json:"email,omitempty"`
		Telephone  *string  `json:"telephone,omitempty"`
		Commission *float32 `json:"commission,omitempty"`
	}

	GetUserResponse struct {
		User *User `json:"user"`
	}
)

func (req CreateUserRequest) Validate() error {
	if req.Name == "" {
		return errdomain.NewUserError("Name cannot be empty.", "email")
	}

	if req.Surname == "" {
		return errdomain.NewUserError("Surname cannot be empty.", "surname")
	}

	if req.Patronymic == "" {
		return errdomain.NewUserError("Patronymic cannot be empty.", "patronymic")
	}

	if req.Password == "" {
		return errdomain.NewUserError("Password cannot be empty.", "password")
	}

	if req.Role < 1 || req.Role > 3 {
		return errdomain.NewUserError("Unknown role specified.", "role")
	}

	if req.Email == nil && req.Telephone == nil {
		return errdomain.NewInvalidRequestError("Either email or telephone must be specified.")
	} else {
		if req.Email != nil && !regexp.MustCompile("").MatchString(*req.Email) {
			return errdomain.NewUserError("Email specified in invalid format.", "email")
		}
		if req.Telephone != nil && !regexp.MustCompile("").MatchString(*req.Telephone) {
			return errdomain.NewUserError("Telephone specified in invalid format.", "telephone")
		}
	}

	if req.Commission != nil && *req.Commission < 0 {
		return errdomain.NewUserError("Commission must be positive number.", "commission")
	}

	return nil
}
