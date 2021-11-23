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
		Timeout: 30 * time.Second,
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

		order := order.Order{
			Region:        record[0],
			Country:       record[1],
			ItemType:      record[2],
			SalesChannel:  record[3],
			OrderPriority: record[4],
			OrderDate:     record[5],
			OrderID:       record[6],
			ShipDate:      record[7],
			UnitsSold:     record[8],
			UnitPrice:     record[9],
			UnitCost:      record[10],
			TotalRevenue:  record[11],
			TotalCost:     record[12],
			TotalProfit:   record[13],
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
