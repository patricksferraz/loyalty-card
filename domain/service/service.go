package service

import (
	"context"
	"time"

	"github.com/patricksferraz/loyalty-card/domain/entity"
	"github.com/patricksferraz/loyalty-card/domain/repo"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateGuest(ctx context.Context, name, doc *string) (*string, error) {
	guest, err := entity.NewGuest(name, doc)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateGuest(ctx, guest); err != nil {
		return nil, err
	}

	return guest.ID, nil
}

func (s *Service) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (s *Service) CreateScore(ctx context.Context, date *time.Time, guestID, description *string) (*string, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	score, err := entity.NewScore(date, description, guest)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateScore(ctx, score); err != nil {
		return nil, err
	}

	return score.ID, nil
}

func (s *Service) FindScore(ctx context.Context, scoreID *string) (*entity.Score, error) {
	score, err := s.Repo.FindScore(ctx, scoreID)
	if err != nil {
		return nil, err
	}

	return score, nil
}

func (s *Service) UseScore(ctx context.Context, scoreID *string) error {
	score, err := s.Repo.FindScore(ctx, scoreID)
	if err != nil {
		return err
	}

	err = score.Use()
	if err != nil {
		return err
	}

	err = s.Repo.SaveScore(ctx, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddTag(ctx context.Context, scoreID *string, tagID *string) error {
	tag, err := s.FindTag(ctx, tagID)
	if err != nil {
		return err
	}

	score, err := s.Repo.FindScore(ctx, scoreID)
	if err != nil {
		return err
	}

	err = score.AddTag(tag)
	if err != nil {
		return err
	}

	err = s.Repo.SaveScore(ctx, score)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateTag(ctx context.Context, name *string) (*string, error) {
	tag, err := entity.NewTag(name)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateTag(ctx, tag); err != nil {
		return nil, err
	}

	return tag.ID, nil
}

func (s *Service) FindTag(ctx context.Context, tagID *string) (*entity.Tag, error) {
	tag, err := s.Repo.FindTag(ctx, tagID)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *Service) SearchTags(ctx context.Context, name *string) ([]*entity.Tag, error) {
	f, err := entity.NewFilterTag(name)
	if err != nil {
		return nil, err
	}

	tags, err := s.Repo.SearchTags(ctx, f)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
