package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gopodcast/feeds"
	"gopodcast/model"
)

func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.Podcast, error) {
	res, err := r.Api.Search(term)
	if err != nil {
		return nil, err
	}
	var podcasts []*model.Podcast
	for _, res := range res.Results {
		podcast := &model.Podcast{
			Artist:        res.ArtistName,
			Name:          res.TrackName,
			FeedURL:       res.FeedURL,
			Thumbnail:     res.ArtworkURL100,
			EpisodesCount: res.TrackCount,
			Genres:        res.Genres,
		}
		podcasts = append(podcasts, podcast)
	}
	return podcasts, nil
}

func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.Feed, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}
	var feeds []*model.Feed
	for _, item := range res.Channel.Item {
		feed := &model.Feed{
			PublishDate: item.PubDate,
			Text:        item.Text,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       nil,
			Summary:     item.Summary,
			LinkURL:     item.Enclosure.URL,
			Duration:    item.Duration,
		}
		feeds = append(feeds, feed)
	}
	return feeds, nil
}
