package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/ifc-simulador/application/route"
)

func main() {
	fmt.Println("Iniciando o simulador...")

	app := fiber.New()

	route.Routes(app)

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}