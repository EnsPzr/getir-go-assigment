package service

import (
	"github.com/enspzr/getir-go-assigment/database"
	"github.com/enspzr/getir-go-assigment/model"
	"github.com/enspzr/getir-go-assigment/service"
	"github.com/enspzr/getir-go-assigment/test/data"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// This test function tests the getall method on the "record" service.
func TestRecordGetAll(t *testing.T) {
	// Init mock mongodb.
	uri, dbName, ms := data.InitMockDB(t)
	defer ms.Stop()
	// Connect mock mongodb.
	err := database.Connect(uri)
	if err != nil {
		t.Error("Database connection error =>", err.Error())
	}

	// Test inputs and outputs.
	date2015 := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	date2016 := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	date2017 := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
	var newTest = []struct {
		input  model.RecordFilter
		output []model.Record
	}{
		{model.RecordFilter{
			StartDate: date2015,
		}, data.OutputStartDate2015},
		{model.RecordFilter{
			StartDate: date2016,
		}, data.OutputStartDate2016},
		{model.RecordFilter{
			EndDate: date2016,
		}, data.OutputEndDate2016},
		{model.RecordFilter{
			EndDate: date2017,
		}, data.OutputEndDate2017},
		{model.RecordFilter{
			MinCount: 100,
		}, data.OutputMinCount100},
		{model.RecordFilter{
			MinCount: 3000,
		}, data.OutputMinCount3000},
		{model.RecordFilter{
			MaxCount: 2000,
		}, data.OutputMaxCount2000},
		{model.RecordFilter{
			MaxCount: 4000,
		}, data.OutputMaxCount4000},
	}
	// Test method with test cases.
	rs := service.NewRecordService(database.DB(), dbName)
	for _, tk := range newTest {
		t.Run("Service-GetAllRecords:"+tk.input.String(), func(t *testing.T) {
			records, err := rs.GetAll(tk.input)
			if err != nil {
				t.Errorf("Get all records error => %s, filter => %s", err.Error(), tk.input.String())
			}
			if !assert.ElementsMatch(t, records, tk.output) {
				t.Errorf("got %v want %v", records, tk.output)
			}
		})
	}
}
