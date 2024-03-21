package domain

import "github.com/google/uuid"

type NewUserRequestPayload struct {
	Email       string `json:"email" binding:"required"`
	AddressType string `json:"address_type"`
	Password    string `json:"password"`
	FirstName   string `json:"firstname" binding:"required"`
	LastName    string `json:"lastname" binding:"required"`
	ZipCode     string `json:"zip_code" binding:"required"`
	Type        string `json:"type"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country" binding:"required"`
	DOB         string `json:"dob"`
	Phone       string `json:"phone" binding:"required"`
	CountryCode string `json:"country_code" binding:"required"`
	Username    string `json:"username"`
}

type RegisterUserResponse struct {
	UserID uuid.UUID `json:"id"`
}

// func (r *NewUserRequest) Validate(v *validator.Validator) bool {

// 	r.Email = strings.ToLower(r.Email)

// 	num, err := common.FormatPhone(r.Phone, r.CountryCode)
// 	if err != nil {
// 		v.AddError("phone", "must be a valid phone number")
// 		return v.Valid()
// 	}

// 	r.Phone = num

// 	if user := v.HasIncompleteRegistrationViaEmail(r.Email); user != nil {
// 		r.ExistingUser = user
// 		v.AddError("email", "INCOMPLETE_REGISTRATION_VIA_EMAIL")
// 		return v.Valid()
// 	}

// 	if user := v.HasIncompleteRegistrationViaPhone(r.Phone); user != nil {
// 		r.ExistingUser = user
// 		v.AddError("phone", "INCOMPLETE_REGISTRATION_VIA_PHONE")
// 		return v.Valid()
// 	}

// 	v.Check(r.Email != "", "email", "email must not be blank")
// 	v.Check(validator.IsEmail(r.Email), "email", "must be a valid email address")
// 	v.Check(len(r.Email) <= 200, "email", "must not be more than 200 bytes long")
// 	v.Check(r.AccountType != "", "account_type", "must not be blank")

// 	v.EmailExists(r.Email)

// 	v.Check(validator.In(r.AccountType, ValidAccountTypes...), "account_type", fmt.Sprintf("account type must be one of %s", strings.Join(ValidAccountTypes, ", ")))

// 	if r.AccountType == AccountTypeBusiness {
// 		v.Check(r.BusinessName != "", "business_name", "must not be blank")
// 	}

// 	v.Check(r.Password != "", "password", "must not be blank")
// 	v.Check(len(r.Password) >= MinPasswordLength, "password", "must be at least 8 characters long")
// 	v.Check(len(r.Password) <= MaxPasswordLength, "password", "the password is too long")

// 	// Check the first name and last name.
// 	v.Check(r.FirstName != "", "first_name", "must not be blank")
// 	v.Check(len(r.FirstName) <= 255, "first_name", "must not be more than 50 bytes long")

// 	v.Check(r.LastName != "", "last_name", "must not be blank")
// 	v.Check(len(r.LastName) <= 255, "last_name", "must not be more than 50 bytes long")

// 	v.Check(len(r.MiddleName) <= 255, "last_name", "must not be more than 50 bytes long")

// 	v.Check(len(r.CountryCode) == 2, "country_code", "must be a valid country code")
// 	v.Check(r.Phone != "", "phone", "must not be blank")

// 	v.PhoneExists(r.Phone)

// 	return v.Valid()

// }
