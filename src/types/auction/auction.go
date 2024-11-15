package auction

import (
	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/kalpio/allsell/src/types/time"
)

type Auction struct {
	ID       uuid.UUID   `json:"id" db:"id"`
	Title    string      `json:"title" db:"title"`
	ExpireAt time.DbTime `json:"expire_at" db:"expire_at"`
	Category string      `json:"category" db:"category"`
	Price    money.Money `json:"price" db:"price"`
}

func (a Auction) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required, validation.Length(1, 256)),
		validation.Field(&a.ExpireAt, validation.Required),
		validation.Field(&a.Category, validation.Required),
		validation.Field(&a.Price, validation.Required))
}
