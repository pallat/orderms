# Order Microservices

This project is a demo microservices for my students.

## How it works

cmd/offline is offline API receieve order and filter only offline channel and insert into MariaDB
cmd/online is online API receive order and filter onlt online channel and insert into MongoDB

both of it use same order package in /order
the difference is in main.go in cmd/offline and cmd/online and .env file

## Makefile

maria: to start MariaDB docker container
mongodb: to start MongoDB docker container
offline-image: to build offline api docker image
online-image: to build online api docker image
offline-container: to start offline api docker container
online-container: to start online api docker container

## Run Order

After start MariaDB, MongoDB, Offline and Online contianer

run command:
> cd cmd/orderfe && go run order.go
