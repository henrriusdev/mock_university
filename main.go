package main

import (
	"github.com/joho/godotenv"
	"mocku/backend"
)

func main() {
	godotenv.Load()
	backend.MountApp()
}
