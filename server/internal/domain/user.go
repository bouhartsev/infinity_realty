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

	UpdateUserRequest struct {
		UserId     int      `json:"-"`
		Name       *string  `json:"name,omitempty"`
		Surname    *string  `json:"surname,omitempty"`
		Patronymic *string  `json:"patronymic,omitempty"`
		Password   *string  `json:"password,omitempty"`
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
		if req.Email != nil && !regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(*req.Email) {
			return errdomain.NewUserError("Email specified in invalid format. Valid format: something@mail.com.", "email")
		}
		if req.Telephone != nil && !regexp.MustCompile("8[0-9]{10}").MatchString(*req.Telephone) {
			return errdomain.NewUserError("Telephone specified in invalid format. Valid format: 8XXXXXXXXX.", "telephone")
		}
	}

	if req.Commission != nil && *req.Commission < 0 {
		return errdomain.NewUserError("Commission must be positive number.", "commission")
	}

	return nil
}

func (req UpdateUserRequest) Validate() error {
	if req.Name != nil && *req.Name == "" {
		return errdomain.NewUserError("Name cannot be empty.", "email")
	}

	if req.Surname != nil && *req.Surname == "" {
		return errdomain.NewUserError("Surname cannot be empty.", "surname")
	}

	if req.Patronymic != nil && *req.Patronymic == "" {
		return errdomain.NewUserError("Patronymic cannot be empty.", "patronymic")
	}

	if req.Password != nil && *req.Password == "" {
		return errdomain.NewUserError("Password cannot be empty.", "password")
	}

	if req.Email != nil && !regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$").MatchString(*req.Email) {
		return errdomain.NewUserError("Email specified in invalid format. Valid format: something@mail.com.", "email")
	}

	if req.Telephone != nil && !regexp.MustCompile("8[0-9]{10}").MatchString(*req.Telephone) {
		return errdomain.NewUserError("Telephone specified in invalid format. Valid format: 8XXXXXXXXX.", "telephone")
	}

	if req.Commission != nil && *req.Commission < 0 {
		return errdomain.NewUserError("Commission must be positive number.", "commission")
	}

	return nil
}
