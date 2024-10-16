package model

import model "videoexample/metadata/pkg"

// VideoDetails includes video metadata its aggregated
// rating.
type VideoDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
