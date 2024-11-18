package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/types/auction"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuctionsHandler struct {
	db             *sqlx.DB
	auctionService services.AuctionService
}

func NewAuctionsHandler(db *sqlx.DB) AuctionsHandler {
	return AuctionsHandler{db, services.NewAuctionService(db)}
}

func (h AuctionsHandler) ListGet(c echo.Context) error {
	return nil
}

func (h AuctionsHandler) CreateGet(c echo.Context) error {
	request := auction.CreateActionRequest{}
	if err := c.Bind(&request); err != nil {
		return err
	}

	res := h.auctionService.Create(c.Request().Context(), request)
	if isNone, err := res.IsNone(); isNone {
		return err
	}

	return c.JSON(http.StatusOK, res.Unwrap())
}
