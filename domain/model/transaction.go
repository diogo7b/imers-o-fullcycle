package model

import (
	"errors"
	"github.com/asakevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	TransacationPending  string = "pendeing"
	TransacationCompleted string = "completed"
	TransacationError string = "error"
	TransacationConfirmed string = "confirmed"
)

type TransacttionRepositoryInterface interface {
	Register(transacation *Transacation) error
	Save(transacation *Transacation) error
	Find(id string) (*Transacation, error)
}

type Transacation struct {
	Base `valid:"required"`
	AccountFrom *Account `valid:"-" `
	Amount float64 `json:"amount" valid:"notnull"`
	PixKeyTo *PixKey `valid:"-"`
	Status string `json:"status" valid:"notnull"`
	Description string `json:"description" valid:"notnull"`
	CancelDescription string `json:"cancel_description" valid:"notnull"`
}

type Transacations struct{
	Transacation []Transacation
}


func (t *Transacation) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("The amount most be greater than 0")
	}

	if t.Status != TransacationPending && t.Status != TransacationCompleted && t.Status != TransacationError {
		return errors.New("invalid status for the transacation")
	}

	if t.PixKeyTo.AccountID == t.AccountFrom.ID {
		return errors.New("the source and destiantion account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string ) (*Account, error){
	transacation := Transacation{
		AccountFrom: accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransacationPending,
		Description: description,
	}

	transacation.ID = uuid.NewV4().String()
	transacation.CreatedAt = time.Now()

	err := transacation.isValid()
	if err != nil {
		return nil,err
	}

	return  &transacation, nil
}

func (t *Transacation) Complete() error {
	t.Status = TransacationCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid
	return err
}

func (t *Transacation) confirmed() error {
	t.Status = TransacationConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid
	return err
}

func (t *Transacation) Cancel(description string) error {
	t.Status = TransacationError
	t.UpdatedAt = time.Now()
	err := t.isValid
	return err
}