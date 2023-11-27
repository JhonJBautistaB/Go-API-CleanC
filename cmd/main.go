package main

import (
	"go-api-cleanc/cmd/routes"
	"go-api-cleanc/cmd/server"
	"go-api-cleanc/config"
	"log"
)

func main() {
	//inicializar la configuración
	cfg := config.LoadConfig()
	//crear el servidor
	serve := server.NewServer(cfg.Port)
	//registar rutas en server
	router := serve.InitRouter()
	routes.InitRouter(router)
	//ejecutar el servidor
	if err := serve.RunServer(router); err != nil {
		log.Fatal(err.Error(), err)
	}
}
