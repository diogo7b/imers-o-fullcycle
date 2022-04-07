import (
	"github.com/asakevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base `valid:"required"`
	OwnerName `jdon:"owner_name" valid:"notnull"`
	Bank *Bank `valid:"-"`
	Number `jdon:"number" valid:"notnull"`
	PixKeys []*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, number string, ownerName string) (*Account, error){
	account := Account{
		OwnerName: ownerName,
		Number: number,
		Bank: bank
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := isValid()
	if err != nil {
		return nil,err
	}

	return  &account, nil
}