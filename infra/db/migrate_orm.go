package db

import (
	"fmt"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

type MigrateOrm struct {
	Db *gorm.DB
	m  *gormigrate.Gormigrate
}

func NewMigrateOrm(db *gorm.DB) *MigrateOrm {
	m := MigrateOrm{
		Db: db,
	}
	m.load()
	return &m
}

func (m *MigrateOrm) load() {
	m.m = gormigrate.New(m.Db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202203231557",
			Migrate: func(db *gorm.DB) error {
				type Base struct {
					ID        *string    `gorm:"type:uuid;primaryKey"`
					CreatedAt *time.Time `gorm:"column:created_at;autoUpdateTime"`
					UpdatedAt *time.Time `gorm:"column:updated_at;autoCreateTime"`
				}
				type Guest struct {
					Base
					Name *string `gorm:"column:name;not null"`
					Doc  *string `gorm:"column:doc;type:varchar(15);unique"`
				}
				type Tag struct {
					Base
					Name *string `gorm:"column:name;not null;unique"`
				}
				type Score struct {
					Base
					Date        *time.Time `gorm:"column:date;not null"`
					Description *string    `gorm:"column:description;type:varchar(255)"`
					WasUsed     *bool      `gorm:"column:was_used;not null"`
					UsedIn      *time.Time `gorm:"column:used_in"`
					GuestID     *string    `gorm:"column:guest_id;type:uuid;not null"`
				}

				return db.AutoMigrate(&Guest{}, &Tag{}, &Score{})
			},
			Rollback: func(db *gorm.DB) error {
				return db.Migrator().DropTable("guests", "tags", "scores")
			},
		},
	})
}

func (m *MigrateOrm) Migrate() error {
	if err := m.m.Migrate(); err != nil {
		return fmt.Errorf("could not migrate: %v", err)
	}
	return nil
}
