package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true) //faz as validaçõs do valid, nesse caso, valida que tem q ter id
}

type Base struct {
	ID        string    `json:"id" valid:"uuid"`
	CreatedAt time.Time `json:"created_at" valid:"-"` // esse valid - indica que pode ou não ser preenchido
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}
