package service

import (
	"context"

	"github.com/c-4u/loyalty-card/domain/entity"
	"github.com/c-4u/loyalty-card/domain/repo"
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

func (s *Service) CreateScore(ctx context.Context, guestID, description *string, tags *[]string) (*string, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	score, err := entity.NewScore(description, tags, guest)
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
