package order

import (
	"fmt"
	"net/http"
)

type storer interface {
	Save(Order) error
}

type Context interface {
	Order() (Order, error)
	JSON(int, interface{})
}

type Handler struct {
	store  storer
	filter string
}

func NewHandler(store storer) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Order(c Context) {
	order, err := c.Order()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if order.SalesChannel != h.filter {
		c.JSON(http.StatusNoContent, map[string]string{
			"message": fmt.Sprintf("%s is not accept", order.SalesChannel),
		})
		return
	}

	if err := h.store.Save(order); err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("%d is saved", order.OrderID),
	})
}
