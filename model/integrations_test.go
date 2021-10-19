package model

import (
	"reflect"
	"testing"
)

func TestIdentity_UnmarshalJSON(t *testing.T) {
	type fields Identity
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "normal", fields: fields{Okta, "admin@unknown.com","project-dev","project-admin"}, args: args{data: []byte("{\"idProvider\":\"okta\",\"principal\":\"admin@unknown.com\",\"defaultGroup\":\"project-dev\",\"privGroup\":\"project-admin\"}")}, wantErr: false},
		{name: "outofrange", fields: fields{Unknown,"admin@unknown.com","project-dev","project-admin"}, args: args{data: []byte("{\"idProvider\":\"forgerock\",\"principal\":\"admin@unknown.com\",\"defaultGroup\":\"project-dev\",\"privGroup\":\"project-admin\"}")}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Identity{
				IdProvider:  tt.fields.IdProvider,
				Principal: tt.fields.Principal,
				DefaultGroup: tt.fields.DefaultGroup,
				PrivGroup: tt.fields.PrivGroup,
			}
			if err := i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Integrations.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIntegrations_MarshalJSON(t *testing.T) {
	type fields Identity

	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{ 
		{name: "normal", fields: fields{Okta, "admin@unknown.com","project-dev","project-admin"}, want: []byte("{\"idProvider\":\"okta\",\"principal\":\"admin@unknown.com\",\"defaultGroup\":\"project-dev\",\"privGroup\":\"project-admin\"}"), wantErr: false},
		{name: "outofrange", fields: fields{0, "admin@unknown.com","project-dev","project-admin"}, want: []byte("{\"idProvider\":\"unknown\",\"principal\":\"admin@unknown.com\",\"defaultGroup\":\"project-dev\",\"privGroup\":\"project-admin\"}"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Identity{
				IdProvider:  tt.fields.IdProvider,
				Principal: tt.fields.Principal,
				DefaultGroup: tt.fields.DefaultGroup,
				PrivGroup: tt.fields.PrivGroup,
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
