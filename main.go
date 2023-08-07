package main

import (
	"go-bankmate/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.Server().Run()
}
