package handlers

import (
	"errors"
	"github.com/enspzr/getir-go-assigment/database"
	"github.com/enspzr/getir-go-assigment/model"
	"github.com/enspzr/getir-go-assigment/service"
	"net/http"
	"strconv"
	"time"
)

var (
	errMethodNotAllowed = errors.New("method not allowed")
)

// RecordGetAll
// This function returns records in database by filters.
// Filter contains startDate, endDate, minCount, maxCount.
// If startDate has value, createdAt returns records greater than startDate.
// If endDate has value, createdAt returns records smaller than endDate.
// Date format is YYYY-MM-DD.
// If minCount has value, sum "counts" returns records greater than minCount.
// If maxCount has value, sum "counts" returns records smaller than maxCount.
// Request url: /records?startDate=2015-10-14&endDate=2016-12-27&minCount=10&maxCount=4000 (GET)
func RecordGetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		successRecordResponse(w, nil)
		return
	}
	if r.Method != "GET" {
		methodNotAllowedError(w, errMethodNotAllowed)
		return
	}
	// Get filter values.
	// If url variables are not in true format, returns error.
	filter, err := getFilter(r)
	if err != nil {
		badRequestError(w, errors.New("Query read error => "+err.Error()))
		return
	}

	rs := service.NewRecordService(database.DB(), database.DbName())
	// Get records in database by filter.
	records, err := rs.GetAll(filter)
	if err != nil {
		internalError(w, err)
		return
	}
	successRecordResponse(w, records)
}

const filterTimeLayout = "2006-01-02"

// This function read query.
// If url variables are not in true format, returns error.
func getFilter(r *http.Request) (model.RecordFilter, error) {
	var recordFilter = model.RecordFilter{}

	// For reading time variables.
	setTime := func(key string) (time.Time, error) {
		query := r.URL.Query().Get(key)
		if query != "" {
			t, err := time.Parse(filterTimeLayout, query)
			if err != nil {
				return t, err
			}
			return t, nil
		}
		return time.Time{}, nil
	}

	// For reading number variables.
	setCount := func(key string) (int, error) {
		query := r.URL.Query().Get(key)
		if query != "" {
			val, err := strconv.Atoi(query)
			if err != nil {
				return 0, err
			}
			return val, nil
		}
		return 0, nil
	}

	var err error
	if recordFilter.StartDate, err = setTime("startDate"); err != nil {
		return recordFilter, err
	}

	if recordFilter.EndDate, err = setTime("endDate"); err != nil {
		return recordFilter, err
	}

	if recordFilter.MinCount, err = setCount("minCount"); err != nil {
		return recordFilter, err
	}

	if recordFilter.MaxCount, err = setCount("maxCount"); err != nil {
		return recordFilter, err
	}

	return recordFilter, nil
}
