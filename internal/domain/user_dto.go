package domain

import (
	"goApiStartetProject/internal/util/validator"
	"strings"

	"github.com/google/uuid"
)

type NewUserRequestPayload struct {
	Email       string `json:"email" binding:"required"`
	AddressType string `json:"address_type"`
	Password    string `json:"password"`
	FirstName   string `json:"firstname" binding:"required"`
	MiddleName   string `json:"middlename"`
	LastName    string `json:"lastname" binding:"required"`
	Username    string `json:"username"`
	ZipCode     string `json:"zip_code"`
	Type        string `json:"type"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country" binding:"required"`
	DOB         string `json:"dob"`
	Phone       string `json:"phone" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
}

type RegisterUserResponse struct {
	UserID uuid.UUID `json:"id"`
}

func (r *NewUserRequestPayload) Validate(v *validator.Validator) bool {

	v.Check(r.Email != "", "email", "email must not be blank")
	v.Check(len(r.Email) <= 200, "email", "must not be more than 200 bytes long")
	r.Email = strings.ToLower(r.Email)


	v.Check(r.Password != "", "password", "must not be blank")
	// v.Check(len(r.Password) >= MinPasswordLength, "password", "must be at least 8 characters long")
	// v.Check(len(r.Password) <= MaxPasswordLength, "password", "the password is too long")

	// Check the first name and last name.
	v.Check(r.FirstName != "", "first_name", "must not be blank")
	v.Check(len(r.FirstName) <= 255, "first_name", "must not be more than 50 bytes long")

	v.Check(r.LastName != "", "last_name", "must not be blank")
	v.Check(len(r.LastName) <= 255, "last_name", "must not be more than 50 bytes long")

	v.Check(len(r.MiddleName) <= 255, "middle_name", "must not be more than 50 bytes long")

	v.Check(len(r.CountryCode) == 2, "country_code", "must be a valid country code")
	v.Check(r.Phone != "", "phone", "must not be blank")

	// v.PhoneExists(r.Phone)

	return v.Valid()

}
