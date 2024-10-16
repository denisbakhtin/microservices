package grpc

import (
	"context"
	"errors"
	"videoexample/gen"
	model "videoexample/metadata/pkg"
	"videoexample/video/internal/controller"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a movie gRPC handler.
type Handler struct {
	gen.UnimplementedVideoServiceServer
	ctrl *controller.Controller
}

// New creates a new movie gRPC handler.
func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMovieDetails returns moviie details by id.
func (h *Handler) GetMovieDetails(ctx context.Context, req *gen.GetVideoDetailsRequest) (*gen.GetVideoDetailsResponse, error) {
	if req == nil || req.VideoId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.VideoId)
	if err != nil && errors.Is(err, controller.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetVideoDetailsResponse{
		VideoDetails: &gen.VideoDetails{
			Metadata: model.MetadataToProto(&m.Metadata),
			Rating:   float32(*m.Rating),
		},
	}, nil
}
