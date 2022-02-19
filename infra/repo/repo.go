package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/loyalty-card/domain/entity"
	"github.com/c-4u/loyalty-card/infra/db"
)

type Repository struct {
	Pg *db.PostgreSQL
}

func NewRepository(pg *db.PostgreSQL) *Repository {
	return &Repository{
		Pg: pg,
	}
}

func (r *Repository) CreateGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Pg.Db.Create(guest).Error
	return err
}

func (r *Repository) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	var e entity.Guest

	r.Pg.Db.First(&e, "id = ?", *guestID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest found")
	}

	return &e, nil
}

func (r *Repository) SaveGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Pg.Db.Save(guest).Error
	return err
}
