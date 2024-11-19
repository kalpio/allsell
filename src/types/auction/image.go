package auction

import "github.com/google/uuid"

type Image struct {
	ID        uuid.UUID `json:"id" db:"id"`
	AuctionID uuid.UUID `json:"auction_id" db:"auction_id"`
	Value     []byte    `json:"value" db:"value"`
}
