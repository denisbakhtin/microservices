package model

import (
	"videoexample/gen"
)

// MetadataToProto converts a Metadata struct into a
// generated proto counterpart.
func MetadataToProto(m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Author:      m.Author,
	}
}

// MetadataFromProto converts a generated proto counterpart
// into a Metadata struct.
func MetadataFromProto(m *gen.Metadata) *Metadata {
	return &Metadata{
		ID:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Author:      m.Author,
	}
}
