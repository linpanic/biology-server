package dto

type ExtraInfo struct {
	Id       int64  `json:"id,omitempty"`
	ExtraKey string `json:"extra_key,omitempty"`
	ExtraVal string `json:"extra_val,omitempty"`
}
