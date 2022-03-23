package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/patricksferraz/loyalty-card/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Tag struct {
	Base   `json:",inline" valid:"-"`
	Name   *string  `json:"name" gorm:"column:name;not null;unique" valid:"required"`
	Scores []*Score `json:"scores,omitempty" gorm:"many2many:scores_tags" valid:"-"`
}

func NewTag(name *string) (*Tag, error) {
	e := Tag{
		Name: name,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Tag) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

type FilterTag struct {
	Name *string `json:"name" valid:"-"`
}

func (e *FilterTag) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func NewFilterTag(name *string) (*FilterTag, error) {
	e := &FilterTag{
		Name: name,
	}

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return e, nil
}
