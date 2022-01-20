package service

import (
	"context"
	"github.com/enspzr/getir-go-assigment/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecordService struct {
	db     *mongo.Client
	dbName string
}

func NewRecordService(db *mongo.Client, name string) *RecordService {
	return &RecordService{db: db, dbName: name}
}

// GetAll
// This method creates query by filter.
// Run created query and return result.
// If any error, returns error.
func (r *RecordService) GetAll(filter model.RecordFilter) ([]model.Record, error) {
	// It is used to convert the "Counts" column to individual rows.
	pp := mongo.Pipeline{
		bson.D{{"$unwind", "$counts"}},
	}

	// If start date where inside the filter has value, adds condition in query.
	// Added query: created at bigger than or equal start date.
	if !filter.StartDate.IsZero() {
		pp = append(pp, bson.D{{"$match", bson.D{
			{"createdAt", bson.D{
				{"$gte", primitive.NewDateTimeFromTime(filter.StartDate)},
			}},
		}}})
	}

	// If end date where inside the filter has value, adds condition in query.
	// Added query: created at smaller than or equal end date.
	if !filter.EndDate.IsZero() {
		pp = append(pp, bson.D{{"$match", bson.D{
			{"createdAt", bson.D{
				{"$lte", primitive.NewDateTimeFromTime(filter.EndDate)},
			}},
		}}})
	}

	pp = append(pp, bson.D{{"$group", bson.D{{
		"_id", "$_id",
	}, {
		"key", bson.D{{"$first", "$key"}},
	}, {
		"createdAt", bson.D{{"$first", "$createdAt"}},
	}, {
		"totalCount", bson.D{{"$sum", "$counts"}},
	}}}})

	// If min count where inside the filter has value, adds condition in query.
	// Added query: sum counts bigger than or equal min count.
	if filter.MinCount != 0 {
		pp = append(pp, bson.D{{"$match", bson.D{
			{"totalCount", bson.D{
				{"$gte", filter.MinCount},
			}},
		}}})
	}

	// If max count where inside the filter has value, adds condition in query.
	// Added query: sum counts smaller than or equal max count.
	if filter.MaxCount != 0 {
		pp = append(pp, bson.D{{"$match", bson.D{
			{"totalCount", bson.D{
				{"$lte", filter.MaxCount},
			}},
		}}})
	}
	// Define returned columns.
	pp = append(pp, bson.D{{"$project", bson.D{{"key", 1}, {"createdAt", 1}, {"totalCount", 1}}}})

	ctx := context.Background()
	// Execute query.
	cursor, err := r.db.Database(r.dbName).
		Collection("records").
		Aggregate(ctx, pp)
	if err != nil {
		return nil, err
	}
	records := make([]model.Record, 0)
	defer cursor.Close(ctx)
	// Transfer query results in records array.
	err = cursor.All(ctx, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
