package order

import (
	"testing"
	"time"
)

type fakeStore struct{}

func (fakeStore) Save(Order) error {
	return nil
}

type spyContext struct {
	code     int
	response map[string]string
}

func (spyContext) Order() (Order, error) {
	return Order{
		Region:        "Sub-Saharan Africa",
		Country:       "South Africa",
		ItemType:      "Fruits",
		SalesChannel:  "Offline",
		OrderPriority: "M",
		OrderDate:     time.Date(2012, time.July, 27, 0, 0, 0, 0, time.UTC),
		OrderID:       443368995,
		ShipDate:      time.Date(2012, time.July, 28, 0, 0, 0, 0, time.UTC),
		UnitsSold:     1593,
		UnitPrice:     9.33,
		UnitCost:      6.92,
		TotalRevenue:  14862.69,
		TotalCost:     11023.56,
		TotalProfit:   3839.13,
	}, nil
}
func (c *spyContext) JSON(code int, v interface{}) {
	c.code = code
	c.response = v.(map[string]string)
}

func TestOrderNotAcceptOfflineChannel(t *testing.T) {
	h := &Handler{
		store:  fakeStore{},
		filter: "Online",
	}

	c := spyContext{}
	h.Order(&c)

	want := "Offline is not accept"

	if want != c.response["message"] {
		t.Errorf("%q is expected but got %q\n", want, c.response["message"])
	}
}
