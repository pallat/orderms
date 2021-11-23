package order

type Order struct {
	Region        string `json:"region"`
	Country       string `json:"country"`
	ItemType      string `json:"item_type"`
	SalesChannel  string `json:"sales_channel"`
	OrderPriority string `json:"order_priority"`
	OrderDate     string `json:"order_date"`
	OrderID       string `json:"order_id"`
	ShipDate      string `json:"ship_date"`
	UnitsSold     string `json:"units_sold"`
	UnitPrice     string `json:"unit_price"`
	UnitCost      string `json:"unit_cost"`
	TotalRevenue  string `json:"total_revenue"`
	TotalCost     string `json:"total_cost"`
	TotalProfit   string `json:"total_profit"`
}
