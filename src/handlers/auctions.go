package handlers

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/types/auction"
	views "github.com/kalpio/allsell/src/views/auctions"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type AuctionsHandler struct {
	db  *sqlx.DB
	srv services.AuctionService
}

func NewAuctionsHandler(db *sqlx.DB) AuctionsHandler {
	return AuctionsHandler{db, services.NewAuctionService(db)}
}

func (h AuctionsHandler) Index(c echo.Context) error {
	return render(c, views.Index())
}

func (h AuctionsHandler) ListGet(c echo.Context) error {
	page, err := strconv.Atoi(param(c, "page", "0"))
	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}
	slog.Info("Render auctions index", "page", page)
	var values []*auction.Auction
	for i := 0; i < 10; i++ {
		values = append(values, auction.NewAuction(fmt.Sprintf("Title of auction no %d", i*page), time.Now().Add(time.Duration(i*40)*time.Hour), fmt.Sprintf("Category #%d", i), *money.New(int64(i*2345), money.PLN)))
	}
	return render(c, views.List(values, page+1))
}

func (h AuctionsHandler) CreateGet(c echo.Context) error {
	return render(c, views.Create())
}

func (h AuctionsHandler) CreatePost(c echo.Context) error {
	request := auction.CreateActionRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	res := h.srv.Create(c.Request().Context(), request)
	if isNone, err := res.IsNone(); isNone {
		return err
	}

	return c.JSON(http.StatusOK, res.Unwrap())
}
