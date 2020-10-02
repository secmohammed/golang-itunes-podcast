package graphql

import (
	"gopodcast/graphql/generated"
	"gopodcast/itunes"
)

type Resolver struct {
	Api *itunes.API
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
