package storage

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SampleStore struct {
	db *MongoDB
}

func NewSampleStore(db *MongoDB) *SampleStore {
	return &SampleStore{db: db}
}

// list with filter
func (s *SampleStore) All() []Sample {
	cursor, err := s.db.GetCollection("Sample").Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Sample
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func (s *SampleStore) Insert(sample *Sample) (string, error) {
	result, err := s.db.GetCollection("Sample").InsertOne(context.TODO(), sample)
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get sample id")
	}
	return id.Hex(), err
}

/*
func (s *SampleStore) FindByID(id string) Sample {
	objID, err := primitive.ObjectIDFromHex(id)

}
func (s *SampleStore) FindByCategory(cid string) Sample {
//https://docs.mongodb.com/manual/tutorial/query-embedded-documents/
}
*/

type UpdateOptions struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	AnyList      []string          `json:"any_list"`
	AnyListItems map[string]string `json:"any_list_items"`
	IntVal       *int              `json:"int_val"`
	FloatVal     *float32          `json:"float_val"`
	Category     string            `json:"category"`
	Title        *string
}

func (s *SampleStore) Update(opts *UpdateOptions) error {
	objID, err := primitive.ObjectIDFromHex(opts.ID)
	if err != nil {
		return err
	}

	newOpts := make(bson.M)
	sampleTags := NewSampleTags()

	if opts.Name != "" {
		newOpts[sampleTags.GetNameTag()] = opts.Name
	}
	if opts.AnyList != nil {
		newOpts[sampleTags.GetAnyListTag()] = opts.AnyList
	} else if opts.AnyListItems != nil {
		for key, val := range opts.AnyListItems {
			newOpts[sampleTags.GetAnyListTag(key)] = val
		}
	}
	if opts.IntVal != nil {
		newOpts["int_val"] = *opts.IntVal
	}
	if opts.FloatVal != nil {
		newOpts["float_val"] = *opts.FloatVal
	}
	if opts.Category != "" {
		newOpts[sampleTags.GetCategoryTag()] = opts.Category
	}
	if opts.Title != nil {
		newOpts[sampleTags.GetTitleTag()] = opts.Title
	}

	result, err := s.db.GetCollection("Sample").UpdateByID(
		context.TODO(),
		objID,
		bson.M{"$set": newOpts},
	)

	fmt.Println(bson.M{"$set": newOpts}, result, err, opts)

	if err == nil && result.MatchedCount == 0 {
		err = mongo.ErrNoDocuments
	}
	if err != nil {
		return fmt.Errorf("failed to update sample: %w", err)
	}

	return nil
}

func (s *SampleStore) SetValidator() {

}
