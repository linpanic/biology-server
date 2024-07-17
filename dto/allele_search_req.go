package dto

import "errors"

type AlleleSearchReq struct {
	Name string `json:"name,omitempty"`
}

func (a *AlleleSearchReq) Verify() error {
	if a.Name == "" {
		return errors.New("name is null")
	}
	return nil
}
