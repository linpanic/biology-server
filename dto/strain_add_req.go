package dto

import "github.com/linpanic/biology-server/etype"

type StrainAddReq struct {
	StrainName     string      `json:"strain_name,omitempty"`
	Number         string      `json:"number,omitempty"`
	ShortName      []string    `json:"short_name,omitempty"`
	StrainAnnotate []string    `json:"strain_annotate,omitempty"`
	StrainExtra    []ExtraInfo `json:"strain_extra,omitempty"`
	Allele         []Allele    `json:"allele,omitempty"`
}

func (s *StrainAddReq) Verify() error {
	if s.Number == "" {
		return etype.NumberEmptyError
	}
	return nil
}
