package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pallat/micro/order"
	"github.com/pallat/micro/router"
	"github.com/pallat/micro/store"
)

func init() {
	err := godotenv.Load("online.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	r := router.New()

	handler := order.NewHandler(store.NewMongoDBStore(os.Getenv("DSN")), os.Getenv("FILTER_CHANNEL"))
	r.POST("api/v1/orders", handler.Order)

	r.ListenAndServe()()
}
