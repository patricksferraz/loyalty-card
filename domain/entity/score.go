package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/c-4u/loyalty-card/utils"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Score struct {
	Base        `json:",inline" valid:"-"`
	Date        *time.Time `json:"date" gorm:"column:date;not null" valid:"required"`
	Description *string    `json:"description,omitempty" gorm:"column:description;type:varchar(255)" valid:"-"`
	WasUsed     *bool      `json:"was_used" gorm:"column:was_used;not null" valid:"-"`
	UsedIn      *time.Time `json:"used_in,omitempty" gorm:"column:used_in" valid:"-"`
	Tags        []*Tag     `json:"tags,omitempty" gorm:"many2many:scores_tags" valid:"-"`
	GuestID     *string    `json:"guest_id" gorm:"column:guest_id;type:uuid;not null" valid:"uuid"`
	Guest       *Guest     `json:"-" valid:"-"`
}

func NewScore(date *time.Time, description *string, guest *Guest) (*Score, error) {
	e := Score{
		Date:        date,
		Description: description,
		WasUsed:     utils.PBool(false),
		GuestID:     guest.ID,
		Guest:       guest,
	}
	e.ID = utils.PString(uuid.NewV4().String())
	e.CreatedAt = utils.PTime(time.Now())

	err := e.IsValid()
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e *Score) IsValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}

func (e *Score) Use() error {
	e.WasUsed = utils.PBool(true)
	e.UsedIn = utils.PTime(time.Now())
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}

func (e *Score) AddTag(tag *Tag) error {
	e.Tags = append(e.Tags, tag)
	e.UpdatedAt = utils.PTime(time.Now())
	err := e.IsValid()
	return err
}
