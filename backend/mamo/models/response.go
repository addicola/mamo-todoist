package models

type NetworkHeadResponse struct {
	ID     uint `json:"id"`
	Result struct {
		Header NetworkHeader `json:"header"`
	} `json:"result"`
}
