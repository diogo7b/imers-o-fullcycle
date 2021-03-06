package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asakevich/govalidator"
)

type Bank struct {
	Base `valid:"required"`
	Code string `json."code" valid:"notnull"`
	Name string `json."name" valid:"notnull"`
	Accounts []*Account `valid:"-"`
}

func (bank *Bank) isValid() error{
	validator, err := govalidator.ValidateStruct(bank)
	return nil
}

func NewBank(code string, name string) (*Bank, error){
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := isValid()
	if err != nil {
		return nil,err
	}

	return  &bank, nil
}