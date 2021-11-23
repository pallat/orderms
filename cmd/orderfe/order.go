package main

import (
	"bytes"
	"crypto/tls"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/pallat/micro/order"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main() {
	client := http.Client{
		Timeout: time.Second,
		Transport: &http.Transport{
			Proxy:               http.ProxyFromEnvironment,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConnsPerHost: 10000,
			MaxConnsPerHost:     0,
		},
	}
	//Region,Country,Item Type,Sales Channel,Order Priority,Order Date,Order ID,Ship Date,Units Sold,Unit Price,Unit Cost,Total Revenue,Total Cost,Total Profit
	f, err := os.Open("./Sales_Records.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	target := map[string]string{
		"Offline": os.Getenv("OFFLINE"),
		"Online":  os.Getenv("ONLINE"),
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panic(err)
		}

		orderDate, _ := time.Parse("01/02/2006", record[5])
		shipDate, _ := time.Parse("01/02/2006", record[7])
		orderID, _ := strconv.Atoi(record[6])
		unitsSold, _ := strconv.Atoi(record[8])
		unitPrice, _ := strconv.ParseFloat(record[9], 64)
		unitCost, _ := strconv.ParseFloat(record[10], 64)
		totalRevenue, _ := strconv.ParseFloat(record[11], 64)
		totalCost, _ := strconv.ParseFloat(record[12], 64)
		totalProfit, _ := strconv.ParseFloat(record[13], 64)

		order := order.Order{
			Region:        record[0],
			Country:       record[1],
			ItemType:      record[2],
			SalesChannel:  record[3],
			OrderPriority: record[4],
			OrderDate:     orderDate,
			OrderID:       uint(orderID),
			ShipDate:      shipDate,
			UnitsSold:     uint(unitsSold),
			UnitPrice:     unitPrice,
			UnitCost:      unitCost,
			TotalRevenue:  totalRevenue,
			TotalCost:     totalCost,
			TotalProfit:   totalProfit,
		}
		payload := &bytes.Buffer{}
		err = json.NewEncoder(payload).Encode(&order)
		if err != nil {
			log.Println(order.OrderID, err)
			continue
		}

		req, err := http.NewRequest(http.MethodPost, target[order.SalesChannel], payload)
		if err != nil {
			log.Println(order.OrderID, err)
			continue
		}

		_, err = client.Do(req)
		if err != nil {
			log.Println(order.OrderID, err)
		}

	}
}
