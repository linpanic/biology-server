package dto

type AlleleAddReq struct {
	Name     string      `json:"name,omitempty"`
	Annotate []string    `json:"annotate,omitempty"`
	Genome   string      `json:"genome,omitempty"`
	Serial   string      `json:"serial,omitempty"`
	Extra    []ExtraInfo `json:"extra,omitempty"`
}
