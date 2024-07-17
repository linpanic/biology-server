package dto

import "github.com/linpanic/biology-server/etype"

type StrainDelReq struct {
	StrainId int64 `json:"strain_id,omitempty"`
}

func (s *StrainDelReq) Verify() error {
	if s.StrainId == 0 {
		return etype.FileNullError
	}
	return nil
}
