package handlers

import (
	"github.com/Rhymond/go-money"
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/types/auction"
	views "github.com/kalpio/allsell/src/views/auctions"
	"github.com/labstack/echo/v4"
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
	return render(c, views.Index([]*auction.Auction{auction.NewAuction("Title #1", time.Now().Add(36*time.Hour), "", *money.New(23000, money.PLN))}))
}

func (h AuctionsHandler) ListGet(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	values := h.srv.List(c.Request().Context(), page)
	if isNone, err := values.IsNone(); isNone {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, values.Unwrap())
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
