package data

import (
	"time"

	"wetransfer.OsbornCollins.net/internal/validator"
)

type Info struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Name       string    `json:"name"`
	Occupation string    `json:"occupation,omitempty"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email,omitempty"`
	Country    string    `json:"country"`
	Address    string    `json:"address"`
}

func ValidateInfo(v *validator.Validator, info *Info) {

	// Use the check() method to execute our validation checks
	//Name
	v.Check(info.Name != "", "name", "must be provided")
	v.Check(len(info.Name) <= 200, "name", "must not be move than 200 bytes long")

	//Occupation
	v.Check(info.Occupation != "", "occupation", "must be provided")
	v.Check(len(info.Occupation) <= 200, "occupation", "must not be move than 200 bytes long")

	//Phone
	v.Check(info.Phone != "", "phone", "must be provided")
	v.Check(validator.Matches(info.Phone, validator.PhoneRx), "phone", "must be a valid phone number")

	//Email
	v.Check(info.Email != "", "email", "must be provided")
	v.Check(validator.Matches(info.Email, validator.EmailRx), "email", "must be a valid email address")

	//Country
	v.Check(info.Country != "", "country", "must be provided")
	v.Check(len(info.Country) <= 200, "country", "must not be move than 200 bytes long")

	//Address
	v.Check(info.Address != "", "address", "must be provided")
	v.Check(len(info.Address) <= 500, "address", "must not be move than 500 bytes long")
}
