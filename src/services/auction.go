package services

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/types/auction"
	"github.com/kalpio/option"
)

const pageSize = 10

type AuctionService struct {
	db *sqlx.DB
}

func NewAuctionService(db *sqlx.DB) AuctionService {
	return AuctionService{db: db}
}

func (srv AuctionService) List(ctx context.Context, pageIndex int) option.Option[[]*auction.Auction] {
	query := `select * from auctions limit :number offset :starting_row;`
	rows, err := srv.db.NamedQueryContext(ctx, query, map[string]interface{}{"number": pageSize, "starting_row": pageSize * pageIndex})
	defer func() {
		if err := rows.Close(); err != nil {
			// log
		}
	}()

	if err != nil {
		return option.None[[]*auction.Auction](err)
	}

	var result []*auction.Auction
	for rows.Next() {
		auc := &auction.Auction{}
		if err = rows.StructScan(&auc); err != nil {
			return option.None[[]*auction.Auction](err)
		}
		result = append(result, auc)
	}

	return option.Some(result)
}

func (srv AuctionService) Create(ctx context.Context, request auction.CreateActionRequest) option.Option[*auction.Auction] {
	query := `insert into auctions (id, title, expire_at, category, price)
values (:id, :title, :expire_at, :category, :price);`

	auc := auction.NewAuction(request.Title, request.ExpireAt, request.Category, request.Price)
	if _, err := srv.db.ExecContext(ctx, query, auc); err != nil {
		return option.None[*auction.Auction](err)
	}

	return option.Some(auc)
}
