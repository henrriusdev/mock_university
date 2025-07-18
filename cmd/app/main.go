package main

import (
	"github.com/joho/godotenv"
	"mocku/api"
)

func main() {
	godotenv.Load()
	api.MountApp()
}
