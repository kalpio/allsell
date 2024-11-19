package auction

import (
	"github.com/Rhymond/go-money"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	allselltime "github.com/kalpio/allsell/src/types/time"
	"sync"
	"time"
)

type Auction struct {
	ID       uuid.UUID          `json:"id" db:"id"`
	Title    string             `json:"title" db:"title"`
	ExpireAt allselltime.DbTime `json:"expire_at" db:"expire_at"`
	Category string             `json:"category" db:"category"`
	Price    money.Money        `json:"price" db:"price"`
	Images   []Image            `json:"images"`

	mu sync.RWMutex
}

func NewAuction(
	title string,
	expireAt time.Time,
	category string,
	price money.Money) *Auction {
	return &Auction{
		ID:       uuid.New(),
		Title:    title,
		ExpireAt: *allselltime.New(expireAt),
		Category: category,
		Price:    price,
	}
}

func (a *Auction) AddImage(image Image) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Images = append(a.Images, image)
}

func (a *Auction) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required, validation.Length(1, 256)),
		validation.Field(&a.ExpireAt, validation.Required),
		validation.Field(&a.Category, validation.Required),
		validation.Field(&a.Price, validation.Required))
}

type CreateActionRequest struct {
	Title    string      `json:"title"`
	ExpireAt time.Time   `json:"expire_at"`
	Category string      `json:"category"`
	Price    money.Money `json:"price"`
}
