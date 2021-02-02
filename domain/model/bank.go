package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Bank struct {
	Base     `valid:"required"` // Como se herdasse tudo do Base
	Code     string             `json:"code" valid:"notnull"` //Essas tags permite que esse model seja convertido em json
	Name     string             `json:"name" valid:"notnull"`
	Accounts []*Account         `valid:"-"`
}

func (bank *Bank) isValid() error { // o primeiro parentende indica que esta atachando isso a classe Bank, logo isso passar a ser um meodo
	_, err := govalidator.ValidateStruct(bank) // valida se os campos estão certos
	if err != nil {
		return err
	}
	return nil
}

func NewBank(code string, name string) (*Bank, error) { // segundo parente indica que retorna ou um Bank ou um erro. O * é pq é um ponteiro
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil // como o retorno é ponteiro(*) tem q usar o & informando que ta enviando um ponteiro
}
