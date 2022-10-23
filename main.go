package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nuttapon-first/summarize-covid19-stats/configs"
	"github.com/nuttapon-first/summarize-covid19-stats/modules/server"
)

func init() {
	err := godotenv.Load("local.env")
	if err != nil {
		log.Printf("[only on local machine] please consider environment variables: %s\n", err)
	}
}

func main() {
	config := new(configs.Configs)

	config.App.Port = os.Getenv("GIN_PORT")

	s := server.NewServer(config)
	s.Start()
}
