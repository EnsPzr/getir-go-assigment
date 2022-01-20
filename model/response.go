package model

// Response
// This struct contains response model fields.
type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Records interface{} `json:"records"`
}
