package main

import (
	"github.com/avtara/travair-api/config"
	"github.com/avtara/travair-api/route"
	"github.com/avtara/travair-api/utils"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}
	e.Logger.Fatal(e.Start(":8080"))
}
