package handlers

import (
	"github.com/jmoiron/sqlx"
	"github.com/kalpio/allsell/src/services"
	"github.com/kalpio/allsell/src/types/auction"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AuctionsHandler struct {
	db             *sqlx.DB
	auctionService services.AuctionService
}

func NewAuctionsHandler(db *sqlx.DB) AuctionsHandler {
	return AuctionsHandler{db, services.NewAuctionService(db)}
}

func (h AuctionsHandler) ListGet(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	values := h.auctionService.List(c.Request().Context(), page)
	if isNone, err := values.IsNone(); isNone {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, values.Unwrap())
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
