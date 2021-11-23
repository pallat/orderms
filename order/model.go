package order

import "time"

type Order struct {
	Region        string    `json:"region"`
	Country       string    `json:"country"`
	ItemType      string    `json:"item_type"`
	SalesChannel  string    `json:"sales_channel"`
	OrderPriority string    `json:"order_priority"`
	OrderDate     time.Time `json:"order_date"`
	OrderID       uint      `json:"order_id"`
	ShipDate      time.Time `json:"ship_date"`
	UnitsSold     uint      `json:"units_sold"`
	UnitPrice     float64   `json:"unit_price"`
	UnitCost      float64   `json:"unit_cost"`
	TotalRevenue  float64   `json:"total_revenue"`
	TotalCost     float64   `json:"total_cost"`
	TotalProfit   float64   `json:"total_profit"`
}

// Sub-Saharan Africa,South Africa,Fruits,Offline,M,7/27/2012,443368995,7/28/2012,1593,9.33,6.92,14862.69,11023.56,3839.13
// Middle East and North Africa,Morocco,Clothes,Online,M,9/14/2013,667593514,10/19/2013,4611,109.28,35.84,503890.08,165258.24,338631.84
// Australia and Oceania,Papua New Guinea,Meat,Offline,M,5/15/2015,940995585,6/4/2015,360,421.89,364.69,151880.40,131288.40,20592.00
