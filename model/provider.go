package model



var ProviderDict = []string{"unknown", "azure", "aws", "azuregov", "awsgov"}

type Region struct {
	DisplayName string   `json:"name"`
	FullName 	string   `json:"fullName"`
	Latitude 	string 	 `json:"latitude,omitempty"`
	Longitude 	string 	 `json:"longitude,omitempty"`
	Name 		string 	 `json:"code"`
	Public 		bool     `json:"public"`
	Zones 		[]string `json:"zones"`
}

type Provider struct {
	CloudProvider   string     `json:"provider,omitempty" validate:"" binding:"required"`
	PrimaryRegion	*Region     `json:"primaryRegion" validate:"" binding:""`
	SecondaryRegion *Region	   `json:"secondaryRegion,omitempty" validate:"" binding:""`
}

type Tenancy struct {
	Provider   		*Provider  	`json:"provider" validate:"required"`
	Subscription 	string     	`json:"subscription" validate:"" binding:""`
	Blueprint    	string	   	`json:"blueprint" validate:"" binding:""`
	SDLCStages 		[]string 	`json:"sdlcStages" validate:"" binding:""`
	ExternalNetwork string		`json:"network" validate:"" binding:""`
	SharedServices  []string 	`json:"sharedServices" validate:"" binding:""`
}
