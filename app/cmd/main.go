package main

import (
	"log"
	"product-service/internal/config"
	"product-service/internal/server"
	"product-service/internal/storage"
)

func main() {
	cfg := config.GetConfig()
	storage, err := storage.NewPostgreStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	server := server.New(cfg, storage)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
