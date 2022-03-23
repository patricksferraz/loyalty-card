package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbOrm struct {
	Db *gorm.DB
	M  *MigrateOrm
}

func NewDbOrm(dsn string, l logger.LogLevel) (*DbOrm, error) {
	pg := &DbOrm{}

	err := pg.connect(dsn, l)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (p *DbOrm) connect(dsn string, l logger.LogLevel) error {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(l),
	})
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	p.Db = db
	p.M = NewMigrateOrm(p.Db)

	return nil
}

func (p *DbOrm) Migrate() error {
	err := p.M.Migrate()
	return err
}
