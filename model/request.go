package model

type Request struct {
	ProjectRequest         *Project      `json:"projectRequest"`
	TenancyRequest         *Tenancy      `json:"tenancyRequest"`
	Integrations           *Integrations `json:"integrations"`
	DevelopmentEnvironment *Development  `json:"developmentEnvironment"`
}
