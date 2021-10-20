package model

type Financial struct {
	BillingCode      string  `json:"billingCode"`
	AuthorizedSpend  int     `json:"authorizedSpend"`
	ApprovingPartner *Person `json:"approvingPartner"`
}
