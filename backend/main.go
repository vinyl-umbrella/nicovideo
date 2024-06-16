package main

import (
	"nicovideo-api/internal/infra/db"
	"nicovideo-api/internal/infra/router"
)

func main() {
	db := db.InitDb()
	r := router.SetupRouter(db)
	r.Run()
}
