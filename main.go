package main

import (
	"github.com/joho/godotenv"
	"log"
	"quqo_challenge/domain/db"
	"quqo_challenge/infrastructure/config"
	"quqo_challenge/presentation/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	a := config.NewAppConfig()

	p := db.NewPersistence(*a)
	p.AutoMigrate()
	defer p.ClosePersistence()

	app := router.InitRouter(p, *a)
	_ = app.Run(":" + a.AppConfig.Port)
}
