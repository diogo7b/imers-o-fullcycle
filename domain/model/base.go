package model

import (
	"time"
	"github.com/asakevich/govalidator"
)

func init() {
	govalidator.setFieldsRequiredByDefault(value:true)
}
type Base struct {
	ID string `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}