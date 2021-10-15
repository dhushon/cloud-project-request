package model

import (
	"reflect"
	"testing"
)

func TestIntegrations_UnmarshalJSON(t *testing.T) {
	type fields struct {
		IdpCustomer  Idp
		IdpDeveloper Idp
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "normal", fields: fields{Okta, AzureAd}, args: args{data: []byte("{\"idpCustomer\":\"okta\",\"idpDeveloper\":\"azuread\"}")}, wantErr: false},
		{name: "outofrange", fields: fields{Okta, Unknown}, args: args{data: []byte("{\"idpCustomer\":\"okta\",\"idpDeveloper\":\"forgerock\"}")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integrations{
				IdpCustomer:  tt.fields.IdpCustomer,
				IdpDeveloper: tt.fields.IdpDeveloper,
			}
			if err := i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Integrations.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntegrations_MarshalJSON(t *testing.T) {
	type fields struct {
		IdpCustomer  Idp
		IdpDeveloper Idp
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{ 
		{name: "normal", fields: fields{Okta, AzureAd}, want: []byte("{\"idpCustomer\":\"okta\",\"idpDeveloper\":\"azuread\"}"), wantErr: false},
		{name: "outofrange", fields: fields{Okta, Unknown}, want: []byte("{\"idpCustomer\":\"okta\",\"idpDeveloper\":\"unknown\"}"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Integrations{
				IdpCustomer:  tt.fields.IdpCustomer,
				IdpDeveloper: tt.fields.IdpDeveloper,
			}
			got, err := i.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Integrations.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Integrations.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
