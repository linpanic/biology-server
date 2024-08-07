package dto

type UserLoginResp struct {
	Token    string `json:"token,omitempty"`
	UserName string `json:"user_name,omitempty"`
}
