package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"time"
)

// RecordFilter
// This structure contain records where  request query filters fields.
type RecordFilter struct {
	StartDate time.Time `query:"startDate"`
	EndDate   time.Time `query:"endDate"`
	MinCount  int       `query:"minCount"`
	MaxCount  int       `query:"maxCount"`
}

// This method translates record filters model to query string.
func (r *RecordFilter) String() string {
	query := ""
	if !r.StartDate.IsZero() {
		query += "startDate=" + r.StartDate.Format("2006-01-02") + "&"
	}
	if !r.EndDate.IsZero() {
		query += "endDate=" + r.EndDate.Format("2006-01-02") + "&"
	}
	if r.MinCount != 0 {
		query += "minCount=" + strconv.Itoa(r.MinCount) + "&"
	}
	if r.MaxCount != 0 {
		query += "maxCount=" + strconv.Itoa(r.MaxCount) + "&"
	}
	if query != "" {
		query = "?" + query
	}
	return query
}

// Record
// This structure contains database query result fields.
type Record struct {
	CreatedAt  time.Time `bson:"createdAt" json:"createdAt"`
	Key        string    `bson:"key" json:"key"`
	TotalCount int       `bson:"totalCount" json:"totalCount"`
}

// MockRecord
// This structure uses mock database for test.
type MockRecord struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	Key       string             `bson:"key"`
	Value     string             `bson:"value"`
	Counts    []int              `bson:"counts"`
}
