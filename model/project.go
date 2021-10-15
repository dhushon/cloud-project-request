package model

import (
	"gopkg.in/go-playground/validator.v9"
	"time"
	"encoding/json"
)


//validate.RegisterValidation("projunique", validators.ProjectNameNotUsed)
func ProjectNameNotUsed(fl validator.FieldLevel) bool {
	return false
}
type Person struct {
	Name 	string `json:"name"`
	Email 	string `json:"email" validate:"email" binding:"required"`
	Phone   string `json:"phone"`
}

type SalesId string

type Contract string

type PTime time.Time //Standard time is RFC3339

type OrgGroup string

type Project struct {
	ProjectName     string     `json:"projectName,omitempty" validate:"max=40,min=5,projunique" binding:"required"`
	ProjectId 		string     `json:"projectId,omitempty" validate:"max"`
	ProjectGroup 	*OrgGroup  `json:"projectGroup" validate:"" binding:"required"`
	ProjectOwner    *Person	   `json:"projectOwner" validate:"person" binding:"required"`
	ProjectEngineer *Person	   `json:"projectEngineer" validate:"person" binding:"required"`
	SalesId			SalesId    `json:"salesId,omitempty" validate:"" binding:""`
	Contract        Contract   `json:"contract,omitempty" validate:"" binding:""`
	Start			PTime	   `json:"expectedStart,omitempty" validate:"" binding:""`
	End				PTime	   `json:"expectectedEnd,omitempty" validate:"" binding:""`
	Regulations     []*string  `json:"regulations" validate:"" binding:""`
	Partners 		[]*string  `json:"developmentPartners" validate:"" binding:""`
}


func (p *Project) UnmarshalJSON(data []byte) (err error) {

	type Alias Project
	aux := &struct {
		*Alias
		// put hybrid encoders here
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	//i.IdpCustomer = toIdp(aux.IdpCustomer)
	validate = getValidator()
	//register all unique validators
	validate.RegisterValidation("projunique", ProjectNameNotUsed)
	// run validators
	err = validate.Struct(p)
	return err
}