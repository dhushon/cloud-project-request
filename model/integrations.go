package model

import (
	"encoding/json"
	"strings"
	"reflect"
	"github.com/go-playground/validator/v10"
)

type Idp int

const (
	Unknown Idp = iota
	Okta
	AzureAd
	end // keep last as crutch for validator
)

var IdpDict = []string{"unknown", "okta", "azuread"}

func (i *Identity) MarshalJSON() ([]byte, error) {
	type Alias Identity
	return json.Marshal(&struct {
		IdProvider string `json:"idProvider,omitempty"`
		*Alias
	}{
		IdProvider: IdpDict[i.IdProvider],
		Alias:      (*Alias)(i),
	})
}

func toIdp(word string) Idp {
	for k, v := range IdpDict {
		if strings.ToLower(word) == v {
			return Idp(k)
		}
	}
	return Unknown
}

func (i *Identity) UnmarshalJSON(data []byte) (err error) {
	type Alias Identity
	aux := &struct {
		IdProvider string `json:"idProvider"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	i.IdProvider = toIdp(aux.IdProvider)
	validate := getValidator()
	validate.RegisterValidation("idpenum", InIdpEnum)
	err = getValidator().Struct(i)
	return err
}

func InIdpEnum(fl validator.FieldLevel) bool {
	field := fl.Field()
	switch field.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return field.Int() <= int64(end)

	case reflect.Float32, reflect.Float64:
		return field.Float() <= float64(end)
	}
	return false
}

type Identity struct {
	IdProvider   Idp    `json:"idProvider" validate:"idpenum" binding:""`
	Principal    string `json:"principal" validate:"omitempty"`
	DefaultGroup string `json:"defaultGroup" validate:"omitempty"`
	PrivGroup    string `json:"privGroup" validate:"omitempty"`
}

type Integrations struct {
	IdpCustomer  *Identity `json:"idpCustomer,omitempty" validate:"omitempty" binding:""`
	IdpDeveloper *Identity `json:"idpDeveloper" validate:"" binding:"required"`
}
