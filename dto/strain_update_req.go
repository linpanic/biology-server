package dto

import "github.com/linpanic/biology-server/etype"

type StrainUpdateReq struct {
	Id             int64       `json:"id,omitempty"`
	StrainName     string      `json:"strain_name,omitempty"`
	Number         string      `json:"number,omitempty"`
	ShortName      []string    `json:"short_name,omitempty"`
	StrainAnnotate []string    `json:"strain_annotate,omitempty"`
	StrainExtra    []ExtraInfo `json:"strain_extra,omitempty"`
	Allele         []Allele    `json:"allele,omitempty"`
}

func (s *StrainUpdateReq) Verify() error {
	if s.Id == 0 || s.Number == "" {
		return etype.FileNullError
	}
	return nil
}
