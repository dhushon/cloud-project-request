package model

import (
	"encoding/json"
	"strings"
	//"validator"
)

type Idp int

const (
	Unknown Idp = iota
    Okta
    AzureAd 
)

var  IdpDict = []string {"unknown","okta", "azuread"}

func (i *Integrations) MarshalJSON() ([]byte, error) {
	type Alias Integrations
	return json.Marshal(&struct {
		IdpCustomer  string `json:"idpCustomer,omitempty"`
		IdpDeveloper string `json:"idpDeveloper,omitempty"`
		*Alias
	}{
		IdpCustomer:  IdpDict[i.IdpCustomer],
		IdpDeveloper: IdpDict[i.IdpDeveloper],
		Alias:        (*Alias)(i),
	})
}
func toIdp(word string) (Idp) {
	for k, v := range IdpDict {
        if strings.ToLower(word) == v {
            return Idp(k)
        }
    }
    return Unknown
}

func (i *Integrations) UnmarshalJSON(data []byte) (err error) {
	type Alias Integrations
	aux := &struct {
		IdpCustomer string `json:"idpCustomer,omitempty"`
		IdpDeveloper string `json:"idpDeveloper" binding:"required"`
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	i.IdpCustomer = toIdp(aux.IdpCustomer)
	i.IdpDeveloper = toIdp(aux.IdpDeveloper)
	err = getValidator().Struct(i)
	return err
}

type Integrations struct {
	IdpCustomer     Idp             `json:"idpCustomer,omitempty" validate:"omitempty, oneof= okta azuread unknown"`
	IdpDeveloper	Idp			    `json:"idpDeveloper" validate:"oneof= okta azuread unknown" binding:"required"`
}