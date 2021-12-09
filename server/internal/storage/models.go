package storage

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Sample struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	AnyList     []string           `bson:"any_list"`
	IntVal      int                `bson:"int_val"`
	FloatVal    float32            `bson:"float_val"`
	InnerDoc    Describe           `bson:"inner_doc"`
	TimeVal     time.Time          `bson:"time"`
	reflectType reflect.Type
}

func NewSampleTags() *Sample {
	s := Sample{}
	s.reflectType = reflect.TypeOf(s)
	s.InnerDoc.reflectType = reflect.TypeOf(s.InnerDoc)
	return &s
}

func (s *Sample) getTag(fieldName string) string {
	field, _ := s.reflectType.FieldByName(fieldName)
	return field.Tag.Get("bson")
}

func (s *Sample) GetNameTag() string {
	return s.getTag("Name")
}

func (s *Sample) GetAnyListTag(index ...string) string {
	listTag := s.getTag("AnyList")
	if index == nil {
		return listTag
	}
	return listTag + "." + index[0]
}

func (s *Sample) GetCategoryTag() string {
	innerDocTag := s.getTag("InnerDoc")
	field, _ := s.InnerDoc.reflectType.FieldByName("Category")
	return innerDocTag + "." + field.Tag.Get("bson")
}

func (s *Sample) GetTitleTag() string {
	innerDocTag := s.getTag("InnerDoc")
	field, _ := s.InnerDoc.reflectType.FieldByName("Title")
	return innerDocTag + "." + field.Tag.Get("bson")
}

type Describe struct {
	Category    string `bson:"category"`
	Title       string `bson:"title"`
	reflectType reflect.Type
}

type Category struct {
	Id   string
	Name string
}

// update category - consistency?
