package models

import (
	"bytes"
	"gopkg.in/mgo.v2/bson"
	"time"
)


type VideoModel struct {
	ID          bson.ObjectId
	Filename    string
	ContentType string
	Duration    float64
	UploadDate  time.Time
	Usage       string
}


type VideoBuffer struct {
	*bytes.Buffer
	Duration  float64
	ContentType string
}


func NewVideoModel() *VideoModel{
	return &VideoModel{
		ID:bson.NewObjectId(),
		UploadDate:time.Now(),
	}
}

