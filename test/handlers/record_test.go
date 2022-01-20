package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/enspzr/getir-go-assigment/database"
	"github.com/enspzr/getir-go-assigment/handlers"
	"github.com/enspzr/getir-go-assigment/model"
	"github.com/enspzr/getir-go-assigment/test/data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// This test method testing records handlers.
func TestRecordGetAll(t *testing.T) {
	// Init mock mongodb.
	uri, _, ms := data.InitMockDB(t)
	defer ms.Stop()
	// Connect mock mongodb.
	err := database.Connect(uri)
	if err != nil {
		t.Error("Database connection error =>", err.Error())
	}

	// Setup router.
	http.HandleFunc("/records", handlers.RecordGetAll)
	// Start http server.
	go func() {
		err = http.ListenAndServe(":8083", nil)
		if err != nil {
			t.Fatal("Http server starting error => " + err.Error())
		}
	}()

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
	// Test endpoint with test cases.
	for _, tk := range newTest {
		t.Run("Handlers-RecordGetAll:"+tk.input.String(), func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8083/records%s", tk.input.String()))
			if err != nil {
				t.Errorf("Get Request Error => %s", err.Error())
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Reading body error => " + err.Error())
			}

			var result data.MockResponse
			err = json.Unmarshal(body, &result)
			if err != nil {
				t.Errorf("Json unmarshall error => " + err.Error())
			}
			if result.Code != 0 {
				t.Errorf("Result is not success => %d", result.Code)
			}
			if !assert.ElementsMatch(t, result.Records, tk.output) {
				t.Errorf("got %v want %v", result.Records, tk.output)
			}
		})
	}
}
