package repo

import (
	"context"

	"github.com/patricksferraz/loyalty-card/domain/entity"
)

type RepoInterface interface {
	CreateGuest(ctx context.Context, guest *entity.Guest) error
	FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error)
	SaveGuest(ctx context.Context, guest *entity.Guest) error

	CreateScore(ctx context.Context, score *entity.Score) error
	FindScore(ctx context.Context, scoreID *string) (*entity.Score, error)
	SaveScore(ctx context.Context, score *entity.Score) error

	CreateTag(ctx context.Context, tag *entity.Tag) error
	FindTag(ctx context.Context, tagID *string) (*entity.Tag, error)
	SearchTags(ctx context.Context, filterTag *entity.FilterTag) ([]*entity.Tag, error)
	SaveTag(ctx context.Context, tag *entity.Tag) error
}
