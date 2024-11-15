package services

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/types/auction"
	"github.com/kalpio/option"
)

type AuctionService struct {
	db *sqlx.DB
}

func NewAuctionService(db *sqlx.DB) AuctionService {
	return AuctionService{db: db}
}

func (srv AuctionService) Create(ctx context.Context, auc auction.Auction) option.Option[auction.Auction] {
	query := `insert into auctions (id, title, expire_at, category, price)
values (?, ?, ?, ?, ?);`

	if _, err := srv.db.ExecContext(
		ctx,
		query,
		auc.ID.String(), auc.Title, auc.ExpireAt.ToDb(), auc.Category, auc.Price); err != nil {
		return option.None[auction.Auction](err)
	}

	return option.Some(auc)
}
