package main

import (
	"github.com/avtara/travair-api/config"
	"github.com/avtara/travair-api/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}