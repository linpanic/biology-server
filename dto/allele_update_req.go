package dto

type AlleleUpdateReq struct {
	Id       int64       `json:"id,omitempty"`
	Name     string      `json:"name,omitempty"`
	Annotate []string    `json:"annotate,omitempty"`
	Genome   string      `json:"genome,omitempty"`
	Serial   string      `json:"serial,omitempty"`
	Extra    []ExtraInfo `json:"extra,omitempty"`
}
