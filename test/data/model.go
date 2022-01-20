package data

import "github.com/enspzr/getir-go-assigment/model"

// MockResponse
// This model was created for the test records handler.
type MockResponse struct {
	Code    int            `json:"code"`
	Msg     string         `json:"msg"`
	Records []model.Record `json:"records"`
}
