package model

type Development struct {
	Languages     []*string             `json:"devLanguages,omitempty"`
	Toolchain     string 				`json:"toolchain,omitempty"`
	OpenSourceDependencies []*string    `json:"openSourceDependencies,omitempty"`
}