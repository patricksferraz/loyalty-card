package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/loyalty-card/domain/entity"
	"github.com/c-4u/loyalty-card/infra/db"
)

type Repository struct {
	Orm *db.DbOrm
}

func NewRepository(orm *db.DbOrm) *Repository {
	return &Repository{
		Orm: orm,
	}
}

func (r *Repository) CreateGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Orm.Db.Create(guest).Error
	return err
}

func (r *Repository) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	var e entity.Guest

	r.Orm.Db.First(&e, "id = ?", *guestID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest found")
	}

	return &e, nil
}

func (r *Repository) SaveGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Orm.Db.Save(guest).Error
	return err
}

func (r *Repository) CreateScore(ctx context.Context, score *entity.Score) error {
	err := r.Orm.Db.Create(score).Error
	return err
}

func (r *Repository) FindScore(ctx context.Context, scoreID *string) (*entity.Score, error) {
	var e entity.Score

	r.Orm.Db.First(&e, "id = ?", *scoreID)

	if e.ID == nil {
		return nil, fmt.Errorf("no score found")
	}

	return &e, nil
}

func (r *Repository) SaveScore(ctx context.Context, score *entity.Score) error {
	err := r.Orm.Db.Save(score).Error
	return err
}
