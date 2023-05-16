package main

import (
	"github.com/Uchel/chels-laundry/delivery"
	_ "github.com/lib/pq"
)

func main() {
	delivery.Exec()
}
