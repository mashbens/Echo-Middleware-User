package main

import (
	"gorm-api/config"
	"gorm-api/route"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8004"))

}
