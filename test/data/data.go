package data

import (
	"embed"
	"encoding/json"
	"github.com/enspzr/getir-go-assigment/model"
)

//go:embed recorddata.json
var file embed.FS

// Reads and returns mock data for test data.
func getMockRecordData() ([]model.MockRecord, error) {
	mockRecords := make([]model.MockRecord, 0)
	jsonFile, err := file.ReadFile("recorddata.json")
	if err != nil {
		return mockRecords, err
	}
	err = json.Unmarshal(jsonFile, &mockRecords)
	return mockRecords, err
}
