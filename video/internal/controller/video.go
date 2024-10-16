package controller

import (
	"context"
	"errors"
	metadatamodel "videoexample/metadata/pkg"
	ratingmodel "videoexample/rating/pkg"
	"videoexample/video/internal/gateway"
	model "videoexample/video/pkg"
)

// ErrNotFound is returned when the video metadata is not
// found.
var ErrNotFound = errors.New("video metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingmodel.RecordID, recordType ratingmodel.RecordType) (float64, error)
}

type metadataGateway interface {
	Get(ctx context.Context, id string) (*metadatamodel.Metadata, error)
}

// Controller defines a video service controller.
type Controller struct {
	ratingGateway   ratingGateway
	metadataGateway metadataGateway
}

// New creates a new video service controller.
func New(ratingGateway ratingGateway, metadataGateway metadataGateway) *Controller {
	return &Controller{ratingGateway, metadataGateway}
}

// Get returns the video details including the aggregated
// rating and video metadata.
// Get returns the video details including the aggregated rating and video metadata.
func (c *Controller) Get(ctx context.Context, id string) (*model.VideoDetails, error) {
	metadata, err := c.metadataGateway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}
	details := &model.VideoDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingmodel.RecordID(id), ratingmodel.RecordTypeVideo)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// no rating yet
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}
	return details, nil
}
