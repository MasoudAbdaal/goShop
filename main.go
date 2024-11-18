package main

import (
	"goShop/api"
	"goShop/models"
)

func init() {
	models.InitialModels()
}

func main() {
	api.APIRegister()

	if err := api.Router.Run("localhost:1080"); err != nil {
		panic(err)
	}
}
